package wechat

import (
	"fmt"
	"testing"
)

func TestNewWxConfig(t *testing.T) {
	wx := NewWxConfig("123456", "225566", "2536")
	fmt.Println(*wx, wx)
}

func TestGetSystemConfigFromJson(t *testing.T) {
	s := GetSystemConfigFromJson()
	fmt.Printf("%T", s)
	fmt.Println(s.system.Urls[2].Query[0])
}
