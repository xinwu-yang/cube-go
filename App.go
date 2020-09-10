package main

import (
	_ "com/tievd/cube/config"
	_ "com/tievd/cube/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
