package util

import (
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniconfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/openplatform"
	openconfig "github.com/silenceper/wechat/v2/openplatform/config"
)

func Wechat() *officialaccount.OfficialAccount {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &offConfig.Config{
		AppID:     "xxx",
		AppSecret: "xxx",
		Token:     "xxx",
		//EncodingAESKey: "xxxx",
		Cache: memory,
	}
	return wc.GetOfficialAccount(cfg)
}

func mini() *miniprogram.MiniProgram {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &miniconfig.Config{
		AppID:     "xxx",
		AppSecret: "xxx",
		//EncodingAESKey: "xxxx",
		Cache: memory,
	}
	return wc.GetMiniProgram(cfg)
}

func open() *openplatform.OpenPlatform {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()
	cfg := &openconfig.Config{
		AppID:     "xxx",
		AppSecret: "xxx",
		Token:     "xxx",
		//EncodingAESKey: "xxxx",
		Cache: memory,
	}
	return wc.GetOpenPlatform(cfg)
}
