#!/usr/local/env python
# coding=utf-8
#
# Author: Archer
# File: LoadDefaultWechatusers.py
# Desc: 载入默认的美丽阅读用户的wechatusers表格
# Date: 5/Aug/2016
#
# Produced By BR
from pymongo import MongoClient

# connect to mongodb
client = MongoClient('mongodb://127.0.0.1:27017/bookshelf')
pclient = MongoClient('mongodb://rio:VFZPhT7y@192.168.200.22:27017/bookshelf')
# client = MongoClient('mongodb://192.168.100.2:27017/bookshelf')
# client = MongoClient('mongodb://rio:VFZPhT7y@192.168.200.22:27017/bookshelf')
db = client['bookshelf']
pdb = pclient['bookshelf']
wc = db['wechatusers']
uc = pdb['userext']

UserIds = []
with open('../data/defaultusers.csv') as F:
    for line in F:
        UserIds.append(line.strip('\n'))

# first, load default user from userext
Users = uc.find({'user_id': {'$in': UserIds}})
# for user in Users:
#     print user['user_name'], user['sex'], user['avatar']

# second, insert to wechatusers
for user in Users:
    data = {}
    data['user_name'] = user['user_name']
    data['sex'] = user['sex']
    data['avatar'] = user['avatar']
    data['default'] = True

    id = wc.insert_one(data).inserted_id
    print id
