package controllers

import (
	"github.com/astaxie/beego"
  // "Goat/models"
  // "Goat/services"
)

type WechatController struct {
	beego.Controller
}

/**
 * 验证微信服务器
 */
func (this *WechatController) Verify() {
  var echostr = this.GetString("echostr")

  this.Ctx.WriteString(echostr)
}
