package controllers

import (
	"fmt"

	"lingo8.cn/authadmin/models"
)

type HomeController struct {
	BaseController
}

func (self *HomeController) Index() {

	fmt.Println(self.GetControllerAndAction())
	self.TplName = "starter.html"
}
func (self *HomeController) Test() {
	self.TplName = "index.tpl"
	userInfo := &models.User{Email: "admin", RealName: "wxg", Password: "123456", IsSuper: true, Status: 1}
	models.InsertUserInfo(userInfo)
}
