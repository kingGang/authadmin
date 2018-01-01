package controllers

import "lingo8.cn/authadmin/models"

type UserController struct {
	BaseController
}

func (user *UserController) Index() {
	user.Layout = "starter.html"
	user.TplName = "user/index.html"
	user.Data["userlist"] = models.GetUserList()
	// user.LayoutSections = make(map[string]string)
	// user.LayoutSections["headcssjs"] = "user/headercssjs.html"
	// user.LayoutSections["footerjs"] = "user/footerjs.html"
}
