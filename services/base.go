package services

import "sync"

var once sync.Once
var m *ServiceList

func GetInstance() *ServiceList {
	once.Do(func() {
		m = &ServiceList{}
		m.ApiDoc = &apiDoc{}
		m.ToMarkDown = &toMarkDown{}
	})
	return m
}

type ServiceList struct {
	ApiDoc     IServiceApiDoc
	ToMarkDown IServiceToMarkDown
}
