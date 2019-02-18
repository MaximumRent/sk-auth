package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sk-auth/auth/entity"
	"sk-auth/mongo"
	"sk-auth/util"
)

// API paths
const (
	REGISTER_USER_PATH = "/register"
	LOGIN_USER_PATH    = "/login"
)

func InitOpenApi(router *gin.Engine) {
	openGroup := router.Group(util.OPEN_API_GROUP)
	openGroup.Use(MessageMappingMiddleware, EntityValidationMiddleware)
	{
		openGroup.POST(REGISTER_USER_PATH, registerUser)
		openGroup.POST(LOGIN_USER_PATH, loginUser)
	}
}

func loginUser(context *gin.Context) {

}

// REGISTER_USER_PATH
func registerUser(context *gin.Context) {
	payload := context.MustGet(PAYLOAD_KEY)
	user := payload.(*entity.User)
	err := mongo.CreateUser(*user)
	if err != nil {
		message := *INVALID_REGISTRATION_MESSAGE
		message.Payload = err
		context.JSON(http.StatusBadRequest, gin.H{"message": message})
	} else {
		message := *SUCCESSFULL_REGISTRATION_MESSAGE
		message.Payload = user
		context.JSON(http.StatusOK, gin.H{"message": message})
	}
}
