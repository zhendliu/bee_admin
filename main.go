package main

import (
	_ "bee_admin/routers"
	_ "bee_admin/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
