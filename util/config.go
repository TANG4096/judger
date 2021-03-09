package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

var ConfigMap map[string]map[string]string
var DbConfigMap map[string]map[string]string

func open(fileName string) map[string]map[string]string {
	configMap := make(map[string]map[string]string)

	//log.Println(dir)
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("打开配置文件失败 %s\n", err.Error())
	}
	data, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("读取配置文件失败 %s\n", err.Error())
	}
	err = json.Unmarshal(data, &configMap)
	if err != nil {
		log.Fatalf(err.Error())
	}
	return configMap
}

func init() {
	dir := GetPath()
	ConfigMap = open(dir + "/config/config.json")
	DbConfigMap = open(dir + "/config/db/config.json")
}

func GetConfigMap(str string) map[string]string {
	mp := ConfigMap[str]
	if len(mp) == 0 {
		return nil
	}
	return mp
}

func GetDbConfigMap(dbName string) map[string]string {
	mp := DbConfigMap[dbName]
	if len(mp) == 0 {
		return nil
	}
	return mp
}
