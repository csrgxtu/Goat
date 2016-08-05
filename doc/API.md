### Goat 后台接口及微信分享信息

#### 部署测试服务器的微信使用
点击下面的链接授权跳转
[wechat auth](https://open.weixin.qq.com/connect/oauth2/authorize?appid=wx628bb9654a8c71c6&redirect_uri=https://dev-goat.beautifulreading.com/goat/wechat/webauth&response_type=code&scope=snsapi_userinfo&state=1&connect_redirect=1#wechat_redirect)

点击下面的使用微信jssdk
[wecaht jssdk](https://dev-goat.beautifulreading.com/static/index.html)


下面是我的微信测试账号信息
```bash
appid: wx628bb9654a8c71c6
secret: d4624c36b6795d1d99dcf0547af5443d

配置的域名：https://dev-goat.beautifulreading.com/goat/wechat/verify
Token: Goat

JS回调域名: dev-goat.beautifulreading.com

```
要测试上面的链接需要关注我的测试账号，扫描如下二维码.
![qr code](http://mmbiz.qpic.cn/mmbiz/jgDtMapChvkGDU4RRRzO9ACOCWVPzwv0Mqn4YJhoVYcDFNw11mLtFx5q960zRXV7NN5V47ia67XgDFkiciaKMmGBg/0)

注意，要测试下面的接口，请使用57a4ca7be7798977cd000001作为用户ｉｄ，或者你可以在100.2的wechatuser里面选择用户。

#### 根据书籍关键词搜索返回
```bash
GET /goat/bookdetail/:query/:userid

curl https://dev-goat.beautifulreading.com/goat/bookdetail/小王子/57a4ca7be7798977cd000001

返回
{
    "msg": "^_^",
    "data": [
        {
            "id": "5555b3af5694ab042b4069c6",
            "wukongdocid": 0,
            "isbn": "7530516388",
            "title": "小王子",
            "author": "(法)安东·德·圣艾修伯里著",
            "clc_sort_num": "I565.88",
            "status": "",
            "version": "0001-01-01T00:00:00Z",
            "createuserid": "admin",
            "updateuserid": "admin",
            "createtime": "",
            "updatetime": ""
        },
        "我就憋着不说话",
        [
            [
                "萌系",
                "自律",
                "阳光",
                "顽固",
                "热情",
                "健忘",
                "知性",
                "任性",
                "直率",
                "性感",
                "聪慧",
                "随和",
                "传统",
                "犀利",
                "精明"
            ],
            [
                "神经质",
                "执行力",
                "应声虫",
                "有恒心",
                "空想家"
            ],
            [
                "不着边际",
                "人文主义",
                "趾高气昂",
                "邋遢大王",
                "多重个性"
            ],
            [
                "抗拒陌生人",
                "有点小脾气",
                "理想主义者"
            ],
            [
                "大自然爱好者"
            ]
        ],
        {
            "id": "57a4ca7be7798977cd000001",
            "openid": "owingv5SwyNfawiY9usxaU1pgvhg",
            "user_id": "",
            "user_name": "archer",
            "avatar": "http://wx.qlogo.cn/mmopen/iawZD3n3LEGPsgwyUwV9diaYNuWMJdoicnqWDbNpCrrhj77TR62ZZGhOwoqTa5QohLhehgjdmwMoUibZ4MxGVzicIy0T5jJ2Guyoh/0",
            "bookdetailids": [],
            "similiraty": 0,
            "default": false,
            "sex": 1,
            "language": "en",
            "city": "广州",
            "province": "广东",
            "country": "中国"
        }
    ]
}
```

#### 获取用户基本信息
```bash
GET /goat/users/:id

curl https://dev-goat.beautifulreading.com/goat/users/57a4ca7be7798977cd000001

返回
{
    "msg": "^_^",
    "data": [
        {
            "id": "57a4ca7be7798977cd000001",
            "openid": "owingv5SwyNfawiY9usxaU1pgvhg",
            "user_id": "",
            "user_name": "archer",
            "avatar": "http://wx.qlogo.cn/mmopen/iawZD3n3LEGPsgwyUwV9diaYNuWMJdoicnqWDbNpCrrhj77TR62ZZGhOwoqTa5QohLhehgjdmwMoUibZ4MxGVzicIy0T5jJ2Guyoh/0",
            "bookdetailids": [
                "5555b3af5694ab042b4069c6"
            ],
            "similiraty": 0,
            "default": false,
            "sex": 1,
            "language": "en",
            "city": "广州",
            "province": "广东",
            "country": "中国"
        }
    ]
}
```

#### 获取相似用户
```bash
GET /goat/users/similiar/:id

curl https://dev-goat.beautifulreading.com/goat/users/similiar/57a4ca7be7798977cd000001

返回
{
    "msg": "^_^",
    "data": [
        {
            "id": "57a4ea8ef38544070dd3d6bd",
            "openid": "",
            "user_id": "",
            "user_name": "Anllela@当知",
            "avatar": "http://7xj2i2.com2.z0.glb.qiniucdn.com/2d36bc281494f54a7060a7bf48508c1c.png",
            "bookdetailids": null,
            "similiraty": 0.98,
            "default": true,
            "sex": 0,
            "language": "",
            "city": "",
            "province": "",
            "country": ""
        },
        {
            "id": "57a4ea8ef38544070dd3d6bf",
            "openid": "",
            "user_id": "",
            "user_name": "白珩",
            "avatar": "http://7xj2i2.com2.z0.glb.qiniucdn.com/3eaa545c5ca853623e176cea6f104321.jpg",
            "bookdetailids": null,
            "similiraty": 0.96,
            "default": true,
            "sex": 0,
            "language": "",
            "city": "",
            "province": "",
            "country": ""
        },
        {
            "id": "57a4ea8ef38544070dd3d6c0",
            "openid": "",
            "user_id": "",
            "user_name": "鲈鱼",
            "avatar": "http://7xj2i2.com2.z0.glb.qiniucdn.com/bbf3516c10ef189ce4a30c3b83631c11.jpg",
            "bookdetailids": null,
            "similiraty": 0.92,
            "default": true,
            "sex": 0,
            "language": "",
            "city": "",
            "province": "",
            "country": ""
        }
    ]
}
```
