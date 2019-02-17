package util

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func ReadConfig() (map[interface{}]interface{}, error) {
	yamlConf := make(map[interface{}]interface{})
	data, err := ioutil.ReadFile(SYS_CONFIG_FILE_PATH)
	if err != nil {
		log.Fatal("Error in reading file. Cause: %s", err)
	}
	err = yaml.Unmarshal([]byte(data), &yamlConf)
	return yamlConf, err
}
