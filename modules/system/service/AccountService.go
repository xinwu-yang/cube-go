package service

import (
	"com/tievd/cube/common/result"
	"com/tievd/cube/common/utils"
	"com/tievd/cube/config"
	"com/tievd/cube/modules/system/entity"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/mojocn/base64Captcha"
	"time"
)

var store = base64Captcha.DefaultMemStore

func Register(registerUser entity.RegisterUser) result.Result {
	if !store.Verify(registerUser.CheckKey, registerUser.Captcha, true) {
		return result.Err("验证码错误！")
	}
	o := orm.NewOrm()
	//验证用户名是否重复
	existUser := entity.SysUser{Username: registerUser.Username}
	_ = o.Read(&existUser, "Username")
	if existUser.Id > 0 {
		return result.Err("用户名已存在！")
	}
	existUser.Phone = registerUser.Phone
	_ = o.Read(&existUser, "Phone")
	if existUser.Id > 0 {
		return result.Err("手机号已存在！")
	}
	salt := utils.RandomSalt(8)
	password := utils.PasswordEncrypt(registerUser.Username, registerUser.Password, salt)
	var sysUser entity.SysUser
	sysUser.Username = registerUser.Username
	sysUser.RealName = registerUser.Username
	sysUser.Phone = registerUser.Phone
	sysUser.Password = password
	sysUser.Salt = salt
	_, err := o.Insert(&sysUser)
	if err != nil {
		return result.Err(err.Error())
	}
	return result.Ok("注册成功！")
}

func Login(loginUser entity.LoginUser) result.Result {
	if !store.Verify(loginUser.CheckKey, loginUser.Captcha, true) {
		return result.Err("验证码错误！")
	}
	o := orm.NewOrm()
	sysUser := entity.SysUser{Username: loginUser.Username}
	err := o.Read(&sysUser, "Username")
	if err == orm.ErrNoRows {
		return result.Err("用户名或密码错误！")
	}
	inputPassword := utils.PasswordEncrypt(loginUser.Username, loginUser.Password, sysUser.Salt)
	if err == orm.ErrNoRows || sysUser.Password != inputPassword {
		return result.Err("用户名或密码错误！")
	}
	token := utils.JwtSign(sysUser.Username, sysUser.Password)
	err = config.Cache.Put("prefix_user_token_"+token, token, time.Hour)
	if err != nil {
		return result.Err(err.Error())
	}
	r := make(map[string]interface{})
	r["userInfo"] = sysUser
	r["token"] = token
	return result.OkWithBody("登录成功！", r)
}

func RandomCaptcha(captchaType string) result.Result {
	var driver base64Captcha.Driver
	if captchaType == "string" {
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
	} else if captchaType == "audio" {
		driver = &base64Captcha.DriverAudio{
			Length:   4,
			Language: "zh",
		}
	}
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		return result.Err(err.Error())
	}
	logs.Info(id)
	logs.Info(store.Get(id, false))
	r := make(map[string]interface{})
	r["captchaId"] = id
	r["b64s"] = b64s
	return result.OkWithBody("生成成功！", r)
}
