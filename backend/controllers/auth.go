package controllers

import (
	"context"

	"go-auth/helpers"
	"go-auth/model"
	"go-auth/redis"
	"go-auth/utils/ids"
	"go-auth/utils/messages"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func HandleAuthCallback(c *gin.Context) {
	res := c.Writer
	req := c.Request
	// Get the value of the "provider" path parameter
	provider := c.Param("provider")
	// Set the provider value in the request context
	ctx := context.WithValue(req.Context(), "provider", provider)
	req = req.WithContext(ctx)
	user, err := gothic.CompleteUserAuth(res, req)
	if err != nil {
		// c.String(http.StatusInternalServerError, fmt.Sprintf("Authentication error: %s", err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.BaseResponse{Message: messages.AUTH_ERROR, Code: ids.AUTH_ERROR})
		return
	}
	session, err := helpers.CreateSession(redis.GetRedisClient(), &user)
	if err != nil {
		// Handle error
		c.AbortWithStatusJSON(http.StatusInternalServerError, model.BaseResponse{Message: messages.WENT_WRONG, Code: ids.WENT_WRONG})
	}
	// Set the session ID in the cookie
	c.SetCookie("session_id", session.ID, 60*60, "/", "localhost", true, false)
	// disable-caching
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	// user.AccessToken
	c.Redirect(http.StatusFound, "http://localhost:5173")
}

func HandleAuthProvider(c *gin.Context) {
	res := c.Writer
	req := c.Request
	provider := c.Param("provider")
	// Set the provider value in the request context
	ctx := context.WithValue(req.Context(), "provider", provider)
	req = req.WithContext(ctx)
	gothic.BeginAuthHandler(res, req)
}

func LogoutHandler(c *gin.Context) {
	sessionId, err := c.Request.Cookie("session_id")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, model.BaseResponse{Message: messages.INVALID_SESSION, Code: ids.INVALID_SESSION})
	}
	helpers.DeleteSession(redis.GetRedisClient(), sessionId.Value)
	c.JSON(http.StatusOK, model.BaseResponse{Message: messages.SUCCESS, Code: ids.SUCCESS})
}
