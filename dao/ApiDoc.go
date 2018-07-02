package dao

import (
	"doc4team/models"
	. "doc4team/tools"
	"errors"
)

type IDaoApiDoc interface {
	Insert(models.ApiDoc) (models.ApiDoc, error)
	GetApiDoc(int64) (models.ApiDoc, error)
	GetApiDocList() (map[int64]models.ApiDoc, error)
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

func (r *apiDoc) GetApiDocList() (map[int64]models.ApiDoc, error) {
	docmods := make(map[int64]models.ApiDoc)

	err := Xdb.Find(docmods)
	if err != nil {
		return make(map[int64]models.ApiDoc), err
	}
	return docmods, nil
}
