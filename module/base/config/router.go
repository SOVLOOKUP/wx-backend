package config

import (
	"gf-app/module/base/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func InitRouter()  {
    s := g.Server()
    s.Group("/",func(g *ghttp.RouterGroup) {
    	g.ALL("/",api.Tasks)
	})
}