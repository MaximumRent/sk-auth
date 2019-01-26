package worker

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Broker interface {
	Start()
	Stop()
}

// Config types
const (
	_CONFIG_FILE_PATH = "./config.yml"
)

// Acceptable config types
const (
	_BROKER_RABBITMQ_TYPE = "rabbitmq"
)

func GetBroker() Broker {
	yamlConf := make(map[interface{}]interface{})
	data, err := ioutil.ReadFile(_CONFIG_FILE_PATH)
	if err != nil {
		log.Fatal("Error in reading file. Cause: %s", err)
	}
	err = yaml.Unmarshal([]byte(data), &yamlConf)
	brokerDef := yamlConf["worker"].(map[interface{}]interface{})["broker"].(map[interface{}]interface{})
	brokerType := brokerDef["type"]
	switch brokerType {
	case _BROKER_RABBITMQ_TYPE:
		return getRabbitMqBroker(brokerDef)
	default:
		return nil
	}
}
