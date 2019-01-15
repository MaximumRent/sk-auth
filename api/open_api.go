package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sk-auth/auth/entity"
)

// API paths
const (
	OPEN_API_GROUP     = "/open"
	REGISTER_USER_PATH = "/register"
)

func InitOpenApi(router *gin.Engine) {
	openGroup := router.Group(OPEN_API_GROUP)
	{
		openGroup.POST(REGISTER_USER_PATH, registerUser)
	}
}

// REGISTER_USER_PATH
func registerUser(context *gin.Context) {
	var user *entity.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	println(user.Password)
	context.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}
