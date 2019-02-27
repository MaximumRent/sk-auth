package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"sk-auth/api"
	"sk-auth/mongo"
	"sk-auth/util"
	"sk-auth/worker"
)

func main() {
	router := gin.Default()
	api.InitOpenApi(router)
	api.InitSecureApi(router)
	mongo.InitMongoDb()
	defer mongo.CloseConnection()
	go worker.GetBroker().Start()
	router.Run(getDefaultAddress())
}

func getDefaultAddress() string {
	yamlConfig, err := util.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	return yamlConfig["sk"].(map[interface{}]interface{})["address"].(string)
}
