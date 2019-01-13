package api

import "github.com/gin-gonic/gin"

// API paths
const (
	REGISTER_USER_PATH = "/register"
)

func InitOpenApi(router *gin.Engine) {
	router.POST(REGISTER_USER_PATH, registerUser)
}

// REGISTER_USER_PATH
func registerUser(context *gin.Context) {

}
