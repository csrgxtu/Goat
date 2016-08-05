package routers

import (
	"Goat/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/goat/bookdetail/:query", &controllers.BookDetailController{}, "get:SearchBookDetail")
	beego.Router("/goat/classification/:clc_sort_num", &controllers.BookDetailController{}, "get:GetClcInfo")
	beego.Router("/goat/indexer", &controllers.WukongController{}, "get:Indexer")
	beego.Router("/goat/searcher/:query", &controllers.WukongController{}, "get:Searcher")

	// 用户信息获取
	// beego.Router("/goat/users/:id", &controllers.UserController{}, "get:GetUserInfoById()")
}
