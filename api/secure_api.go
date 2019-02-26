package api

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sk-auth/auth/entity"
	"sk-auth/mongo"
	"sk-auth/util"
)

const (
	LOGOUT_USER_PATH    = "/logout"
	VALIDATE_TOKEN_PATH = "/validate"
	UPDATE_USER_INFO    = "/user/update"
)

func InitSecureApi(router *gin.Engine) {
	secureGroup := router.Group(util.SECURE_API_GROUP)
	secureGroup.Use(MessageMappingMiddleware, EntityValidationMiddleware, TokenValidationMiddleware)
	{
		secureGroup.POST(LOGOUT_USER_PATH, logoutUser)
		secureGroup.POST(VALIDATE_TOKEN_PATH, validateToken)
		secureGroup.POST(UPDATE_USER_INFO, updateUserInfo)
	}
}

func updateUserInfo(context *gin.Context) {
	payload := context.MustGet(USER_INFO_KEY)
	user := payload.(*entity.User)
	if user.Email == "" || user.Password == "" || user.Nickname == "" {
		message := *INVALID_UPDATE_USERINFO_MESSAGE
		message.Payload = errors.New("Email, Password or Nickname can't be empty!")
		context.JSON(http.StatusBadRequest, gin.H{"message": message})
	}
	err := mongo.UpdateUserInfo(*user)
	if err != nil {
		message := *INVALID_UPDATE_USERINFO_MESSAGE
		message.Payload = err
		context.JSON(http.StatusBadRequest, gin.H{"message": message})
	} else {
		message := *SUCCESSFULL_UPDATE_USERINFO_MESSAGE
		message.Payload = user
		context.JSON(http.StatusOK, gin.H{"message": message})
	}
}

// VALIDATE_TOKEN_PATH
// Main validation will do in TokenValidationMiddleware
// If validation will be success, this endpoint return success message
func validateToken(context *gin.Context) {
	message := *SUCCESSFULL_TOKEN_VALIDATION_MESSAGE
	message.Payload = context.MustGet(SHORT_USER_INFO_KEY)
	context.JSON(http.StatusOK, gin.H{"message": message})
}

func logoutUser(context *gin.Context) {
	payload := context.MustGet(PAYLOAD_KEY)
	logoutInfo := payload.(*entity.LogoutUserInfo)
	err := mongo.Logout(logoutInfo.Email, logoutInfo.Nickname, logoutInfo.AuthDevice, logoutInfo.Token)
	if err != nil {
		message := *INVALID_LOGOUT_MESSAGE
		message.Payload = err
		context.JSON(http.StatusBadRequest, gin.H{"message": message})
	} else {
		message := *SUCCESSFULL_LOGOUT_MESSAGE
		message.Payload = ""
		context.JSON(http.StatusOK, gin.H{"message": message})
	}
}
