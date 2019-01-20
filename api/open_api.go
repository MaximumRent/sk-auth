package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// API paths
const (
	OPEN_API_GROUP     = "/open"
	REGISTER_USER_PATH = "/register"
)

func InitOpenApi(router *gin.Engine) {
	openGroup := router.Group(OPEN_API_GROUP)
	openGroup.Use(MessageMappingMiddleware, EntityValidationMiddleware)
	{
		openGroup.POST(REGISTER_USER_PATH, registerUser)
	}
}

// REGISTER_USER_PATH
func registerUser(context *gin.Context) {
	payload := context.MustGet(PAYLOAD_KEY)
	context.JSON(http.StatusOK, gin.H{"status": payload})
}
