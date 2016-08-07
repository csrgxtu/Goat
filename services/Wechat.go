package services

import (
  "github.com/parnurzeal/gorequest"
  "github.com/astaxie/beego"
  // "gopkg.in/mgo.v2/bson"
  "encoding/json"
  "Goat/models"
  "crypto/sha1"
  "strconv"
  "strings"
  "errors"
  "fmt"
  "io"
)

// get wechat api token
func GetAPIToken() (err error, token string) {
  var url = beego.AppConfig.String("Wechat_API_Token")

  request := gorequest.New()
  _, body, errs := request.Get(url).End()
  if len(errs) != 0 {
    beego.Info(errs)
    err = errors.New("Server Internal Error")
    return
  }

  var tokenData models.API_Token
  err = json.Unmarshal([]byte(body), &tokenData)
  if err != nil {
    beego.Info(err)
    err = errors.New("Server Internal Error")
    return
  }

  token = tokenData.ACCESS_TOKEN

  return
}

// 根据ｃｏｄｅ获取ｔｏｋｅｎ
func WechatWebAuthGetAccessToken(code string) (err error, rtv models.WebAuth_AcessToken) {
  var tokenUrl = beego.AppConfig.String("Wechat_WebAuth_AcessToken") + code

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


// jssdk　signature
func GetTicket(token string) (err error, rtv models.JSSDK_Ticket) {
  request := gorequest.New()
  var ticketApi = beego.AppConfig.String("Wechat_JSSDK_Ticket")

  var ticketUrl = ticketApi + "access_token=" + token
  _, body, errs := request.Get(ticketUrl).End()
  if len(errs) != 0 {
    beego.Info(errs)
    err = errors.New("Server Internal Error")
    return
  }

  var jsonData models.JSSDK_Ticket
  err = json.Unmarshal([]byte(body), &jsonData)
  if err != nil {
    beego.Info(err)
    err = errors.New("Server Internal Error")
    return
  }

  if jsonData.Ticket == "" {
    err = errors.New("Server Internal Error")
    return
  }
  // ticket = jsonData.Ticket
  rtv = jsonData

  return
}

// jssdk signature
func GetSignature(noncestr, jsapi_ticket, url string, timestamp int64) (err error, signature string) {
  var hashStr = "jsapi_ticket=" + jsapi_ticket + "&noncestr=" + noncestr + "&timestamp=" + strconv.FormatInt(timestamp, 10) + "&url=" + url

  h := sha1.New()
  io.WriteString(h, hashStr)
  signature = strings.Replace(fmt.Sprintf("% x", h.Sum(nil)), " ", "", -1)

  return
}
