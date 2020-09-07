package entity

import "time"

type SysUser struct {
	Id           int64
	Username     string
	RealName     string
	Password     string    `json:"-"`
	Salt         string    `json:"-"`
	Avatar       string    `orm:"null"`
	Birthday     time.Time `orm:"null"`
	Sex          int       `orm:"null"`
	Email        string    `orm:"null"`
	Phone        string    `orm:"null"`
	OrgCode      string    `orm:"null"`
	OrgCodeTxt   string    `orm:"null"`
	State        int       `orm:"default(0)" description:"状态：0 正常 1 冻结"`
	DelFlag      int       `orm:"default(0)" description:"逻辑删除：0 正常 1 删除"`
	WorkNo       string    `orm:"null"`
	Post         string    `orm:"null"`
	Telephone    string    `orm:"null"`
	CreateBy     string    `orm:"null"`
	CreateTime   time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateBy     string    `orm:"null"`
	UpdateTime   time.Time `orm:"auto_now;type(datetime)"`
	UserIdentity int       `orm:"null"`
	DepartIds    string    `orm:"null"`
	ThirdId      string    `orm:"null"`
	ThirdType    string    `orm:"null"`
	RelTenantIds string    `orm:"null"`
}

type LoginUser struct {
	Username string
	Password string
	Captcha  string
	CheckKey string
}

type RegisterUser struct {
	LoginUser
	Phone string
}
