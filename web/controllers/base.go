package controllers

import (
	"doc4team/services"

	"github.com/kataras/iris/mvc"
)

var servicesL *services.ServiceList

func init() {
	servicesL = services.GetInstance()
}

func badResponse(path string, err error, code int) mvc.Response {
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
