package util

import (
	"MindaZepeto/models"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"strconv"
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

//GetGroupPhotoTableName 返回用户的相册应当存放在哪一张表中。
func GetGroupPhotoTableName(id string) string {
	id_num, _ := strconv.Atoi(id)
	if id_num <= 2000 {
		return "grouphotoaaa"
	} else if id_num <= 4000 {
		return "grouphotoaab"
	} else if id_num <= 6000 {
		return "grouphotoaac"
	} else if id_num <= 8000 {
		return "grouphotoaba"
	} else if id_num <= 10000 {
		return "grouphotoabb"
	} else if id_num <= 12000 {
		return "grouphotoabc"
	} else if id_num <= 14000 {
		return "grouphotoaca"
	} else if id_num <= 16000 {
		return "grouphotoacb"
	} else if id_num <= 18000 {
		return "grouphotoacc"
	} else if id_num <= 20000 {
		return "grouphotobaa"
	} else if id_num <= 22000 {
		return "grouphotobab"
	} else if id_num <= 24000 {
		return "grouphotobac"
	} else if id_num <= 26000 {
		return "grouphotobba"
	} else {
		return "grouphotoend"
	}
}
