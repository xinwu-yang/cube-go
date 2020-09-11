package api

import (
	"com/tievd/cube/common/result"
	"com/tievd/cube/modules/system/entity"
	accountService "com/tievd/cube/modules/system/service"
	"encoding/json"
	"github.com/astaxie/beego"
)

var Mapping beego.LinkNamespace

func init() {
	Mapping = beego.NSNamespace("/sys",
		beego.NSRouter("/randomCaptcha/:type", &AccountApi{}, "get:RandomCaptcha"),
		beego.NSRouter("/login", &AccountApi{}, "post:Login"),
		beego.NSRouter("/register", &AccountApi{}, "post:Register"),
	)
}

type AccountApi struct {
	beego.Controller
}

func (this *AccountApi) Register() {
	var registerUser entity.RegisterUser
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &registerUser)
	if err != nil {
		this.Data["json"] = result.Err(err.Error())
		this.ServeJSON()
		return
	}
	if registerUser.Captcha == "" || registerUser.CheckKey == "" {
		this.Data["json"] = result.Err("验证码无效！")
		this.ServeJSON()
		return
	}
	this.Data["json"] = accountService.Register(registerUser)
	this.ServeJSON()
}

func (this *AccountApi) Login() {
	var loginUser entity.LoginUser
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &loginUser)
	if err != nil {
		this.Data["json"] = result.Err(err.Error())
		this.ServeJSON()
		return
	}
	if loginUser.Captcha == "" || loginUser.CheckKey == "" {
		this.Data["json"] = result.Err("验证码无效！")
		this.ServeJSON()
		return
	}
	this.Data["json"] = accountService.Login(loginUser)
	this.ServeJSON()
}

func (this *AccountApi) RandomCaptcha() {
	this.Data["json"] = accountService.RandomCaptcha(this.Ctx.Input.Param(":type"))
	this.ServeJSON()
}
