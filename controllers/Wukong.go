package controllers

import (
	"github.com/astaxie/beego"
  "Goat/models"
  "Goat/services"
)

type WukongController struct {
	beego.Controller
}


func (this *WukongController) Indexer() {
	var rt models.Result

	err, rtv := services.Indexer()
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

func (this *WukongController) Searcher() {
  var rt models.Result
  var query = this.GetString(":query")
	var userid = this.GetString(":userid")

  err, rtv, rtvc := services.Searcher(query)
	if err != nil {
		rt.Msg = "o_o"
		beego.Info(err)
		this.Ctx.ResponseWriter.WriteHeader(500)
	} else {
		rt.Msg = "^_^"
		rt.Data = make([]models.Recs, 2)
		rt.Data[0] = rtv
		rt.Data[1] = services.GetTagClouds(rtvc.Tags)
	}
	services.AppendBookDetailId(rtv.Id.Hex(), userid)

  this.Data["json"] = &rt
  this.ServeJSON()
}
