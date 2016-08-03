package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	BookDetail struct {
		Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
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
	}
)
