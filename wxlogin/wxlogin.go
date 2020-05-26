package wxlogin

import (
	"MindaZepeto/config"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type WXLoginResp struct {
	OpenID     string `json:"OpenID"`
	SessionKey string `json:"SessionKey"`
	UnionID    string `json:"UnionID"`
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
}

//WXLogin  以 code 作为输入, 返回调用微信接口得到的对象指针和异常情况
func WXLogin(code string) (*WXLoginResp, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

	// 合成url,
	url = fmt.Sprintf(url, config.Sysconfig.AppID, config.Sysconfig.SecretKey, code)

	// 创建http get请求
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 解析http请求中body 数据到我们定义的结构体中
	wxResp := WXLoginResp{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&wxResp); err != nil {
		return nil, err
	}

	// 判断微信接口返回的是否是一个异常情况
	if wxResp.ErrCode != 0 {
		return nil, errors.New(fmt.Sprintf("ErrCode:%s  ErrMsg:%s", wxResp.ErrCode, wxResp.ErrMsg))
	}

	return &wxResp, nil
}
