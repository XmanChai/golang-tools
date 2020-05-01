package main

import (
	"fmt"
	"xmctools/wechat"
)

func main() {
	//configs := wechat.GetSystemConfigFromJson().GetConfigs()
	//fmt.Printf("type of %v\n %T\n", configs.Urls[2].Query[0], configs)
	wxc := wechat.NewWxConfig("wxd754a949bbcf1130", "dd9476b135e469184a62df91fb30f78b", "")
	configs := wechat.GetSystemConfigFromJson().GetConfigs().Uri
	fmt.Println(configs["access_token"])
	builderDirector := wechat.BuildDirector{}
	accessTokenBuilder := &wechat.AccessTokenBuilder{Wxconfig: *wxc}
	builderDirector.SetBuilder(accessTokenBuilder)
	builderDirector.Constructor()
	token := accessTokenBuilder.GetResult().(wechat.AccessToken)
	fmt.Println(token.Timestamp, token.Token)
	//fmt.Printf("token is %v", token)

}
