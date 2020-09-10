package api

import (
	"com/tievd/cube/config"
	"com/tievd/cube/entity"
	"com/tievd/cube/model/result"
	"com/tievd/cube/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/mojocn/base64Captcha"
	"strconv"
	"time"
)

type AccountApi struct {
	beego.Controller
}

var store = base64Captcha.DefaultMemStore

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
	if !store.Verify(registerUser.CheckKey, registerUser.Captcha, true) {
		this.Data["json"] = result.Err("验证码错误！")
		this.ServeJSON()
		return
	}
	salt := utils.RandomSalt(8)
	password := utils.PasswordEncrypt(registerUser.Username, registerUser.Password, salt)
	var sysUser entity.SysUser
	sysUser.Username = registerUser.Username
	sysUser.RealName = registerUser.Username
	sysUser.Phone = registerUser.Phone
	sysUser.Password = password
	sysUser.Salt = salt
	_, err = orm.NewOrm().Insert(&sysUser)
	if err != nil {
		this.Data["json"] = result.Err(err.Error())
		this.ServeJSON()
		return
	}
	this.Data["json"] = result.Ok("注册成功！")
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
	if !store.Verify(loginUser.CheckKey, loginUser.Captcha, true) {
		this.Data["json"] = result.Err("验证码错误！")
		this.ServeJSON()
		return
	}
	o := orm.NewOrm()
	sysUser := entity.SysUser{Username: loginUser.Username}
	err = o.Read(&sysUser, "Username")
	if err == orm.ErrNoRows {
		this.Data["json"] = result.Err("用户名或密码错误！")
		this.ServeJSON()
		return
	}
	inputPassword := utils.PasswordEncrypt(loginUser.Username, loginUser.Password, sysUser.Salt)
	if err == orm.ErrNoRows || sysUser.Password != inputPassword {
		this.Data["json"] = result.Err("用户名或密码错误！")
		this.ServeJSON()
		return
	}
	token := utils.JwtSign(sysUser.Username, sysUser.Password)
	err = config.Cache.Put("prefix_user_token_"+strconv.FormatInt(sysUser.Id, 10), token, time.Hour)
	if err != nil {
		this.Data["json"] = result.Err(err.Error())
		this.ServeJSON()
		return
	}
	r := make(map[string]interface{})
	r["userInfo"] = sysUser
	r["token"] = token
	this.Data["json"] = result.OkWithBody("登录成功！", r)
	this.ServeJSON()
}

func (this *AccountApi) RandomCaptcha() {
	codeType := this.Ctx.Input.Param(":type")
	var driver base64Captcha.Driver
	if codeType == "string" {
		randColor := base64Captcha.RandColor()
		driverString := base64Captcha.DriverString{
			Width:           105,
			Height:          35,
			Length:          4,
			Source:          base64Captcha.RandomId(),
			ShowLineOptions: 5,
			BgColor:         &randColor,
		}
		driver = driverString.ConvertFonts()
	} else if codeType == "audio" {
		driver = &base64Captcha.DriverAudio{
			Length:   4,
			Language: "zh",
		}
	}
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		this.Data["json"] = result.Err(err.Error())
		this.ServeJSON()
		return
	}
	logs.Info(id)
	logs.Info(store.Get(id, false))
	r := make(map[string]interface{})
	r["captchaId"] = id
	r["b64s"] = b64s
	this.Data["json"] = result.OkWithBody("生成成功！", r)
	this.ServeJSON()
}
