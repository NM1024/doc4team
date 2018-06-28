package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// HomeController localhost:6060/ 注意这里不是localhost:6060/home
// 因为在注册此comtroller的时候并没有加home
type HomeController struct {
	Ctx iris.Context
}

// Get  get:localhost:6060
func (c *HomeController) Get() mvc.Result {
	return mvc.View{
		Name: "home/index.html",
	}
}
