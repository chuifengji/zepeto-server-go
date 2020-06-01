package config

import (
	"io/ioutil"

	jsoniter "github.com/json-iterator/go"
)

var Sysconfig = &sysconfig{}

func init() {
	// path, _ := filepath.Abs("./config/config.json")
	// fmt.Println(path)
	b, err := ioutil.ReadFile("./config.json") ///root/server/mindazepeto/config.json 线上版路径
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
	DBUserName  string `json:"DBUserName"`
	DBPassword  string `json:"DBPassword"`
	DBIp        string `json:"DBIp"`
	DBPort      string `json:"DBPort"`
	DBName      string `json:"DBName"`
	UserToken   string `json:"UserToken"`
}
