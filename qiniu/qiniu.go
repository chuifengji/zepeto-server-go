package qiniu

import (
	"MindaZepeto/config"
	"fmt"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

func GetUptoken(fileName string) string {
	accessKey := config.Sysconfig.AccessKeyQn
	secretKey := config.Sysconfig.SecretKeyQn
	bucket := config.Sysconfig.BucketName
	keyToOverwrite := fileName //覆盖上传需要获取文件名
	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", bucket, keyToOverwrite),
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

func GetUptokenPhotos() string {
	accessKey := config.Sysconfig.AccessKeyQn
	secretKey := config.Sysconfig.SecretKeyQn
	bucket := config.Sysconfig.BucketName
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
