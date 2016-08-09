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
	var userid = this.GetString(":userid")

	err, rtv, rtvc := services.SearchBookdetail(query)
	if err != nil {
		rt.Msg = "o_o"
		beego.Info(err)
		this.Ctx.ResponseWriter.WriteHeader(500)
	} else {
		rt.Msg = "^_^"
		rt.Data = make([]models.Recs, 5)
		rt.Data[0] = rtv
		rt.Data[1] = rtvc.Main
		rt.Data[2] = rtvc.Img
		// 随机选择A类活着B类标签
		if int(services.RangeRandomFloat(0, 2)) == 0 {
			beego.Info("chose a tag")
			rt.Data[3] = services.GetTagClouds(rtvc.TagA)
		} else {
			beego.Info("chose b tag")
			rt.Data[3] = services.GetTagClouds(rtvc.TagB)
		}
		// beego.Info(rtvc.TagA)
		// beego.Info(len(rtvc.TagA))
		// rt.Data[3] = services.GetTagClouds(rtvc.TagA)

		_, rt.Data[4] = services.GetUserInfoById(userid)
	}
	services.AppendBookDetailId(rtv.Id.Hex(), userid)

	this.Data["json"] = &rt
	this.ServeJSON()
}

func (this *BookDetailController) GetClcInfo() {
	var rt models.Result
	var clc = this.GetString(":clc_sort_num")

	err, rtv := services.GetClcInfo(clc)
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
