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
	REGISTER_USER_PATH 			= "/registration"
	LOGIN_USER_PATH    			= "/login"
	GOOGLE_OAUTH    			= "/login/google"
	GOOGLE_OAUTH_CALLBACK    	= "/oauth/google/callback"
)

func InitOpenApi(router *gin.Engine) {
	openGroup := router.Group(util.OPEN_API_GROUP)
	openGroup.Use(MessageMappingMiddleware, EntityValidationMiddleware)
	{
		openGroup.POST(REGISTER_USER_PATH, registerUser)
		openGroup.POST(LOGIN_USER_PATH, login)
		openGroup.POST(GOOGLE_OAUTH, googleOauth)
		openGroup.POST(GOOGLE_OAUTH_CALLBACK, googleOauthCallback)
	}
}

func googleOauth(context *gin.Context) {

}

func googleOauthCallback(context *gin.Context) {

}

// LOGIN_USER_PATH
func login(context *gin.Context) {
	payload := context.MustGet(PAYLOAD_KEY)
	loginUserInfo := payload.(*entity.LoginUserInfo)
	err, token, user := mongo.LoginUser(*loginUserInfo)
	if err != nil {
		message := *INVALID_LOGIN_MESSAGE
		message.Payload = err
		context.JSON(http.StatusBadRequest, gin.H{"message": message})
	} else {
		message := *SUCCESSFULL_LOGIN_MESSAGE
		message.Payload = &entity.ShortUser{
			Email:    user.Email,
			Nickname: user.Nickname,
			Token:    token.Token,
		}
		context.JSON(http.StatusOK, gin.H{"message": message})
	}
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
