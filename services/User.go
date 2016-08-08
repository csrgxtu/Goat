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

// save a wechat user
func CreateWechatUser(user models.WechatUsers) (err error, rtv string){
  if CheckAndReconnect() != nil {
    return
  }

  // first, check if wechat user already in
  var wechatUser models.WechatUsers
  var criteria = bson.M{"openid": user.OpenId}
  err = Session.DB(DB).C(WechatUsersCollection).Find(criteria).One(&wechatUser)
  if err == nil {
    rtv = wechatUser.Id.Hex()
    return
  }

  user.Id = bson.NewObjectId()
  err = Session.DB(DB).C(WechatUsersCollection).Insert(user)
  if err != nil {
    beego.Info(err)
    err = errors.New("Server Internal Error")
    return
  }
  rtv = user.Id.Hex()

  return
}

// 如果用户搜索命中一本书，则添加到他的相关书籍里面
func AppendBookDetailId(bid, userid string) (err error) {
  if CheckAndReconnect() != nil {
    return
  }

  var criteria = bson.M{"_id": bson.ObjectIdHex(userid)}
  var change = bson.M{"$push": bson.M{"bookdetailids": bid}}
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
  criteria = bson.M{"bookdetailids": bson.M{"$in": wechatUser.BookDetailIds}, "_id": bson.M{"$ne": bson.ObjectIdHex(id)}}
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
  beego.Info(tmpStruct)
  if len(tmpStruct) == 0 {
    // 取三个异性用户
    beego.Info(wechatUser.Sex)
    criteria = bson.M{"sex": bson.M{"$ne": wechatUser.Sex}, "default": true}
    err = Session.DB(DB).C(WechatUsersCollection).Find(criteria).All(&rtv)
    if err != nil {
      beego.Info(err)
      err = errors.New("Server Internal Error")
      return
    }
    rtv[0].Similiraty = 0.98
    rtv[1].Similiraty = 0.96
    rtv[2].Similiraty = 0.92
    rtv = rtv[0:3]
  } else if len(tmpStruct) == 1 {
    // 取两个异性用户
    criteria = bson.M{"sex": bson.M{"$ne": wechatUser.Sex}, "default": true}
    err = Session.DB(DB).C(WechatUsersCollection).Find(criteria).All(&rtv)
    if err != nil {
      beego.Info(err)
      err = errors.New("Server Internal Error")
      return
    }
    rtv[0].Similiraty = 0.98
    rtv[1].Similiraty = 0.96
    rtv = rtv[0:2]
    var tmpWechatUser models.WechatUsers
    criteria = bson.M{"_id": bson.ObjectIdHex(tmpStruct[0].Key)}
    err = Session.DB(DB).C(WechatUsersCollection).Find(criteria).One(&tmpWechatUser)
    if err != nil {
      beego.Info(err)
      err = errors.New("Server Internal Error")
      return
    }
    tmpWechatUser.Similiraty = float64(tmpStruct[0].Value) / float64(len(wechatUser.BookDetailIds))
    rtv = append(rtv, tmpWechatUser)
  } else if len(tmpStruct) == 2 {
    // 取一个异性用户
    criteria = bson.M{"sex": bson.M{"$ne": wechatUser.Sex}, "default": true}
    err = Session.DB(DB).C(WechatUsersCollection).Find(criteria).All(&rtv)
    if err != nil {
      beego.Info(err)
      err = errors.New("Server Internal Error")
      return
    }
    rtv[0].Similiraty = 0.98
    rtv = rtv[0:1]

    var tmpWechatUsers [2]models.WechatUsers
    err, tmpWechatUsers[0] = GetUserInfoById(tmpStruct[0].Key)
    if err != nil {
      beego.Info(err)
      err = errors.New("Server Internal Error")
      return
    }
    err, tmpWechatUsers[1] = GetUserInfoById(tmpStruct[1].Key)
    if err != nil {
      beego.Info(err)
      err = errors.New("Server Internal Error")
      return
    }
    tmpWechatUsers[0].Similiraty = float64(tmpStruct[0].Value) / float64(len(wechatUser.BookDetailIds))
    tmpWechatUsers[1].Similiraty = float64(tmpStruct[1].Value) / float64(len(wechatUser.BookDetailIds))
    rtv = append(rtv, tmpWechatUsers[1])
    rtv = append(rtv, tmpWechatUsers[0])
  } else if len(tmpStruct) >= 3 {
    // 返回这三个用户
    rtv = make([]models.WechatUsers, 3)
    for i := 0; i < 3; i++ {
      err, rtv[i] = GetUserInfoById(tmpStruct[i].Key)
      rtv[i].Similiraty = float64(tmpStruct[i].Value) / float64(len(wechatUser.BookDetailIds))
    }
  }

  // 上面计算的相似度因为用户书籍等数据太少，所以肖总要求使用80-99之间的随机数替代，不过上面代码我不删除
  for i := 0; i < 3; i++ {
    rtv[i].Similiraty = RangeRandomFloat(80, 99)
  }

  return
}
