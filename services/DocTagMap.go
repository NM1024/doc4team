package services

import (
	"doc4team/models"
)

type IServiceDocTagMap interface {
	Create(models.DocTagMap) (int64, error)
	GetDocTagMap(int64) (models.DocTagMap, error)
	GetDocTagMapByDocId(int64) (map[int64]models.DocTag, error)
	GetDocUsedAndUnusedTags(pageindex int, pagesize int, docid int64) (map[int64]models.DocTag, map[int64]models.DocTag, error)
}

type docTagMap struct {
}

// Create
func (r *docTagMap) Create(tag models.DocTagMap) (int64, error) {

	tag, err := daoL.DocTagMap.Insert(tag)
	if err != nil {
		return 0, err
	}
	return tag.TagId, nil
}

// GetDocTagMap
func (r *docTagMap) GetDocTagMap(adid int64) (models.DocTagMap, error) {
	return daoL.DocTagMap.GetDocTagMap(adid)
}

// GetDocTagMapByDocId
func (r *docTagMap) GetDocTagMapByDocId(docid int64) (map[int64]models.DocTag, error) {
	return daoL.DocTagMap.GetDocTagMapByDocId(docid)
}

// GetDocUsedAndUnusedTags
func (r *docTagMap) GetDocUsedAndUnusedTags(pageindex int, pagesize int, docid int64) (map[int64]models.DocTag, map[int64]models.DocTag, error) {
	start := pagesize * (pageindex - 1)
	count := pagesize

	used, err := r.GetDocTagMapByDocId(docid)
	if err != nil {
		return nil, nil, err
	}
	tagids := make([]int64, 0)
	for _, v := range used {
		tagids = append(tagids, v.TagId)
	}
	unused, err := daoL.DocTag.GetDocUnusedTags(start, count, docid, tagids)
	if err != nil {
		return nil, nil, err
	}
	return used, unused, nil
}
