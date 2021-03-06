package services

import (
  "github.com/astaxie/beego"
  "gopkg.in/mgo.v2/bson"
  "Goat/models"
  "errors"
  // "strings"
)

var BookDetailCollection = beego.AppConfig.String("BookDetailCollection")
var ClassificationCollection = beego.AppConfig.String("ClassificationCollection")

func SearchBookdetailDefault(query string) (err error, rtv models.BookDetail, rtvc models.Classification) {
  if CheckAndReconnect() != nil {
    return
  }

  var criteria = bson.M{"clc_sort_num": bson.M{"$ne": nil}}
  err = Session.DB(DB).C(BookDetailCollection).Find(criteria).One(&rtv)
  if err != nil {
    beego.Info(err)
    err = errors.New("Server Internal Error")
    return
  }

  err, rtvc = GetClcInfo("default")
  return
}

func SearchBookdetail(query string) (err error, rtv models.BookDetail, rtvc models.Classification) {
  if CheckAndReconnect() != nil {
    return
  }

  var errb error
  var criteria = bson.M{"title": query}
  erra := Session.DB(DB).C(BookDetailCollection).Find(criteria).One(&rtv)
  if erra != nil {
    beego.Info(erra)
    criteria = bson.M{"title": bson.M{"$regex": bson.RegEx{".*" + query + "*.", ""}}}
    errb = Session.DB(DB).C(BookDetailCollection).Find(criteria).One(&rtv)
    if errb != nil {
      beego.Info(errb)
      errb = errors.New("Server Internal Error")
      // return
    }
  }

  // 是在搜不到，就按照V2来搜
  // if errb != nil {
  //   beego.Info("使用默认")
  //   var rtvs []models.BookDetail
  //   // criteria = bson.M{"clc_sort_num": bson.M{"$regex": bson.RegEx{".*V2*.", ""}}}
  //   criteria = bson.M{"clc_sort_num": "V2"}
  //   // errc := Session.DB(DB).C(BookDetailCollection).Find(criteria).One(&rtv)
  //   errc := Session.DB(DB).C(BookDetailCollection).Find(criteria).Limit(60).All(&rtvs)
  //   if errc != nil {
  //     beego.Info(errc)
  //     errc = errors.New("Server Internal Error")
  //     return
  //   }
  //
  //   rtv = rtvs[int(RangeRandomFloat(0, 60))]
  // }

  // 如果找不到，使用默认的
  DefaultClc := [8]string{"default", "Z2", "Z3", "Q95", "Q94", "I27", "I287.8", "I287.45"}
  if errb != nil {

    err, rtvc = GetClcInfo(DefaultClc[int(RangeRandomFloat(0, 8))])
  } else {
    // 在classification里面寻找主文案，标签云
    if len(rtv.ClcSortNum) > 0 {
      err, rtvc = GetClcInfo(rtv.ClcSortNum)
    } else {
      err, rtvc = GetClcInfo(DefaultClc[int(RangeRandomFloat(0, 8))])
    }
  }

  if err != nil {
    err, rtvc = GetClcInfo(DefaultClc[int(RangeRandomFloat(0, 8))])
  }

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
    // criteria = bson.M{"clc_sort_num": bson.M{"$regex": bson.RegEx{".*V2*.", ""}}}

    // var criteria = bson.M{"clc_sort_num": bson.M{"$regex": bson.RegEx{".*" + clc[0:len(clc) - i] + "*.", ""}}}
    var criteria = bson.M{"clc_sort_num": clc[0:len(clc) - i]}
    err = Session.DB(DB).C(ClassificationCollection).Find(criteria).One(&rtv)
    if err == nil {
      break
    }
  }

  return
}
