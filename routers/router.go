package routers

import (
	"Goat/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/goat/bookdetail/:query", &controllers.BookDetailController{}, "get:SearchBookDetail")
	beego.Router("/goat/classification/:clc_sort_num", &controllers.BookDetailController{}, "get:GetClcInfo")
}
