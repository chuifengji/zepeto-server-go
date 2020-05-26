package config

import (
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

var Sysconfig = &sysconfig{}

func init() {
	//指定对应的json配置文件
	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("Sys config read err")
	} else {
		err = jsoniter.Unmarshal(b, Sysconfig)
		if err != nil {
			panic(err)
		}
	}
}

type sysconfig struct {
	Port       string `json:"Port"`
	AppID      string `json:"AppID"`
	SecretKey  string `json:"SecretKey"`
	DBUserName string `json:"DBUserName"`
	DBPassword string `json:"DBPassword"`
	DBIp       string `json:"DBIp"`
	DBPort     string `json:"DBPort"`
	DBName     string `json:"DBName"`
}