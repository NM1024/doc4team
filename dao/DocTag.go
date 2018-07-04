package dao

import (
	"doc4team/models"
	. "doc4team/tools"
	"errors"
)

type IDaoDocTag interface {
	Insert(models.DocTag) (models.DocTag, error)
	GetDocTag(int64) (models.DocTag, error)
	GetDocTagList(int, int) (map[int64]models.DocTag, error)
	GetDocUnusedTags(start int, count int, docid int64, tagids []int64) (map[int64]models.DocTag, error)
}

type docTag struct {
}

func (r *docTag) Insert(tagmod models.DocTag) (models.DocTag, error) {

	affected, err := Xdb.Insert(&tagmod)
	if err != nil {
		return models.DocTag{}, err
	}
	if affected <= 0 {
		return models.DocTag{}, errors.New("插入失败")
	}
	return tagmod, nil
}

func (r *docTag) GetDocTag(id int64) (models.DocTag, error) {

	tagmod := new(models.DocTag)
	has, err := Xdb.Id(id).Get(tagmod)

	if err != nil {
		return models.DocTag{}, err
	}
	if !has {
		return models.DocTag{}, errors.New("无数据")
	}
	return *tagmod, nil
}

func (r *docTag) GetDocTagList(start int, count int) (map[int64]models.DocTag, error) {
	tags := make(map[int64]models.DocTag)

	err := Xdb.Limit(count, start).Find(tags)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *docTag) GetDocUnusedTags(start int, count int, docid int64, tagids []int64) (map[int64]models.DocTag, error) {
	tags := make(map[int64]models.DocTag)
	err := Xdb.NotIn("tag_id", tagids).Limit(count, start).Find(tags)

	if err != nil {
		return nil, err
	}
	return tags, nil
}
