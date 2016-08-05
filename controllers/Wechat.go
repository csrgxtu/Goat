package controllers

import (
	"github.com/astaxie/beego"
  "Goat/models"
  "Goat/services"
  "strconv"
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

/**
 * 网页授权
 */
func (this *WechatController) WebAuth() {
  var rt models.Result
  var code = this.GetString("code")

  // 授权失败，没有code，无法继续后面的拉取用户信息操作
  if len(code) == 0 {
    rt.Msg = "o_o"
		this.Ctx.ResponseWriter.WriteHeader(500)
    this.Data["json"] = &rt
    this.ServeJSON()
  }

  // get access_token
  err, rtv := services.WechatWebAuthGetAccessToken(code)
  if err != nil {
    rt.Msg = "o_o"
    beego.Info(err)
    this.Ctx.ResponseWriter.WriteHeader(500)
    this.Data["json"] = &rt
    this.ServeJSON()
  }

  // get wechat user Info
  erra, rtva := services.WechatWebAuthGetWechatUserInfo(rtv.ACCESS_TOKEN, rtv.OPENID, "zh_CN")
  if erra != nil {
    rt.Msg = "o_o"
    beego.Info(erra)
    this.Ctx.ResponseWriter.WriteHeader(500)
    this.Data["json"] = &rt
    this.ServeJSON()
  }

  // save this to the database
  var user models.WechatUsers
  user.OpenId = rtva.OpenId
  user.UserName = rtva.UserName
  user.Avatar = rtva.Avatar
  user.Sex = rtva.Sex
  user.Language = rtva.Language
  user.Province = rtva.Province
  user.City = rtva.City
  user.Country = rtva.Country
  errb, rtvb := services.CreateWechatUser(user)
  if errb != nil {
    rt.Msg = "o_o"
    beego.Info(errb)
    this.Ctx.ResponseWriter.WriteHeader(500)
    this.Data["json"] = &rt
    this.ServeJSON()
  }

  // Successful, redirect with user id
  this.Redirect(beego.AppConfig.String("Wechat_WebAuth_Redirect") + "#userid=" + rtvb, 302)
}

/**
 * jssdk 获取signature
 */
func (this *WechatController) Signature() {
  var rt models.Result

  // get ticket
  err, rtv := services.GetAPIToken()
  if err != nil {
    rt.Msg = "o_o"
    beego.Info(err)
    this.Ctx.ResponseWriter.WriteHeader(500)
    this.Data["json"] = &rt
    this.ServeJSON()
  }

  erra, rtva := services.GetTicket(rtv)
  if erra != nil {
    rt.Msg = "o_o"
    beego.Info(erra)
    this.Ctx.ResponseWriter.WriteHeader(500)
    this.Data["json"] = &rt
    this.ServeJSON()
  }

  // get signature
  var noncestr = beego.AppConfig.String("Wechat_JSSDK_Noncestr")
  var url = beego.AppConfig.String("Wechat_JSSDK_Url")
  timestamp, err := strconv.ParseInt(beego.AppConfig.String("Wechat_JSSDK_Timestamp"), 10, 64)
  if err != nil {
    rt.Msg = "o_o"
    beego.Info(err)
    this.Ctx.ResponseWriter.WriteHeader(500)
    this.Data["json"] = &rt
    this.ServeJSON()
  }

  errc, rtvc := services.GetSignature(noncestr, rtva, url, int64(timestamp))
  if erra != nil {
    rt.Msg = "o_o"
    beego.Info(errc)
    this.Ctx.ResponseWriter.WriteHeader(500)
    this.Data["json"] = &rt
    this.ServeJSON()
  } else {
    rt.Msg = "^_^"
    var data models.JSSDK_Signature
    data.Nonestr = noncestr
    data.JSAPI_Ticket = rtva
    data.Timestamp = int64(timestamp)
    data.Signature = rtvc
    data.Url = url
    data.AppId = beego.AppConfig.String("Wechat_APPID")
    rt.Data = make([]models.Recs, 1)
    rt.Data[0] = data
  }

  this.Data["json"] = &rt
  this.ServeJSON()
}
