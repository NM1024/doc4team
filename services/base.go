package services

import (
	"doc4team/dao"
	"sync"
)

var once sync.Once
var m *ServiceList

// GetInstance 设置ServiceList的单例
func GetInstance() *ServiceList {
	once.Do(func() {
		m = &ServiceList{}
		m.ApiDoc = &apiDoc{}
		m.ToMarkDown = &toMarkDown{}
		m.DocTag = &docTag{}
		m.DocTagMap = &docTagMap{}
	})
	return m
}

// ServiceList 定义所有的services接口
type ServiceList struct {
	ApiDoc     IServiceApiDoc
	ToMarkDown IServiceToMarkDown
	DocTag     IServiceDocTag
	DocTagMap  IServiceDocTagMap
}

var daoL *dao.DaoList

func init() {
	daoL = dao.GetInstance()
}
