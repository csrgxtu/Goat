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
		UserName string `json:"nickname"`
    Sex int `json:"sex"`
    Province string `json:"province"`
    City string `json:"city"`
    Country string `json:"country"`
		Avatar string `json:"headimgurl"`
		Language string `json:"language"`
	}

  JSSDK_Ticket struct {
  	Error_Code int    `json:"errorcode"`
  	Error_Msg  string `json:"errmsg"`
  	Ticket     string `json:"ticket"`
  	Expires_In int    `json:"expires_in"`
  }

  API_Token struct {
    ACCESS_TOKEN string `json:"access_token"`
    EXPIRES_IN uint64 `json:"expires_in"`
  }

  JSSDK_Signature struct {
    Nonestr string `json:"noncestr"`
    JSAPI_Ticket string `json:"jsapi_ticket"`
    Timestamp int64 `json:"timestamp"`
    Signature string `json:"signature"`
    Url string `json:"url"`
    AppId string `json:"appid"`
  }
)
