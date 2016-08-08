package services

import (
  "github.com/astaxie/beego"
  "gopkg.in/mgo.v2/bson"
  "Goat/models"
  "errors"
  "strings"
)

var BookDetailCollection = beego.AppConfig.String("BookDetailCollection")
var ClassificationCollection = beego.AppConfig.String("ClassificationCollection")

func SearchBookdetail(query string) (err error, rtv models.BookDetail, rtvc models.Classification) {
  if CheckAndReconnect() != nil {
    return
  }

  var criteria = bson.M{"title": query}
  err = Session.DB(DB).C(BookDetailCollection).Find(criteria).One(&rtv)
  if err != nil {
    criteria = bson.M{"title": bson.M{"$regex": bson.RegEx{".*" + query + "*.", ""}}}
    err = Session.DB(DB).C(BookDetailCollection).Find(criteria).One(&rtv)
    if err != nil {
      beego.Info(err)
      err = errors.New("Server Internal Error")
      // return
    }
  }
  // 是在搜不到，就按照V2来搜
  criteria = bson.M{"clc_sort_num": bson.M{"$regex": bson.RegEx{".*V2*.", ""}}}
  err = Session.DB(DB).C(BookDetailCollection).Find(criteria).One(&rtv)
  if err != nil {
    beego.Info(err)
    err = errors.New("Server Internal Error")
    return
  }

  // 在classification里面寻找主文案，标签云
  err, rtvc = GetClcInfo(rtv.ClcSortNum)

  return
}

func GetClcInfo(clc string) (err error, rtv models.Classification) {
  if CheckAndReconnect() != nil {
    return
  }

  // if strings.Index(clc, ".") > 0 {
  //   clc = clc[0:strings.Index(clc, ".")]
  // }

  for i := 0; i < len(clc); i++ {
    // beego.Info(clc[0:len(clc) - i])
    var criteria = bson.M{"clc_sort_num": clc[0:len(clc) - i]}
    err = Session.DB(DB).C(ClassificationCollection).Find(criteria).One(&rtv)
    if err == nil {
      break
    }
  }

  return
}
