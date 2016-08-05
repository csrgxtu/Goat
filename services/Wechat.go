package services

import (
  "github.com/parnurzeal/gorequest"
  "github.com/astaxie/beego"
  // "gopkg.in/mgo.v2/bson"
  "encoding/json"
  "Goat/models"
  "errors"
)

// 根据ｃｏｄｅ获取ｔｏｋｅｎ
func WechatWebAuthGetAccessToken(code string) (err error, rtv models.WebAuth_AcessToken) {
  var tokenUrl = beego.AppConfig.String("Wechat_WebAuth_Redirect") + code

  request := gorequest.New()
  _, body, errs := request.Get(tokenUrl).End()
  if len(errs) != 0 {
    beego.Info(errs)
    err = errors.New("Server Internal Error")
    return
  }

  var tokenData models.WebAuth_AcessToken
  err = json.Unmarshal([]byte(body), &tokenData)
  if err != nil {
    beego.Info(err)
    err = errors.New("Server Internal Error")
    return
  }

  if tokenData.ACCESS_TOKEN == "" {
    beego.Info("Access_Token is empty")
    err = errors.New("Server Internal Error")
    return
  }

  rtv = tokenData
  return
}

// get wechat auth user info according token and openid
func WechatWebAuthGetWechatUserInfo(token, openid, lang string) (err error, rtv models.WebAuth_WechatUser) {
  var userUrl = beego.AppConfig.String("Wechat_WebAuth_UserInfo") + "access_token=" + token + "&openid=" + openid + "&lang=" + lang
  request := gorequest.New()
  _, body, errs := request.Post(userUrl).End()
  if len(errs) != 0 {
    beego.Info(errs)
    err = errors.New("Server Internal Error")
    return
  }

  err = json.Unmarshal([]byte(body), &rtv)
  if err != nil {
    beego.Info(err)
    err = errors.New("Server Internal Error")
    return
  }

  return
}
