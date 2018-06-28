package controllers

import (
	"doc4team/services"
)

var servicesL *services.ServiceList

func init() {
	servicesL = services.GetInstance()
}
