package dao

import (
	"doc4team/models"
	. "doc4team/tools"
	"errors"
)

type IDaoDocTagMap interface {
	Insert(models.DocTagMap) (models.DocTagMap, error)
	GetDocTagMap(int64) (models.DocTagMap, error)
	GetDocTagMapByDocId(int64) (map[int64]models.DocTag, error)
}

type docTagMap struct {
}

func (r *docTagMap) Insert(tagmod models.DocTagMap) (models.DocTagMap, error) {

	affected, err := Xdb.Insert(&tagmod)
	if err != nil {
		return models.DocTagMap{}, err
	}
	if affected <= 0 {
		return models.DocTagMap{}, errors.New("插入失败")
	}
	return tagmod, nil
}

func (r *docTagMap) GetDocTagMap(id int64) (models.DocTagMap, error) {

	tagmod := new(models.DocTagMap)
	has, err := Xdb.Id(id).Get(tagmod)

	if err != nil {
		return models.DocTagMap{}, err
	}
	if !has {
		return models.DocTagMap{}, errors.New("无数据")
	}
	return *tagmod, nil
}

// GetDocTagMapByDocId
func (r *docTagMap) GetDocTagMapByDocId(docid int64) (map[int64]models.DocTag, error) {
	tags := make(map[int64]models.DocTag)
	err := Xdb.Table("doctagmap").Select("doctag.*").Join("INNER", "doctag", "doctagmap.tag_id = doctag.tag_id").Where("doctagmap.api_id = ?", docid).Find(tags)

	if err != nil {
		return nil, err
	}

	return tags, nil
}
