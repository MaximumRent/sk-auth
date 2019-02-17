package main

import (
	"github.com/gin-gonic/gin"
	"sk-auth/api"
	"sk-auth/mongo"
	"sk-auth/worker"
)

const (
	_DEFAULT_ADDRESS = ":7070"
)

func main() {
	router := gin.Default()
	api.InitMiddleware(router)
	api.InitOpenApi(router)
	api.InitSecureApi(router)
	mongo.InitMongoDb()
	defer mongo
	go worker.GetBroker().Start()
	router.Run(_DEFAULT_ADDRESS)
}
