package main

import (
	_ "com/tievd/cube/config"
	_ "com/tievd/cube/router"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
