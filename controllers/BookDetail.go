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
 * 根据用户id搜索相关的scan image 或者获取某一个用户的ScanImage
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
