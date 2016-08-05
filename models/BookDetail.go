package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	BookDetail struct {
		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
		WukongDocId uint64 `json:"wukongdocid" bson:"wukongdocid"`
		Isbn string `json:"isbn" bson:"isbn"`
		Title string `json:"title" bson:"title"`
		Author string `json:"author" bson:"author"`
		ClcSortNum string `json:"clc_sort_num" bson:"clc_sort_num"`
		Status string `json:"status" bson:"status"`
		Version time.Time `json:"version" bson:"version"`
		CreateUserId string `json:"createuserid" bson:"createuserid"`
		UpdateUserId string `json:"updateuserid" bson:"updateuserid"`
		CreateTime string `json:"createtime" bson:"createtime"`
		UpdateTime string `json:"updatetime" bson:"updatetime"`
	}

	Classification struct {
		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
		ClcSortNum string `json:"clc_sort_num" bson:"clc_sort_num"`
		Description string `json:"description" bson:"description"`
		Main string `json:"main" bson:"main"`
		Tags []string `json:"tags" bson:"tags"`
	}

	// 微信授权的用户
	WechatUsers struct {
		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
		OpenId string `json:"openid" bson:"openid"`
		UserId string `json:"user_id" bson:"user_id"`
		UserName string `json:"user_name" bson:"nickname"`
		Avatar string `json:"avatar" bson:"headimgurl"`
		BookDetailIds []string `json:"bookdetailids" bson:"bookdetailids"`
		Similiraty float64 `json:"similiraty" bson:"similiraty"` // 这个用于快速返回相似性的
		Default bool `json:"default" bson:"default"`
		Sex bool `json:"sex" bson:"sex"`
		Language string `json:"language" bson:"language"`
		City string `json:"city" bson:"city"`
		Province string `json:"province" bson:"province"`
		Country string `json:"country" bson:"country"`
	}

	// 美丽阅读的用户
	BRUsers struct {
		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
		UserId string `json:"user_id" bson:"user_id"`
		UserName string `json:"user_name" bson:"nickname"`
		Avatar string `json:"avatar" bson:"headimgurl"`
	}
)
