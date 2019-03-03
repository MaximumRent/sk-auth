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
	USER_HAS_ACCESS_TO    = "/user/access"
	ADD_NEW_ROLE_TO_USER = "/user/add/role"
)

func InitSecureApi(router *gin.Engine) {
	secureGroup := router.Group(util.SECURE_API_GROUP)
	secureGroup.Use(MessageMappingMiddleware, EntityValidationMiddleware, TokenValidationMiddleware)
	{
		secureGroup.POST(LOGOUT_USER_PATH, logoutUser)
		secureGroup.POST(VALIDATE_TOKEN_PATH, validateToken)
		secureGroup.POST(UPDATE_USER_INFO, updateUserInfo)
		secureGroup.POST(USER_HAS_ACCESS_TO, checkPermissions)
		secureGroup.POST(ADD_NEW_ROLE_TO_USER, addRole)
	}
}

func addRole(context *gin.Context) {
	payload := context.MustGet(PAYLOAD_KEY)
	shortUserInfo := context.MustGet(SHORT_USER_INFO_KEY).(*entity.ShortUser)

}

// USER_HAS_ACCESS_TO
func checkPermissions(context *gin.Context) {
	payload := context.MustGet(PAYLOAD_KEY)
	shortUserInfo := context.MustGet(SHORT_USER_INFO_KEY).(*entity.ShortUser)
	accessRequest := payload.(*entity.AccessRequest)
	err := mongo.UserHasAccessTo(accessRequest, shortUserInfo)
	if err != nil {
		message := *INVALID_PERMISSION_MESSAGE
		message.Payload = err
		context.JSON(http.StatusBadRequest, gin.H{"message": message})
	} else {
		message := *SUCCESSFULL_PERMISSION_MESSAGE
		message.Payload = accessRequest
		context.JSON(http.StatusOK, gin.H{"message": message})
	}
}

// UPDATE_USER_INFO
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

// LOGOUT_USER_PATH
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
