package qiniu

import (
	"MindaZepeto/config"
	"fmt"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"github.com/qiniu/x/rpc"
)

//GetUptoken 覆盖上传所需的token
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

//GetUptokenPhotos 获取图片上传的token.
func GetUptokenPhotos() string {
	accessKey := config.Sysconfig.AccessKeyQn
	secretKey := config.Sysconfig.SecretKeyQn
	bucket := config.Sysconfig.BucketNameb
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 14400
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

// DeleteSingleImage 删除单个图片
func DeleteSingleImage(fileName string) {
	accessKey := config.Sysconfig.AccessKeyQn
	secretKey := config.Sysconfig.SecretKeyQn
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		UseHTTPS: true, // 是否使用https域名进行资源管理
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	bucket := config.Sysconfig.BucketName
	key := fileName
	err := bucketManager.Delete(bucket, key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

//DeleteImages 批量删除图片
func DeleteImages() {
	accessKey := config.Sysconfig.AccessKeyQn
	secretKey := config.Sysconfig.SecretKeyQn
	mac := qbox.NewMac(accessKey, secretKey)
	cfg := storage.Config{
		UseHTTPS: true, // 是否使用https域名进行资源管理
	}
	bucketManager := storage.NewBucketManager(mac, &cfg)
	bucket := config.Sysconfig.BucketName
	keys := []string{
		"github1.png",
		"github2.png",
		"github3.png",
		"github4.png",
		"github5.png",
	}
	deleteOps := make([]string, 0, len(keys))
	for _, key := range keys {
		deleteOps = append(deleteOps, storage.URIDelete(bucket, key))
	}
	rets, err := bucketManager.Batch(deleteOps)
	if err != nil {
		// 遇到错误
		if _, ok := err.(*rpc.ErrorInfo); ok {
			for _, ret := range rets {
				// 200 为成功
				fmt.Printf("%d\n", ret.Code)
				if ret.Code != 200 {
					fmt.Printf("%s\n", ret.Data.Error)
				}
			}
		} else {
			fmt.Printf("batch error, %s", err)
		}
	} else {
		// 完全成功
		for _, ret := range rets {
			// 200 为成功
			fmt.Printf("%d\n", ret.Code)
		}
	}
}
