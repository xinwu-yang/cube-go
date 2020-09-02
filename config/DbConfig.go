package config

import (
	"com/tievd/cube/entity"
	"github.com/astaxie/beego/orm"
)
import _ "github.com/go-sql-driver/mysql"

func init() {
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	// username:password@protocol(ip:port)/dbname?param=value
	_ = orm.RegisterDataBase("default", "mysql", "root:chengxun@/demo?charset=utf8")

	orm.RegisterModel(new(entity.SysUser))
	// 自动建表
	//_ = orm.RunSyncdb("default", true, true)
}
