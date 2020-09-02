package api

import (
	"com/tievd/cube/config"
	"com/tievd/cube/entity"
	"com/tievd/cube/model/result"
	"com/tievd/cube/utils"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type AccountApi struct {
	beego.Controller
}

func (this *AccountApi) Login() {
	var loginUser entity.SysLoginModel
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &loginUser)
	if err != nil {
		this.Data["json"] = result.Err(err.Error())
	} else {
		if loginUser.Captcha == "" {
			this.Data["json"] = result.Err("验证码无效！")
			this.ServeJSON()
			return
		}
		lowerCaseCaptcha := strings.ToLower(loginUser.Captcha)
		realKey := utils.MD5Encode(lowerCaseCaptcha + loginUser.CheckKey)
		checkCode := config.Cache.Get(realKey)
		if checkCode == nil || checkCode != lowerCaseCaptcha {
			this.Data["json"] = result.Err("验证码错误！")
			this.ServeJSON()
			return
		}
		o := orm.NewOrm()
		sysUser := entity.SysUser{Username: loginUser.Username}
		err := o.Read(&sysUser)
		if err != nil {
			this.Data["json"] = result.Err(err.Error())
			this.ServeJSON()
			return
		}
		inputPassword := utils.PasswordEncrypt(loginUser.Username, loginUser.Password, sysUser.Salt)
		if sysUser.Password != inputPassword {
			this.Data["json"] = result.Err("用户名或密码错误！")
			this.ServeJSON()
			return
		}
		token := utils.JwtSign(sysUser.Username, sysUser.Password)
		_ = config.Cache.Put("prefix_user_token_"+token, token, time.Duration(3600))
		var ok = result.Ok("登录成功！")
		r := make(map[string]interface{})
		r["userInfo"] = sysUser
		r["token"] = token
		ok.Result = &r
		this.Data["json"] = &ok
	}
	this.ServeJSON()
}
