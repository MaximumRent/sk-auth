package api

type Message struct {
	Message string      `json:"message"`
	Code    int64       `json:"code"`
	Payload interface{} `json:"payload"`
}

// SUCCESSFUL MESSAGES
var (
	SUCCESSFULL_REGISTRATION_MESSAGE = &Message{
		Message: "Registration completed successful.",
		Code:    _SUCCESSFULL_REGISTRATION,
	}
	SUCCESSFULL_LOGIN_MESSAGE = &Message{
		Message: "Login completed successful.",
		Code:    _SUCCESSFULL_LOG_IN,
	}
	SUCCESSFULL_UPDATE_USERINFO_MESSAGE = &Message{
		Message: "User info updated successful.",
		Code:    _SUCCESSFULL_UPDATE_USERINFO,
	}
	SUCCESSFULL_TOKEN_VALIDATION_MESSAGE = &Message{
		Message: "Your token is valid.",
		Code:    _SUCCESSFULL_TOKEN_VALIDATION,
	}
	SUCCESSFULL_LOGOUT_MESSAGE = &Message{
		Message: "Your logout completed successful.",
		Code:    _SUCCESSFULL_LOGOUT,
	}
)

// ERROR MESSAGES
var (
	INVALID_REGISTRATION_MESSAGE = &Message{
		Message: "Registration wasn't completed successful.",
		Code:    _INVALID_REGISTRATION,
	}
	INVALID_LOGIN_MESSAGE = &Message{
		Message: "Invalid login or password.",
		Code:    _INVALID_LOG_IN,
	}
	INVALID_UPDATE_USERINFO_MESSAGE = &Message{
		Message: "User info didn't update.",
		Code:    _INVALID_UPDATE_USERINFO,
	}
	INVALID_TOKEN_MESSAGE = &Message{
		Message: "You haven't valid token.",
		Code:    _INVALID_TOKEN,
	}
	INVALID_LOGOUT_MESSAGE = &Message{
		Message: "Something went wrong across your logout. Please try again.",
		Code:    _INVALID_LOGOUT,
	}
)
