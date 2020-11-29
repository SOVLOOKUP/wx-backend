package config

import (
	"gf-app/module/base/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func InitRouter()  {
    s := g.Server()
    //prefix := "/api/v1/wxrobot"
	prefix := ""

    s.Group(prefix+"/",func(g *ghttp.RouterGroup) {
    	// 基础消息接收
    	//g.ALL("/",api.MessageHandler)

    	// 消息模板发送
		g.ALL("/oatemplate",api.TemplateHandler)

		// 获取用户openid列表
		g.ALL("/alluser",api.ListAllUserOpenIDs)

		// 获取tags列表
		g.ALL("/alltags",api.GetTagList)

		// 获取tag下的用户
		g.ALL("/tagsuser",api.OpenIDListByTag)

		// 获取templatelist
		g.ALL("/templatelist",api.ListTemplate)

		// 获取templatelist
		g.ALL("/clearquota",api.ClearQuota)
	})
}