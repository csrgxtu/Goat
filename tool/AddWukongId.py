#!/usr/local/env python
# coding=utf-8
#
# Author: Archer
# File: AddWukongId.py
# Desc: 因为悟空检索的DocId不支持string格式，只能是int，所以这里像数据库的bookdeail
# 添加int类型的id
# Date: 4/Aug/2016
from pymongo import MongoClient

# connect to mongodb
# client = MongoClient('mongodb://192.168.100.2:27017/bookshelf')
# client = MongoClient('mongodb://rio:VFZPhT7y@192.168.200.22:27017/bookshelf')
client = MongoClient('mongodb://localhost:27017/bookshelf')
db = client['bookshelf']
dc = db['bookdetail']

# foreach doment, update $set
books = dc.find()
WukongDocId = int(0)
for book in books:
    print book['_id'], WukongDocId
    WukongDocId = WukongDocId + 1
    dc.update_one({'_id': book['_id']}, {'$set': {'wukongdocid': WukongDocId}})
