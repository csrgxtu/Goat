package services

import (
  "github.com/astaxie/beego"
  "gopkg.in/mgo.v2/bson"
  "Goat/models"
  "errors"
  "sort"
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

// 如果用户搜索命中一本书，则添加到他的相关书籍里面
func AppendBookDetailId(bid, userid string) (err error) {
  if CheckAndReconnect() != nil {
    return
  }

  var wechatUsers models.WechatUsers
  var criteria = bson.M{"_id": bson.ObjectIdHex(userid)}
  err = Session.DB(DB).C(WechatUsersCollection).Find(criteria).One(&wechatUsers)
  if err != nil {
    beego.Info(err)
    err = errors.New("Server Internal Error")
    return
  }

  bids := wechatUsers.BookDetailIds
  bids[len(bids)] = bid
  var change = bson.M{"$set": bson.M{"bookdetailids": bids}}
  err = Session.DB(DB).C(WechatUsersCollection).Update(criteria, change)
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

  // 首先获取当前微信用户的信息及藏书
  var wechatUser models.WechatUsers
  var criteria = bson.M{"_id": bson.ObjectIdHex(id)}
  err = Session.DB(DB).C(WechatUsersCollection).Find(criteria).One(&wechatUser)
  if err != nil {
    // 妈的，这个用户不存在，不用计算相似度了
    beego.Info(err)
    err = errors.New("Server Internal Error")
    return
  }

  // 查询共同书籍用户
  var wechatUsers []models.WechatUsers
  criteria = bson.M{"bookdetailids": bson.M{"$in": wechatUser.BookDetailIds}}
  err = Session.DB(DB).C(WechatUsersCollection).Find(criteria).All(&wechatUsers)
  if err != nil {
    // 找不到
    beego.Info(err)
    err = errors.New("Server Internal Error")
    return
  }

  // 找共同书籍数目, 并存储
  tmpStruct := make(PairList, len(wechatUsers))
  for ix, value := range wechatUsers {
    tmp := IntersectString(wechatUser.BookDetailIds, value.BookDetailIds)
    if len(tmp) == 0 {
      continue
    }
    tmpStruct[ix].Key = value.Id.Hex()
    tmpStruct[ix].Value = len(tmp)
  }
  // 对上述结果进行排序，取出前三,  否则去找默认用户，下面的代码之所以这么长，是因为产品的奇怪需求
  sort.Sort(tmpStruct)
  if len(tmpStruct) == 0 {
    // 取三个异性用户
  } else if len(tmpStruct) == 1 {
    // 取两个异性用户
  } else if len(tmpStruct) == 2 {
    // 取一个异性用户
  } else if len(tmpStruct) >= 3 {
    // 返回这三个用户
    for i := 0; i < 3; i++ {
      rtv[i].Id = bson.ObjectIdHex(tmpStruct[i].Key)
      rtv[i].Similiraty = float64(tmpStruct[i].Value) / float64(len(wechatUser.BookDetailIds))
    }
  }

  return
}
