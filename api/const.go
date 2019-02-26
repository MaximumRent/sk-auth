package api

// Map codes
const (
	_MAP_PAYLOAD_TO_USER         = 1
	_MAP_PAYLOAD_TO_SHORT_USER   = 2
	_MAP_PAYLOAD_TO_COMPLEX_USER = 3
	_MAP_PAYLOAD_TO_LOGIN_INFO   = 4
	_MAP_PAYLOAD_TO_LOGOUT_USER  = 5
)

// Success message codes
const (
	_SUCCESSFULL_REGISTRATION     = 100
	_SUCCESSFULL_LOG_IN           = 101
	_SUCCESSFULL_UPDATE_USERINFO  = 102
	_SUCCESSFULL_TOKEN_VALIDATION = 103
	_SUCCESSFULL_LOGOUT           = 104
)

// Error message codes
const (
	_INVALID_REGISTRATION    = 200
	_INVALID_LOG_IN          = 201
	_INVALID_UPDATE_USERINFO = 202
	_INVALID_TOKEN           = 203
	_INVALID_LOGOUT          = 204
)
