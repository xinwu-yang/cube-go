// @APIVersion 1.0.0
// @Title Cube API
// @Description Cube Project API
// @Contact xinwuy@icloud.com
package routers

import (
	systemApi "com/tievd/cube/modules/system/api"
	"github.com/astaxie/beego"
)

func init() {
	//定义路由
	ns := beego.NewNamespace("/v1",
		systemApi.Mapping,
	)
	beego.AddNamespace(ns)
}
