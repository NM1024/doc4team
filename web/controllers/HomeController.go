package controllers

import (
	"github.com/kataras/iris/mvc"
)

// MovieController is our /movies controller.
type HomeController struct {
	// Our MovieService, it's an interface which
	// is binded from the main application.
	// Service services.ApiDetails
}

func (c *HomeController) Get() mvc.Result {
	return mvc.View{
		Name: "home/index.html",
	}
	// return mvc.Response{Path: "home/index.html"}
	// ctx.View("home/index.html")
}
