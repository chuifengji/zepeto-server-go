package config

import (
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

var Sysconfig = &sysconfig{}

func init() {
	b, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic(err)
	} else {
		err = jsoniter.Unmarshal(b, Sysconfig)
		if err != nil {
			panic(err)
		}
	}
}

type sysconfig struct {
	Port        string `json:"Port"`
	AppID       string `json:"AppID"`
	SecretKeyWx string `json:"SecretKeyWx"`
	AccessKeyQn string `json:"AccessKeyQn"`
	SecretKeyQn string `json:"SecretKeyQn"`
	BucketName  string `json:"BucketName"`
	BucketNameb string `json:"BucketNameb"`
	DBUserName  string `json:"DBUserName"`
	DBPassword  string `json:"DBPassword"`
	DBIp        string `json:"DBIp"`
	DBPort      string `json:"DBPort"`
	DBName      string `json:"DBName"`
	UserToken   string `json:"UserToken"`
}
