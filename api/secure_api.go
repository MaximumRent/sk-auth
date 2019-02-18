package api

import (
	"github.com/gin-gonic/gin"
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
	println("updateUserInfo")
}

func validateToken(context *gin.Context) {
	println("validateToken")
}

func logoutUser(context *gin.Context) {
	println("logoutUser")
}
