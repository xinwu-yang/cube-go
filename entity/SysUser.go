package entity

import "time"

type SysUser struct {
	Id           int64
	Username     string
	RealName     string
	Password     string
	Salt         string
	Avatar       string
	Birthday     time.Time
	Sex          int
	Email        string
	Phone        string
	OrgCode      string
	OrgCodeTxt   string
	State        int
	DelFlag      int
	WorkNo       string
	Post         string
	Telephone    string
	CreateBy     string
	CreateTime   time.Time
	UpdateBy     string
	UpdateTime   time.Time
	UserIdentity int
	DepartIds    string
	ThirdId      string
	ThirdType    string
	RelTenantIds string
}

type SysLoginModel struct {
	Username string
	Password string
	Captcha  string
	CheckKey string
}
