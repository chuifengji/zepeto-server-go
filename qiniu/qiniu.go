package qiniu

import (
	"MindaZepeto/config"

	"github.com/qiniu/api.v7/storage"
	"github.com/qiniu/api.v7/v7/auth/qbox"
)

func getUptoken() string {
	accessKey := config.Sysconfig.AccessKeyQn
	secretKey := config.Sysconfig.SecretKeyQn
	bucket := config.Sysconfig.BucketName
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 14400 //upToken的有效期为４个小时。
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
