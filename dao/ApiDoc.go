package dao

import (
	"doc4team/models"
	. "doc4team/tools"
	"errors"
)

type IDaoApiDoc interface {
	Insert(models.ApiDoc, []models.DocTagMap) (models.ApiDoc, error)
	GetApiDoc(int64) (models.ApiDoc, error)
	GetApiDocList() (map[int64]models.ApiDoc, error)
}

type apiDoc struct {
}

func (r *apiDoc) Insert(docmod models.ApiDoc, tagmap []models.DocTagMap) (models.ApiDoc, error) {
	session := Xdb.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return models.ApiDoc{}, err
	}

	affected, err := Xdb.Insert(&docmod)
	if err != nil {
		session.Rollback()
		return models.ApiDoc{}, err
	}
	if affected <= 0 {
		session.Rollback()
		return models.ApiDoc{}, errors.New("插入失败")
	}

	affected, err = Xdb.Insert(&tagmap)
	if err != nil {
		session.Rollback()
		return models.ApiDoc{}, err
	}
	if affected <= 0 {
		session.Rollback()
		return models.ApiDoc{}, errors.New("插入失败")
	}

	err = session.Commit()
	if err != nil {
		return models.ApiDoc{}, err
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
		return nil, err
	}
	return docmods, nil
}
