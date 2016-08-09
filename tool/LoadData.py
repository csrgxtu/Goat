#!/usr/bin/env python
# coding=utf-8
#
# Author: Archer Reilly
# File: LoadData.py
# Desc: 将分类描述载入到数据库里面去
# Date: 03/Aug/2016
from pymongo import MongoClient

# connect to mongodb
# client = MongoClient('mongodb://127.0.0.1:27017/bookshelf')
# client = MongoClient('mongodb://192.168.100.2:27017/bookshelf')
client = MongoClient('mongodb://rio:VFZPhT7y@192.168.200.22:27017/bookshelf')
db = client['bookshelf']
cc = db['classification']

RAW_DATA = []
with open('test.csv') as F:
    for line in F:
        RAW_DATA.append(line.strip('\n\r'))

TagsA = []
with open('taga.csv') as F:
    for line in F:
        TagsA.append(line.strip('\n\r'))

TagsB = []
with open('tagb.csv') as F:
    for line in F:
        TagsB.append(line.strip('\n\r'))

# print RAW_DATA
for data in RAW_DATA:
    print data
    # print data.split(' ')[0], data.split(' ')[1]
    if len(data.split(',')[3]) == 1:
        name = '00' + data.split(',')[3]
    elif len(data.split(',')[3]) == 2:
        name = '0' + data.split(',')[3]
    else:
        name = data.split(',')[3]

    document = {
        "clc_sort_num": data.split(',')[0],
        "description": data.split(',')[1],
        "main": data.split(',')[2],
        "img": 'http://7xj2i2.com2.z0.glb.qiniucdn.com/' + name + ".png",
        "taga": TagsA,
        "tagb": TagsB
    }
    id = cc.insert_one(document).inserted_id
    print id
