package service

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/silenceper/wechat/v2/officialaccount/server"
	"net/http"
)

var cfg = g.Cfg()

type WxRobot struct {
	Wxcfg *offConfig.Config		//公众号配置
	Name string		//实例名称
	CreateTime *gtime.Time	//实例创建时间
	Account *officialaccount.OfficialAccount	//公众号实例
}

func (a *WxRobot) Start(name string) *WxRobot {
	wc := wechat.NewWechat()

	//名称和时间
	a.Name = name
	a.CreateTime = gtime.Now()
	memory := cache.NewMemory()

	//配置实例
	a.Wxcfg = &offConfig.Config{
		AppID:     cfg.GetString("werobot.wxAppId"),
		AppSecret: cfg.GetString("werobot.wxAppSecret"),
		Token:     cfg.GetString("werobot.wxToken"),
		//EncodingAESKey: "xxxx",
		Cache: memory,
	}
	a.Account = wc.GetOfficialAccount(a.Wxcfg)
	glog.Printf(`
WxRobot start success!
name %v
crateTime %v`,a.Name,a.CreateTime)
	return a
}

func (a *WxRobot) GetServer(rw http.ResponseWriter, req *http.Request) *server.Server {
	return a.Account.GetServer(req, rw)
}