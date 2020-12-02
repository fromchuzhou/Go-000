package main

import (
	"myapp/controller"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New() // 实例一个iris对象
	app.Logger().SetLevel("debug")
	//配置路由
	// app.Get("/", func(ctx iris.Context) {
	// 	ctx.WriteString("Hello Iris")
	// })

	mvc.New(app.Party("hello")).Handle(new(controller.MovieController))
	// 启动服务器
	app.RegisterView(iris.HTML("D:/myapp/views", ".html"))
	app.Run(iris.Addr(":7999"))
}
