package config

import (
	"gf-app/module/base/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func InitRouter()  {
    s := g.Server()
    prefix := "/api/v1/wxrobot"

    s.Group(prefix+"/",func(g *ghttp.RouterGroup) {
    	// 基础消息接收
    	g.ALL("/",api.MessageHandler)

    	// 消息模板发送
		g.ALL("/oatemplate",api.TemplateHandler)
	})
}