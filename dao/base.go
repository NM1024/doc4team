package dao

import "sync"

var once sync.Once
var m *DaoList

// GetInstance 设置DaoList的单例
func GetInstance() *DaoList {
	once.Do(func() {
		m = &DaoList{}
		m.ApiDoc = &apiDoc{}
		m.DocTag = &docTag{}
		m.DocTagMap = &docTagMap{}
	})
	return m
}

// DaoList 定义所有的Dao接口
type DaoList struct {
	ApiDoc    IDaoApiDoc
	DocTag    IDaoDocTag
	DocTagMap IDaoDocTagMap
}
