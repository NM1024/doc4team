package main

import (
	"doc4team/web/controllers"
	"os"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/view"

	"doc4team/tools"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	// 加载view模版
	views := iris.HTML("./views", ".html")
	// 加载模版函数
	templateFunc(views)
	app.RegisterView(views)
	// 静态资源
	app.StaticWeb("/", "./public")

	// 注册一个controller 这里的doc是下边定义的一个方法
	// app.Party("/doc")定义了起始地址 例如http://localhost:6060/doc/....
	mvc.Configure(app.Party("/doc"), doc)
	// 这也是注册controller 没有Party("...") 则指向http://localhost:6060/....
	// 这样使用的话切记不要出现相同地址
	mvc.New(app).Handle(new(controllers.HomeController))

	// 读取.env文件的配置
	port := os.Getenv("PORT")
	app.Run(
		// Start the web server at localhost:6060
		iris.Addr("localhost:"+port),
		// disables updates:
		iris.WithoutVersionChecker,
		// skip err server closed when CTRL/CMD+C pressed:
		iris.WithoutServerError(iris.ErrServerClosed),
		// enables faster json serialization and more:
		iris.WithOptimizations,
	)
}

// 在iris mvc的框架下 这里需要使用mvc.Application
func doc(app *mvc.Application) {
	// 这里可以为某个controller单独注册中间件或者其他包的内容
	// app.Router.Use(middleware.BasicAuth)
	// app.Register(services.GetApiDocInstance())
	// app.Register(services.GetToMarkDownInstance())

	app.Handle(new(controllers.DocController))
}

// 注册模版函数，自带的模版函数是不够用的，需要独立添加
// 以下的模版函数可以在tools/templateFunc.go中找到
// 可以将模版函数封装到一个struct中，在这里反射注入进行注册
func templateFunc(html *view.HTMLEngine) {
	html.AddFunc("HttpMethodList", tools.HttpMethodList)
	html.AddFunc("Obj2String", tools.Obj2String)
}
