package services

import (
	"doc4team/models"
	"strconv"
	"strings"
)

type IServiceApiDoc interface {
	Create(models.ApiDoc, string) (int64, error)
	GetApiDoc(int64) (models.ApiDoc, error)
	GetApiDocList() (map[int64]models.ApiDoc, error)
	SearchApiDoc(string, string) (map[int64]models.ApiDoc, error)
	Edit(models.ApiDoc, string) (int64, error)
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
func (r *apiDoc) Create(ma models.ApiDoc, tags string) (int64, error) {

	tagsmap := make([]*models.DocTagMap, 0)
	artags := strings.Split(tags, ",")
	for _, tag := range artags {
		if tag != "" {
			tid, err := strconv.ParseInt(tag, 10, 64)
			if err != nil {
				return 0, err
			}
			tmap := new(models.DocTagMap)
			tmap.TagId = tid

			tagsmap = append(tagsmap, tmap)
		}
	}
	ma, err := daoL.ApiDoc.Insert(ma, tagsmap)
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

// Edit
func (r *apiDoc) Edit(ma models.ApiDoc, tags string) (int64, error) {

	tagsmap := make([]*models.DocTagMap, 0)
	artags := strings.Split(tags, ",")
	for _, tag := range artags {
		if tag != "" {
			tid, err := strconv.ParseInt(tag, 10, 64)
			if err != nil {
				return 0, err
			}
			tmap := new(models.DocTagMap)
			tmap.ApiId = ma.ApiDocid
			tmap.TagId = tid

			tagsmap = append(tagsmap, tmap)
		}
	}
	ma, err := daoL.ApiDoc.Update(ma, tagsmap)
	if err != nil {
		return 0, err
	}
	return ma.ApiDocid, nil
}

// SearchApiDoc
func (r *apiDoc) SearchApiDoc(keyword string, tags string) (map[int64]models.ApiDoc, error) {
	doctags := make([]int64, 0)
	artags := strings.Split(tags, ",")
	for _, tag := range artags {
		if tag != "" {
			tid, err := strconv.ParseInt(tag, 10, 64)
			if err != nil {
				return nil, err
			}
			doctags = append(doctags, tid)
		}
	}

	return daoL.ApiDoc.SearchApiDoc(keyword, doctags)
}
