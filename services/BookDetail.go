package services

import (
  "github.com/astaxie/beego"
  "gopkg.in/mgo.v2/bson"
  "Goat/models"
  "errors"
)

var BookDetailCollection = beego.AppConfig.String("BookDetailCollection")

func SearchBookdetail(query string) (err error, rtv models.BookDetail) {
  if CheckAndReconnect() != nil {
    return
  }

  var criteria = bson.M{"title": query}
  err = Session.DB(DB).C(BookDetailCollection).Find(criteria).One(&rtv)
  if err != nil {
    beego.Info(err)
    err = errors.New("Server Internal Error")
    return
  }

  return
}
