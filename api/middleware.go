package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"sk-auth/auth/entity"
	"sk-auth/mongo"
	"sk-auth/validation"
)

const (
	MESSAGE_KEY         = "message"
	PAYLOAD_KEY         = "message.payload"
	USER_INFO_KEY       = "message.payload.userinfo"
	SHORT_USER_INFO_KEY = "message.payload.short"
)

// Middleware which validates user token
func TokenValidationMiddleware(context *gin.Context) {
	payload := context.MustGet(SHORT_USER_INFO_KEY)
	shortUser := payload.(*entity.ShortUser)
	err := mongo.ValidateAuthToken(shortUser.Email, shortUser.Nickname, shortUser.Token)
	if err != nil {
		message := *INVALID_TOKEN_MESSAGE
		message.Payload = err
		context.JSON(http.StatusUnauthorized, gin.H{"message": message})
		context.Abort()
		return
	}
	context.Set(SHORT_USER_INFO_KEY, shortUser)
}

func MessageMappingMiddleware(context *gin.Context) {
	var message *Message
	if err := context.BindJSON(&message); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.Set(MESSAGE_KEY, message)
}

func EntityValidationMiddleware(context *gin.Context) {
	message, isPresented := context.Get(MESSAGE_KEY)
	if !isPresented {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Entity validation middleware can't find message!"})
		return
	}
	payload := extractPayload(message.(*Message), context)
	if err := payload.SelfValidate(); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.Set(PAYLOAD_KEY, payload)
}

// Extracts self-validatable data from payload based on message code
// Current codes:
// 1 - payload contains user info
func extractPayload(message *Message, context *gin.Context) validation.SelfValidatable {
	payload := message.Payload.(map[string]interface{})
	switch message.Code {
	case _MAP_PAYLOAD_TO_USER:
		var user = entity.CreateUser()
		if err := mapstructure.Decode(payload, user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return user
	case _MAP_PAYLOAD_TO_SHORT_USER:
		shortUser := new(entity.ShortUser)
		if err := mapstructure.Decode(payload, shortUser); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return shortUser
	case _MAP_PAYLOAD_TO_COMPLEX_USER:
		shortUser := new(entity.ShortUser)
		if err := mapstructure.Decode(payload, shortUser); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		var user = entity.CreateUser()
		if err := mapstructure.Decode(payload, user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		context.Set(USER_INFO_KEY, user)
		context.Set(SHORT_USER_INFO_KEY, shortUser)
		return shortUser
	case _MAP_PAYLOAD_TO_LOGIN_INFO:
		loginInfo := new(entity.LoginUserInfo)
		if err := mapstructure.Decode(payload, loginInfo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return loginInfo
	case _MAP_PAYLOAD_TO_LOGOUT_USER:
		logoutInfo := new(entity.LogoutUserInfo)
		if err := mapstructure.Decode(payload, logoutInfo); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		shortUser := new(entity.ShortUser)
		if err := mapstructure.Decode(payload, shortUser); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		context.Set(SHORT_USER_INFO_KEY, shortUser)
		return logoutInfo
	case _MAP_PAYLOAD_TO_TOKEN_VALIDATION:
		accessRequest := new(entity.AccessRequest)
		if err := mapstructure.Decode(payload, accessRequest); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return accessRequest
	default:
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid code type"})
		return nil
	}
}
