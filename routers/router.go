package routers

import (
	"github.com/astaxie/beego"
	"lingo8.cn/authadmin/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{}, "*:Index")
	beego.Router("/test", &controllers.HomeController{}, "GET:Test")
	beego.Router("/user", &controllers.UserController{}, "GET:Index")
}
