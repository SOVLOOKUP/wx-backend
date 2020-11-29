package middleware

import (
	"gf-app/utils/resp"
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

//请求拦截器
func MiddlewareErrorHandler(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.Status == http.StatusNotFound {
		r.Response.ClearBuffer()
		r.Response.WriteJsonExit(resp.NotFound("NoDataFound",""))
	}
}

//鉴权登录
func Login(r *ghttp.Request) (string, interface{}) {
	username := r.GetString("username")
	passwd := r.GetString("passwd")

	//仅供测试
	username = "sss"
	userdata := passwd
	// TODO 进行登录校验

	return username, userdata
}

//跨域
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}