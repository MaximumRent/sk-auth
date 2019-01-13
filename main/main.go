package main

import (
	"github.com/gin-gonic/gin"
	"sk-auth/api"
)

const (
	_DEFAULT_ADDRESS = ":8080"
)

func main() {
	router := gin.Default()
	api.InitOpenApi(router)
	api.InitSecureApi(router)
	router.Run(_DEFAULT_ADDRESS)
}
