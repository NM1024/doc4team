package dao

import (
	"doc4team/models"
	. "doc4team/tools"
	"errors"
)

type IDaoApiDoc interface {
	Insert(models.ApiDoc, []*models.DocTagMap) (models.ApiDoc, error)
	GetApiDoc(int64) (models.ApiDoc, error)
	GetApiDocList() (map[int64]models.ApiDoc, error)
	Update(models.ApiDoc, []*models.DocTagMap) (models.ApiDoc, error)
}

type apiDoc struct {
}

func (r *apiDoc) Insert(docmod models.ApiDoc, tagmap []*models.DocTagMap) (models.ApiDoc, error) {
	session := Xdb.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return models.ApiDoc{}, err
	}

	// 插入apidoc
	affected, err := session.Insert(&docmod)
	if err != nil {
		session.Rollback()
		return models.ApiDoc{}, err
	}
	if affected <= 0 {
		session.Rollback()
		return models.ApiDoc{}, errors.New("插入失败35")
	}

	// 验证tagmap len
	if len(tagmap) > 0 && tagmap != nil {
		for _, t := range tagmap {
			t.ApiId = docmod.ApiDocid
		}
		// 插入
		affected, err = session.Insert(&tagmap)
		if err != nil {
			session.Rollback()
			return models.ApiDoc{}, err
		}
		if affected <= 0 {
			session.Rollback()
			return models.ApiDoc{}, errors.New("插入失败45")
		}
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

func (r *apiDoc) Update(docmod models.ApiDoc, tagmap []*models.DocTagMap) (models.ApiDoc, error) {
	session := Xdb.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		return models.ApiDoc{}, err
	}
	// 更新apidoc
	affected, err := session.Id(docmod.ApiDocid).AllCols().Update(&docmod)
	if err != nil {
		session.Rollback()
		return models.ApiDoc{}, err
	}
	if affected <= 0 {
		session.Rollback()
		return models.ApiDoc{}, errors.New("更新失败96")
	}

	// 查询tagmap条数
	var dtm models.DocTagMap
	dtm.ApiId = docmod.ApiDocid
	counts, err := session.Count(&dtm)
	if err != nil {
		session.Rollback()
		return models.ApiDoc{}, err
	}
	if counts > 0 {
		// 删除tagmap
		affected, err = session.Delete(&dtm)
		if err != nil {
			session.Rollback()
			return models.ApiDoc{}, err
		}
		if affected <= 0 {
			session.Rollback()
			return models.ApiDoc{}, errors.New("更新失败108")
		}
	}

	//重新插入tagmap
	affected, err = session.Insert(&tagmap)
	if err != nil {
		session.Rollback()
		return models.ApiDoc{}, err
	}
	if affected <= 0 {
		session.Rollback()
		return models.ApiDoc{}, errors.New("更新失败118")
	}

	err = session.Commit()
	if err != nil {
		return models.ApiDoc{}, err
	}

	return docmod, nil
}
