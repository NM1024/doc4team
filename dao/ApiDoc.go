package dao

import (
	"errors"
	"doc4team/models"
	. "doc4team/tools"
	"sync"
)

type IDaoApiDoc interface {
	Insert(models.ApiDoc) (models.ApiDoc, error)
	GetApiDoc(int64) (models.ApiDoc, error)
}

var m IDaoApiDoc
var once sync.Once

func GetApiDocInstance() IDaoApiDoc {
	once.Do(func() {
		m = &apiDoc{}
	})
	return m
}

type apiDoc struct {
}

func (r *apiDoc) Insert(docmod models.ApiDoc) (models.ApiDoc, error) {

	affected, err := Xdb.Insert(&docmod)
	if err != nil {
		return models.ApiDoc{}, err
	}
	if affected <= 0 {
		return models.ApiDoc{}, errors.New("插入失败")
	}
	return docmod, nil
}

func (r *apiDoc) GetApiDoc(id int64) (models.ApiDoc, error) {

	docmod := new(models.ApiDoc)
	has, err := Xdb.Id(id).Get(docmod)

	if err != nil {
		return models.ApiDoc{}, err
	}
	if !has {
		return models.ApiDoc{}, errors.New("无数据")
	}
	return *docmod, nil
}
