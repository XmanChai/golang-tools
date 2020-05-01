package wechat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//TODO 使用单例模式创建系统配置并加载默认不可更改选项
var (
	systemSettingInstance *systemSingleton
)

type systemConfig struct {
	AccountCapacity int               `json:"accountc_apacity"` //最大公众号容量
	Version         string            `json:"version"`          //系统版本
	Addr            string            `json:"addr"`             //服务地址
	Port            int               `json:"port"`             //http服务端口
	Framework       string            `json:"framework"`        //http框架 [gin/beego/custom]
	Cache           string            `json:"cache"`            //SDK缓存 [redis]
	Storage         string            `json:"storage"`          //token ticket等缓存的方式 [file/redis]
	Redis           redisConfig       `json:"redis"`            //redis配置
	Mysql           mysqlConfig       `json:"mysql"`            //mysql配置
	Uri             map[string]string `json:"uri"`              //微信API接口URI
	Wechats         []wxConfig        `json:"wechats"`
}

type redisConfig struct {
	Addr           string `json:"addr"`
	Port           int    `json:"port"`
	Password       string `json:"password"`
	ConnectTimeout int    `json:"connect_timeout"`
	ConnectRetry   int    `json:"connect_retry"`
	SyncTimeout    int    `json:"sync_timeout"`
}

func (r redisConfig) MakeDbString() string {
	return ""
}

type mysqlConfig struct {
	Addr     string `json:"addr"`
	Port     int    `json:"port"`
	Schema   string `json:"schema"`
	Charset  string `json:"charset"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type systemSingleton struct {
	system systemConfig
}

func GetSystemConfigFromJson() *systemSingleton {
	if systemSettingInstance == nil {
		systemSettingInstance = new(systemSingleton)
		err := systemSettingInstance.loadFromJson()
		if err != nil {
			fmt.Println("读取配置文件失败，SDK无法提供服务")
			os.Exit(1)
		}
	}
	return systemSettingInstance
}

func (s *systemSingleton) GetConfigs() systemConfig {
	return s.system
}

//MEMO 读取JSON文件格式配置文件
func (s *systemSingleton) loadFromJson() error {
	configMap := systemConfig{}
	pwd, _ := os.Getwd()
	data, err := ioutil.ReadFile(pwd + "/config.json")
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &configMap)
	if err != nil {
		return err
	}
	s.system = configMap
	return nil
}

//MEMO 简单工厂模式返回微信配置
type wxConfig struct {
	appid  string
	secret string
	token  string
}

func NewWxConfig(appid, secret, token string) *wxConfig {
	wxc := new(wxConfig)
	wxc.appid = appid
	wxc.secret = secret
	wxc.token = token
	return wxc
}
