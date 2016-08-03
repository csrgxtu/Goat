package routers

import (
	"Goat/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/goat/bookdetail/:query", &controllers.BookDetailController{}, "get:SearchBookDetail")
}
