package services

import (
  "github.com/astaxie/beego"
  "gopkg.in/mgo.v2/bson"
  "Goat/models"
  "errors"
  // "strings"
)

var WechatUsersCollection = beego.AppConfig.String("WechatUsersCollection")

func GetUserInfoById(id string) (err error, rtv models.WechatUsers) {
  if CheckAndReconnect() != nil {
    return
  }

  var criteria = bson.M{"_id": bson.ObjectIdHex(id)}
  err = Session.DB(DB).C(WechatUsersCollection).Find(criteria).One(&rtv)
  if err != nil {
    beego.Info(err)
    err = errors.New("Server Internal Error")
    return
  }

  return
}

func GetSimiliar(id string) (err error, rtv []models.WechatUsers) {
  if CheckAndReconnect() != nil {
    return
  }

  // do the similiar calculation, if cant find any, use the default


  return
}
