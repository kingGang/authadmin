package controllers

import (
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"lingo8.cn/authadmin/models"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) JsonResult(code models.JRResultCode, msg string, obj interface{}) {
	c.Data["json"] = &models.JsonResult{code, msg, obj}
	c.ServeJSON()
	c.StopRun()
}

func (c *BaseController) PageShow(index, size, count int) string {
	// 	<nav aria-label="...">
	//   <ul class="pagination">
	//     <li class="disabled"><a href="#" aria-label="Previous"><span aria-hidden="true">&laquo;</span></a></li>
	//     <li class="active"><a href="#">1 <span class="sr-only">(current)</span></a></li>
	//     ...
	//   </ul>
	// </nav>
	var num int
	if num = count % size; num > 0 {
		num = count/size + 1
	} else {
		num = count / size
	}
	html := "<nav aria-label='...'><ul class='pagination'>"
	c.Input().Set("page", "15")

	fmt.Printf("%#v", c.Input().Encode())
	html += "<li><a href='#' aria-label='Previous'><span aria-hidden='true'>«</span></a></li>"
	for i := 1; i <= num; i++ {
		if i == index {
			html += "<li class='active'><a href='#'>" + strconv.Itoa(i) + "<span class='sr-only'>(current)</span></a></li>"
		} else {
			c.Input().Set("page", strconv.Itoa(i))
			html += "<li><a href='" + c.Ctx.Request.URL.Path + "?" + c.Input().Encode() + "' aria-label='Previous'><span aria-hidden='true'>" + strconv.Itoa(i) + "</span></a></li>"
		}
	}
	html += "<li><a href='#' aria-label='Next'><span aria-hidden='true'>»</span></a></li>"
	html += "</ul></nav>"
	return html
}
