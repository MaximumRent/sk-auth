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
	payload, isPresented := context.Get(PAYLOAD_KEY)
	if !isPresented {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Payload not presented!"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"status": payload})
}
