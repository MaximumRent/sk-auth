package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"sk-auth/api"
	"sk-auth/mongo"
	"sk-auth/util"
)

const (
	_DEFAULT_ADDRESS = ":7070"
)

func main() {
	router := gin.Default()
	api.InitOpenApi(router)
	api.InitSecureApi(router)
	mongo.InitMongoDb()
	defer mongo.CloseConnection()
	//go worker.GetBroker().Start()
	router.Run(getDefaultAddress())
}

func getDefaultAddress() string {
	yamlConfig, err := util.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	return yamlConfig["sk"].(map[interface{}]interface{})["address"].(string)
}
