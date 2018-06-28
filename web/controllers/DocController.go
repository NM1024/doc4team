package controllers

import (
	"fmt"
	"html/template"
	"doc4team/models"
	"doc4team/tools"
	"strconv"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// MovieController is our /movies controller.
type DocController struct {
	// Our MovieService, it's an interface which
	// is binded from the main application.
	// Service services.ApiDetails
	// SvApiDoc     services.IServiceApiDoc
	// SvtoMarkDown services.IServiceToMarkDown
	Ctx iris.Context
}

func (c *DocController) GetCreate() mvc.Result {

	damap := make(map[string]interface{})
	damap["pagetitle"] = "建立新文档"
	damap["pageedit"] = false
	damap["pageform"] = "/doc/create"

	return mvc.View{
		Name: "doc/edit.html",
		Data: iris.Map{"data": damap},
	}
}

func (c *DocController) GetEditBy(id int64) mvc.Result {

	damap := make(map[string]interface{})
	damap["pagetitle"] = "修改文档"
	damap["pageedit"] = true
	damap["pageform"] = "/doc/edit"
	res, err := servicesL.ApiDoc.GetApiDoc(id)
	if err != nil {
		tools.LogErr(err)
		return banReq("/", err, 400)
	}

	damap["doc"] = res
	return mvc.View{
		Name: "doc/edit.html",
		Data: iris.Map{"data": damap},
	}
}

func (c *DocController) PostCreate() mvc.Result {
	// username := c.Ctx.FormValue("username")
	// password := c.Ctx.FormValue("password")

	mapi := models.ApiDoc{}

	mapi.Name = c.Ctx.FormValue("Name")
	mapi.Version = c.Ctx.FormValue("Version")
	mapi.Address = c.Ctx.FormValue("Address")
	mapi.Method = c.Ctx.FormValue("Method")
	mapi.Parameters = c.Ctx.FormValue("Parameters")
	mapi.Header = c.Ctx.FormValue("Header")
	mapi.Body = c.Ctx.FormValue("Body")
	mapi.Response = c.Ctx.FormValue("Response")
	mapi.Describe = c.Ctx.FormValue("Describe")
	mapi.Remark = c.Ctx.FormValue("Remark")

	// if err := c.Ctx.ReadForm(&mapi); err != nil {
	// 	tools.LogErr(err)
	// 	return banReq("/", err, 400)
	// }

	res, err := servicesL.ApiDoc.Create(mapi)

	if err != nil {
		tools.LogErr(err)
		return banReq("/", err, 400)
	}

	return mvc.Response{
		Path: "/doc/" + strconv.FormatInt(res, 10),
	}
}

func (c *DocController) GetBy(id int64) mvc.Result {

	res, err := servicesL.ApiDoc.GetApiDoc(id)
	if err != nil {
		tools.LogErr(err)
		return banReq("/", err, 400)
	}
	// res, err := services.GetApiDocInstance().GetApiDoc(id)
	// if err != nil {
	// 	tools.LogErr(err)
	// }
	// ctx.StatusCode(iris.StatusOK)
	// ctx.ViewData("ApiDoc", res)
	// ctx.View("doc/detail.html")

	// ctx.Markdown(markdownContents)
	mkc, err := servicesL.ToMarkDown.ApiDoc2MarkDown(res)
	if err != nil {
		tools.LogErr(err)
		return banReq("/", err, 400)
	}
	fmt.Println(mkc)
	// mkdata := blackfriday.Run([]byte(mkc), blackfriday.WithNoExtensions())
	return mvc.View{
		Name: "doc/detail.html",
		// Data: iris.Map{"mkdata": template.HTML(string(mkdata))},
		Data: iris.Map{"mkdata": template.HTML(mkc)},
	}
}

func banReq(path string, err error, code int) mvc.Response {
	return mvc.Response{
		// if not nil then this error will be shown instead.
		Err: err,
		// redirect to /user/me.
		Path: "/",
		// When redirecting from POST to GET request you -should- use this HTTP status code,
		// however there're some (complicated) alternatives if you
		// search online or even the HTTP RFC.
		// Status "See Other" RFC 7231, however iris can automatically fix that
		// but it's good to know you can set a custom code;
		Code: code,
	}
}
