package middleware

import (
	"go-auth/model"
	"go-auth/redis"
	"go-auth/utils/ids"
	"go-auth/utils/messages"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId, err := c.Request.Cookie("session_id")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.BaseResponse{Message: messages.INVALID_SESSION, Code: ids.INVALID_SESSION})
		}
		// Verify session data in Redis
		_, err = redis.GetValue(sessionId.Value)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.BaseResponse{Message: messages.INVALID_SESSION, Code: ids.INVALID_SESSION})
		}
		// Session is valid, continue processing
		c.Next()
	}
}
