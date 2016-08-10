# -*- coding: utf-8 -*-
# flake8: noqa

from qiniu import Auth, put_file, etag, urlsafe_base64_encode
import qiniu.config

#需要填写你的 Access Key 和 Secret Key
access_key = '_4TUdWfMQGZ5f2DFFmXbARs7pQLWmiPK-IFbSsw5'
secret_key = '1x0lUvV11qxbWQO1G_XrMm6v-MSsDWJWNCJk2K67'

#构建鉴权对象
q = Auth(access_key, secret_key)

#要上传的空间
bucket_name = 'brpublic'

Files = ["rere161.png"]
# with open('names.csv') as F:
#     for line in F:
#         Files.append(line.strip('\n'))

for f in Files:
    print f
    token = q.upload_token(bucket_name, f, 3600)
    localfile = '../static/img/' + f
    ret, info = put_file(token, f, localfile)
    print info
#
# #上传到七牛后保存的文件名
# key = 'my-python-logo.png';
#
# #生成上传 Token，可以指定过期时间等
# token = q.upload_token(bucket_name, key, 3600)
#
# #要上传文件的本地路径
# localfile = './sync/bbb.jpg'
#
# ret, info = put_file(token, key, localfile)
# print(info)
# assert ret['key'] == key
# assert ret['hash'] == etag(localfile)
