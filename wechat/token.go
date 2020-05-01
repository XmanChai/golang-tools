package wechat

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type AccessTokenBuilder struct {
	Wxconfig  wxConfig
	RawUrl    *url.URL
	RawResult map[string]interface{}
	Token     AccessToken
}

type AccessToken struct {
	Token     string
	Timestamp string
}

func (akb *AccessTokenBuilder) SetConfigs() BuildProcess {
	//获取配置appid secret grant_type 拼接请求rawquery
	uri := GetSystemConfigFromJson().GetConfigs().Uri["access_token"]

	u, err := url.Parse(uri)
	if err != nil {
		panic(err)
	}

	parse := &url.Values{}
	parse.Set("appid", akb.Wxconfig.appid)
	parse.Set("secret", akb.Wxconfig.secret)
	parse.Set("grant_type", "client_credential")
	u.RawQuery = parse.Encode()
	akb.RawUrl = u
	return akb
}

func (akb *AccessTokenBuilder) SetRequest() BuildProcess {
	//根据获取的HTTP请求地址
	res, err := http.Get(akb.RawUrl.String())
	if res != nil && res.Body != nil {
		defer res.Body.Close()
	}
	if err != nil {
		panic("获取TOKEN失败")
	}
	jMap := make(map[string]interface{})
	err = json.NewDecoder(res.Body).Decode(&jMap)
	if err != nil {
		panic("解析TOKEN失败")
	}
	akb.RawResult = jMap
	return akb
}

func (akb *AccessTokenBuilder) SetResult() BuildProcess {
	//根据返回的MAP判断是否成功
	//根据缓存的类型 把生成的TOKEN存放在文件或redis中
	result := akb.RawResult
	if result["errcode"] == nil || result["errcode"] == 0 {
		akb.Token.Token = result["access_token"].(string)
		akb.Token.Timestamp = time.Now().String()
	} else {
		akb.Token.Token = ""
		akb.Token.Timestamp = time.Now().String()
	}
	return akb
}

func (akb *AccessTokenBuilder) GetResult() interface{} {
	return akb.Token
}
