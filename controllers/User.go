package controllers

import (
	"github.com/astaxie/beego"
  "Goat/models"
  "Goat/services"
)

type UserController struct {
	beego.Controller
}

/**
 * 根据id获取微信用户信息
 */
func (this *UserController) GetUserInfoById() {
	var rt models.Result
	var id = this.GetString(":id")

	err, rtv := services.GetUserInfoById(id)
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

/**
 * 根据微信用户id获取3个相似用户，若不足，则用美丽阅读的用户作为默认
 */
func (this *UserController) GetSimiliar() {
  var rt models.Result
	var id = this.GetString(":id")

  err, rtv := services.GetSimiliar(id)
  if err != nil {
    rt.Msg = "o_o"
    beego.Info(err)
    this.Ctx.ResponseWriter.WriteHeader(500)
  } else {
    rt.Msg = "^_^"
    rt.Data = make([]models.Recs, len(rtv))
    for ix, value := range rtv {
      rt.Data[ix] = value
    }
  }

  this.Data["json"] = &rt
  this.ServeJSON()
}
