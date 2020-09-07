package config

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
)

var Cache cache.Cache

func init() {
	newCache, err := cache.NewCache("redis", `{"key":"cube","conn":":6379","dbNum":"1"}`)
	if err != nil {
		logs.Error("Redis初始化失败！")
	}
	Cache = newCache
}
