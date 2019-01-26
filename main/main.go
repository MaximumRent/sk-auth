package main

import (
	"github.com/gin-gonic/gin"
	"sk-auth/api"
	"sk-auth/worker"
)

const (
	_DEFAULT_ADDRESS = ":8080"
)

func main() {
	router := gin.Default()
	api.InitMiddleware(router)
	api.InitOpenApi(router)
	api.InitSecureApi(router)
	go worker.GetBroker().Start()
	router.Run(_DEFAULT_ADDRESS)
}
