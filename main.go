package main

import (
	"github.com/astaxie/beego"
	_ "lingo8.cn/authadmin/common"
	_ "lingo8.cn/authadmin/routers"
)

func main() {
	beego.Run()
}
