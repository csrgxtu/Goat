package models

import (
	"gopkg.in/mgo.v2/bson"
)

type (
  // 网页授权登陆根据code获取的access token结构体
  WebAuth_AcessToken struct {
    ACCESS_TOKEN string `json:"access_token"`
    EXPIERS_IN uint64 `json:"expires_in"`
    REFRESH_TOKEN string `json:"refresh_token"`
    OPENID string `json:"openid"`
    SCOPE string `json:"scope"`
    UNIONID string `json:"unionid"`
  }

  // 微信错误返回
  Error struct {
    ERRCODE uint64 `json:"errcode"`
    ERRMSG string `json:"errmsg"`
  }

  // 微信授权的用户
	WebAuth_WechatUser struct {
		Id bson.ObjectId `json:"id"`
		OpenId string `json:"openid"`
		UserName string `json:"user_name"`
    Sex int `json:"sex"`
    Province string `json:"province"`
    City string `json:"city"`
    Country string `json:"country"`
		Avatar string `json:"headimgurl"`
		Language string `json:"language"`
	}
)
