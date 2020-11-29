package boot

import (
	"gf-app/middleware"
	_ "gf-app/packed"
	"gf-app/utils/gtime"
	"gf-app/utils/gtoken"
	"github.com/gogf/gf/frame/g"
)

func init() {
	InitConfig()
	InitModules()
}

func InitConfig()  {
	s := g.Server()

	//response拦截器
	s.Use(middleware.MiddlewareErrorHandler)

	//跨域
	s.Use(middleware.MiddlewareCORS)

	//开启Gtoken鉴权
	gtoken.Tokenizer.Start()

	//admin平滑重启
	s.EnableAdmin()

	//定时任务
	gtime.Start()
}


