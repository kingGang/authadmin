package controllers

import "lingo8.cn/authadmin/models"
import "fmt"

type UserController struct {
	BaseController
}

func (user *UserController) Index() {
	pindex, err := user.GetInt("page")
	if err != nil {
		pindex = 1
	}
	user.Layout = "starter.html"
	user.TplName = "user/index.html"
	psize := 5
	list, count := models.GetUserList(pindex, psize)
	user.Data["userlist"] = list
	user.Data["page"] = user.PageShow(pindex, psize, count)
	user.LayoutSections = make(map[string]string)
	user.LayoutSections["headcssjs"] = "user/headercssjs.html"
	user.LayoutSections["footerjs"] = "user/footerjs.html"
}
func (user *UserController) Edit() {
	if user.Ctx.Request.Method == "POST" {
		uinfo := models.User{}
		err := user.ParseForm(&uinfo)
		if err == nil {
			fmt.Printf("%#v", uinfo)
		}
		id := models.InsertUserInfo(&uinfo)
		if id > 0 {
			user.JsonResult(models.JRCodeSucc, "添加成功", id)
		}
		user.JsonResult(models.JRCodeFailed, "添加失败", id)

		return
	}
	user.TplName = "user/edit.html"
	user.Layout = "base/layout_pullbox.html"
	user.LayoutSections = make(map[string]string)
	user.LayoutSections["footerjs"] = "user/edit_footerjs.html"
}
func (user *UserController) Delete() {
	if user.Ctx.Request.Method != "POST" {
		user.JsonResult(models.JRCode401, "请求方式错误", nil)
	}
	id, err := user.GetInt("id")
	if err != nil {
		user.JsonResult(models.JRCodeFailed, "参数错误", nil)
	}
	num := models.DelUserInfo(id)
	if num > 0 {
		user.JsonResult(models.JRCodeSucc, "删除成功", num)
	}
	user.JsonResult(models.JRCodeFailed, "删除失败", id)
}
