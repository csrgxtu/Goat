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
		Tags []string `json:"tags" bson:"tags"`
	}

	WechatUsers struct {
		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
		OpenId string `json:"openid" bson:"openid"`
		NickName string `json:"nickname" bson:"nickname"`
		Sex bool `json:"sex" bson:"sex"`
		Language string `json:"language" bson:"language"`
		City string `json:"city" bson:"city"`
		Province string `json:"province" bson:"province"`
		Country string `json:"country" bson:"country"`
		HeadImgUrl string `json:"headimgurl" bson:"headimgurl"`
		BookDetailIds []string `json:"bookdetailids" bson:"bookdetailids"`
	}
)
