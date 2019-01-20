package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"net/http"
	"sk-auth/auth/entity"
	"sk-auth/validation"
)

// Message codes
const (
	_MAP_PAYLOAD_TO_USER_CODE = 1
)

const (
	MESSAGE_KEY = "message"
	PAYLOAD_KEY = "message.payload"
)

func InitMiddleware(router *gin.Engine) {

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
	case _MAP_PAYLOAD_TO_USER_CODE:
		var user = entity.CreateUser()
		if err := mapstructure.Decode(payload, user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return user
	default:
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid code type"})
		return nil
	}
}