package routers

import (
	"Goat/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/goat/bookdetail/:query/:userid", &controllers.BookDetailController{}, "get:SearchBookDetail")
	beego.Router("/goat/classification/:clc_sort_num", &controllers.BookDetailController{}, "get:GetClcInfo")
	beego.Router("/goat/indexer", &controllers.WukongController{}, "get:Indexer")
	beego.Router("/goat/searcher/:query/:userid", &controllers.WukongController{}, "get:Searcher")

	// 用户信息获取
	beego.Router("/goat/users/:id", &controllers.UserController{}, "get:GetUserInfoById")
	beego.Router("/goat/users/similiar/:id", &controllers.UserController{}, "get:GetSimiliar")

  // wechat
  beego.Router("/goat/wechat/verify", &controllers.WechatController{}, "get:Verify")
  beego.Router("/goat/wechat/webauth", &controllers.WechatController{}, "get:WebAuth")
  beego.Router("/goat/wechat/signature", &controllers.WechatController{}, "get:Signature")
}
