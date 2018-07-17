package controllers

import (
	"github.com/kataras/golog"
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

	damap := make(map[string]interface{})
	// damap["pageform"] = "/search"

	unused, err := servicesL.DocTag.GetDocTagList(1, 10)
	if err != nil {
		golog.Warn(err)
		return badResponse("/", err, 400)
	}
	damap["unused"] = unused

	return mvc.View{
		Name: "home/index.html",
		Data: iris.Map{"data": damap},
	}

}

// GetSearch get:localhost:6060/Search
func (c *HomeController) GetSearch() mvc.Result {
	keyword := c.Ctx.URLParam("kw")
	tags := c.Ctx.URLParam("s")

	docs, err := servicesL.ApiDoc.SearchApiDoc(keyword, tags)
	if err != nil {
		golog.Warn(err)
		return badResponse("/", err, 400)
	}

	damap := make(map[string]interface{})
	damap["docs"] = docs
	return mvc.View{
		Name: "doc/list.html",
		Data: iris.Map{"data": damap},
	}
}
