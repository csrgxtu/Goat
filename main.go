package main

import (
	_ "Goat/routers"
	"github.com/astaxie/beego"
  "github.com/astaxie/beego/plugins/cors"
)

func main() {
  beego.InsertFilter("*", beego.BeforeRouter,cors.Allow(&cors.Options{
  AllowOrigins: []string{"*"},
  AllowMethods: []string{"*"},
  AllowHeaders: []string{"Origin"},
  ExposeHeaders: []string{"Content-Length"},
  AllowCredentials: true,
  }))

	beego.SetStaticPath("/", "static")

	beego.Run()
}
