package services

import (
	"doc4team/models"
)

type IServiceDocTag interface {
	Create(models.DocTag) (int64, error)
	GetDocTag(int64) (models.DocTag, error)
	GetDocTagList(pageindex int, pagesize int) (map[int64]models.DocTag, error)
}

type docTag struct {
}

// Create
func (r *docTag) Create(tag models.DocTag) (int64, error) {

	tag, err := daoL.DocTag.Insert(tag)
	if err != nil {
		return 0, err
	}
	return tag.TagId, nil
}

// GetDocTag
func (r *docTag) GetDocTag(adid int64) (models.DocTag, error) {
	return daoL.DocTag.GetDocTag(adid)
}

// GetDocTagList
func (r *docTag) GetDocTagList(pageindex int, pagesize int) (map[int64]models.DocTag, error) {
	start := pagesize * (pageindex - 1)
	count := pagesize
	return daoL.DocTag.GetDocTagList(start, count)
}
