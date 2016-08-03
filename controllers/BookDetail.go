package controllers

import (
	"github.com/astaxie/beego"
  "Goat/models"
  "Goat/services"
)

type BookDetailController struct {
	beego.Controller
}

/**
 * 根据书籍名称获取其信息
 */
func (this *BookDetailController) SearchBookDetail() {
	var rt models.Result
	var query = this.GetString(":query")

	err, rtv := services.SearchBookdetail(query)
	if err != nil {
		rt.Msg = "o_o"
		beego.Info(err)
		this.Ctx.ResponseWriter.WriteHeader(500)
	} else {
		rt.Msg = "^_^"
		rt.Data = make([]models.Recs, 1)
		rt.Data[0] = rtv
	}

	this.Data["json"] = &rt
	this.ServeJSON()
}
