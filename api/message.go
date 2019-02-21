package api

type Message struct {
	Message string      `json:"message"`
	Code    int64       `json:"code"`
	Payload interface{} `json:"payload"`
}

var (
	SUCCESSFULL_REGISTRATION_MESSAGE = &Message{
		Message: "Registration completed succesfull",
		Code:    _SUCCESSFULL_REGISTRATION,
	}
	SUCCESSFULL_LOGIN_MESSAGE = &Message{
		Message: "Login completed succesfull",
		Code:    _SUCCESSFULL_LOG_IN,
	}
	INVALID_REGISTRATION_MESSAGE = &Message{
		Message: "Registration wasn't completed successfull",
		Code:    _INVALID_REGISTRATION,
	}
	INVALID_LOGIN_MESSAGE = &Message{
		Message: "Invalid login or password",
		Code:    _INVALID_LOG_IN,
	}
)
