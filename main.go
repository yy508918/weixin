package main

import (
	"net/http"

	"github.com/rixingyike/wechat" // 微信SDK包
)

func main() {
	wechat.Debug = true

	cfg := &wechat.WxConfig{
		Token:          "master",
		AppId:          "wx843991a75ff6c743_123",
		Secret:         "2fd0d0a8ae95227f9bf584214622b134",
		EncodingAESKey: "",
	}

	app := wechat.New(cfg)
	app.SendText("@all", "Hello,World!")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		app.VerifyURL(w, r).NewText("客服消息1").Send().NewText("客服消息2").Send().NewText("查询OK").Reply()
	})

	http.ListenAndServe(":9090", nil)
}
