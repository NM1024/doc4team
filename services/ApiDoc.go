package services

import (
	"doc4team/dao"
	"doc4team/models"
	"doc4team/tools"
)

type IServiceApiDoc interface {
	Create(models.ApiDoc) (int64, error)
	GetApiDoc(int64) (models.ApiDoc, error)
}

// var once sync.Once
// var m *apiDoc

// func GetApiDocInstance() IServiceApiDoc {

// 	once.Do(func() {
// 		m = &apiDoc{}
// 	})
// 	return m
// }

// func GetApiDocInstance() IServiceApiDoc {
// 	return &apiDoc{}
// }

type apiDoc struct {
}

func (r *apiDoc) Create(ma models.ApiDoc) (int64, error) {

	ma, err := dao.GetApiDocInstance().Insert(ma)
	if err != nil {
		tools.LogErr(err)
	}
	return ma.ApiDocid, nil
}
func (r *apiDoc) GetApiDoc(adid int64) (models.ApiDoc, error) {
	return dao.GetApiDocInstance().GetApiDoc(adid)
}
