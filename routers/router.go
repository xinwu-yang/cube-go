// @APIVersion 1.0.0
// @Title Cube API
// @Description Cube Project API
// @Contact xinwuy@icloud.com
package routers

import (
	"com/tievd/cube/api"
	"github.com/astaxie/beego"
)

func init() {
	//定义路由
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/account",
			beego.NSRouter("/randomCaptcha/:type", &api.AccountApi{}, "get:RandomCaptcha"),
			beego.NSRouter("/login", &api.AccountApi{}, "post:Login"),
			beego.NSRouter("/register", &api.AccountApi{}, "post:Register"),
		),
	)
	beego.AddNamespace(ns)
}
