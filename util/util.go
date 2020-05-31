package util

import (
	"MindaZepeto/models"
	"crypto/sha1"
	"encoding/hex"
	"io"
)

//GetSHAEncode 加密
func GetSHAEncode(str string) string {
	w := sha1.New()
	io.WriteString(w, str)            //将str写入到w中
	bw := w.Sum(nil)                  //w.Sum(nil)将w的hash转成[]byte格式
	shastr2 := hex.EncodeToString(bw) //将 bw 转成字符串
	return shastr2
}

// GetReturnData 对返回data的包裹
func GetReturnData(list interface{}, msgString string) *models.Result {
	result := new(models.Result)
	result.Data = list
	result.Code = 200
	result.Msg = msgString
	return result
}
