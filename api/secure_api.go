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
	ADD_NEW_ROLE_TO_USER = "/user/role/add"
	DELETE_USER_ROLE = "/user/role/delete"
	GET_USER_INFO = "/user/info"
)

func InitSecureApi(router *gin.Engine) {
	secureGroup := router.Group(util.SECURE_API_GROUP)
	secureGroup.Use(MessageMappingMiddleware, EntityValidationMiddleware, TokenValidationMiddleware)
	{
		secureGroup.GET(GET_USER_INFO, getUser)
		secureGroup.POST(LOGOUT_USER_PATH, logoutUser)
		secureGroup.POST(VALIDATE_TOKEN_PATH, validateToken)
		secureGroup.POST(UPDATE_USER_INFO, updateUserInfo)
		secureGroup.POST(USER_HAS_ACCESS_TO, checkPermissions)
		secureGroup.POST(ADD_NEW_ROLE_TO_USER, addRole)
		secureGroup.POST(DELETE_USER_ROLE, deleteRole)
	}
}

// GET_USER_INFO
func getUser(context *gin.Context) {
	payload := context.MustGet(PAYLOAD_KEY)
	shortUserInfo := payload.(*entity.ShortUser)
	user := new(entity.User)
	err := mongo.GetCurrentUser(shortUserInfo, user)
	if err != nil {
		message := *INVALID_GET_USER_MESSAGE
		message.Payload = err
		context.JSON(http.StatusBadRequest, gin.H{"message": message})
	} else {
		message := SUCCESSFULL_GET_USER_MESSAGE
		message.Payload = user
		context.JSON(http.StatusOK, gin.H{"message": message})
	}
}

// DELETE_USER_ROLE
func deleteRole(context *gin.Context) {
	payload := context.MustGet(PAYLOAD_KEY)
	shortUserInfo := context.MustGet(SHORT_USER_INFO_KEY).(*entity.ShortUser)
	changeRoleRequest := payload.(*entity.ChangeRoleRequest)
	err := mongo.DeleteUserRole(shortUserInfo.Email, shortUserInfo.Nickname, changeRoleRequest.RoleName)
	if err != nil {
		message := *INVALID_DELETE_ROLE_MESSAGE
		message.Payload = err
		context.JSON(http.StatusBadRequest, gin.H{"message": message})
	} else {
		message := SUCCESSFULL_DELETE_ROLE_MESSAGE
		message.Payload = changeRoleRequest
		context.JSON(http.StatusOK, gin.H{"message": message})
	}
}

// ADD_NEW_ROLE_TO_USER
func addRole(context *gin.Context) {
	payload := context.MustGet(PAYLOAD_KEY)
	shortUserInfo := context.MustGet(SHORT_USER_INFO_KEY).(*entity.ShortUser)
	changeRoleRequest := payload.(*entity.ChangeRoleRequest)
	err := mongo.AddRoleToUser(shortUserInfo.Email, shortUserInfo.Nickname, changeRoleRequest.RoleName)
	if err != nil {
		message := *INVALID_ADD_ROLE_MESSAGE
		message.Payload = err
		context.JSON(http.StatusBadRequest, gin.H{"message": message})
	} else {
		message := SUCCESSFULL_ADD_ROLE_MESSAGE
		message.Payload = changeRoleRequest
		context.JSON(http.StatusOK, gin.H{"message": message})
	}
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
