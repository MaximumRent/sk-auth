package mongo

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"sk-auth/util"
)

var client *mongo.Client

var MongoUrl string

func readMongoConfig() {
	yamlConfig, err := util.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	MongoUrl = yamlConfig["mongo"].(map[interface{}]interface{})["url"].(string)
}

func InitMongoDb() {
	var err error
	readMongoConfig()
	client, err = mongo.Connect(context.TODO(), MongoUrl)
	if err != nil {
		log.Fatal("Can't create connection to mongodb. Cause: %s", err)
	}

}

func CloseConnection() {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
