package controllers

import (
	"net/http"

	"go-auth/model"
	"go-auth/utils/ids"
	"go-auth/utils/messages"

	"github.com/gin-gonic/gin"
)

func CheckHandler(c *gin.Context) {
	c.JSON(http.StatusOK, model.BaseResponse{Message: messages.SUCCESS, Code: ids.SUCCESS})
}
