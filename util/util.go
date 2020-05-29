package util

import (
	"MindaZepeto/models"
	"crypto/md5"
	"encoding/hex"
)

//GetMD5Encode 使用ＭＤ５加密
func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// GetReturnData 对返回data的包裹
func GetReturnData(list interface{}, msgString string) *models.Result {
	result := new(models.Result)
	result.Data = list
	result.Code = 200
	result.Msg = msgString
	return result
}
