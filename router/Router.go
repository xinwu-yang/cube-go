package router

import (
	"com/tievd/cube/api"
	"github.com/astaxie/beego"
)

func init() {
	//定义路由
	beego.Router("/login", &api.AccountApi{}, "post:Login")
	beego.Router("/randomImage/:type", &api.AccountApi{}, "get:RandomImage")
}
