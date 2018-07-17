package controllers

import (
	"doc4team/models"
	"fmt"
	"html/template"
	"strconv"

	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// DocController localhost:6060/doc
type DocController struct {
	// 如果在这里声明一个services接口，那么这个接口的实例化就需要到注册controller的时候讲需要的service实例注册，并这里变量名称必须大写（public）
	// 像这样：SvtoMarkDown services.IServiceToMarkDown

	// 在此项目中，并没有使用这样的做法，而是在services/base.go定义了接口的list并添加单例初始化方法GetInstance()
	// 在web/controllers/base.go中声明并初始化的了services接口

	Ctx iris.Context
}

// GetCreate get:localhost:6060/doc/create
func (c *DocController) GetCreate() mvc.Result {

	damap := make(map[string]interface{})
	damap["pagetitle"] = "建立新文档"
	damap["pageedit"] = false
	damap["pageform"] = "/doc/create"

	unused, err := servicesL.DocTag.GetDocTagList(1, 10)
	if err != nil {
		golog.Warn(err)
		return badResponse("/", err, 400)
	}
	damap["unused"] = unused

	return mvc.View{
		Name: "doc/edit.html",
		Data: iris.Map{"data": damap},
	}
}

// GetList get:localhost:6060/doc/list
func (c *DocController) GetList() mvc.Result {

	res, err := servicesL.ApiDoc.GetApiDocList()
	if err != nil {
		golog.Warn(err)
		return badResponse("/", err, 400)
	}

	return mvc.View{
		Name: "doc/list.html",
		Data: iris.Map{"data": res},
	}
}

// GetEditBy get:localhost:6060/doc/edit/{id}
func (c *DocController) GetEditBy(id int64) mvc.Result {

	damap := make(map[string]interface{})
	damap["pagetitle"] = "修改文档"
	damap["pageedit"] = true
	damap["pageform"] = "/doc/edit"
	// servicesL是定义在web/controllers/base.go中
	res, err := servicesL.ApiDoc.GetApiDoc(id)
	if err != nil {
		golog.Warn(err)
		return badResponse("/", err, 400)
	}

	used, unused, err := servicesL.DocTagMap.GetDocUsedAndUnusedTags(1, 10, id)
	if err != nil {
		golog.Warn(err)
		return badResponse("/", err, 400)
	}

	damap["doc"] = res
	damap["used"] = used
	damap["unused"] = unused
	return mvc.View{
		Name: "doc/edit.html",
		Data: iris.Map{"data": damap},
	}
}

// PostCreate post:localhost:6060/doc/create
func (c *DocController) PostCreate() mvc.Result {
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

	tags := c.Ctx.FormValue("usedtags")
	// if err := c.Ctx.ReadForm(&mapi); err != nil {
	// 	tools.LogErr(err)
	// 	return banReq("/", err, 400)
	// }

	res, err := servicesL.ApiDoc.Create(mapi, tags)

	if err != nil {
		golog.Warn(err)
		return badResponse("/", err, 400)
	}

	return mvc.Response{
		Path: "/doc/" + strconv.FormatInt(res, 10),
	}
}

// PostEdit get:localhost:6060/doc/edit
func (c *DocController) PostEdit() mvc.Result {
	mapi := models.ApiDoc{}

	adid, err := strconv.ParseInt(c.Ctx.FormValue("ApiDocid"), 10, 64)
	if err != nil {
		golog.Warn(err)
		return badResponse("/", err, 400)
	}

	mapi.ApiDocid = adid
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

	tags := c.Ctx.FormValue("usedtags")
	// if err := c.Ctx.ReadForm(&mapi); err != nil {
	// 	tools.LogErr(err)
	// 	return banReq("/", err, 400)
	// }

	res, err := servicesL.ApiDoc.Edit(mapi, tags)

	if err != nil {
		golog.Warn(err)
		return badResponse("/", err, 400)
	}

	return mvc.Response{
		Path: "/doc/" + strconv.FormatInt(res, 10),
	}
}

// GetBy get:localhost:6060/doc/{id}
func (c *DocController) GetBy(id int64) mvc.Result {

	res, err := servicesL.ApiDoc.GetApiDoc(id)
	if err != nil {
		golog.Warn(err)
		return badResponse("/", err, 400)
	}
	mkc, err := servicesL.ToMarkDown.ApiDoc2MarkDown(res)
	if err != nil {
		golog.Warn(err)
		return badResponse("/", err, 400)
	}

	usedtags, err := servicesL.DocTagMap.GetDocTagMapByDocId(id)
	if err != nil {
		golog.Warn(err)
		// return badResponse("/", err, 400)
	}
	fmt.Println(mkc)

	damap := make(map[string]interface{})
	damap["mkc"] = template.HTML(mkc)
	damap["used"] = usedtags
	return mvc.View{
		Name: "doc/detail.html",
		Data: iris.Map{"data": damap},
	}
}
