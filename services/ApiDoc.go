package services

import (
	"doc4team/models"
)

type IServiceApiDoc interface {
	Create(models.ApiDoc) (int64, error)
	GetApiDoc(int64) (models.ApiDoc, error)
	GetApiDocList() (map[int64]models.ApiDoc, error)
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

// Create
func (r *apiDoc) Create(ma models.ApiDoc) (int64, error) {

	ma, err := daoL.ApiDoc.Insert(ma)
	if err != nil {
		return 0, err
	}
	return ma.ApiDocid, nil
}

// GetApiDoc
func (r *apiDoc) GetApiDoc(adid int64) (models.ApiDoc, error) {
	return daoL.ApiDoc.GetApiDoc(adid)
}

// GetApiDocList
func (r *apiDoc) GetApiDocList() (map[int64]models.ApiDoc, error) {
	return daoL.ApiDoc.GetApiDocList()
}
