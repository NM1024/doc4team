package services

import (
	"encoding/json"
	"errors"
	"doc4team/models"
	"doc4team/tools"
)

type IServiceToMarkDown interface {
	ApiDoc2MarkDown(models.ApiDoc) (string, error)
}

// var m *toMarkDown

// func GetToMarkDownInstance() IToMarkDown {

// 	m.once.Do(func() {
// 		m = &toMarkDown{}
// 	})
// 	return m
// }

// func GetToMarkDownInstance() IServiceToMarkDown {
// 	return &toMarkDown{}
// }

type toMarkDown struct {
}

func (r *toMarkDown) ApiDoc2MarkDown(ml models.ApiDoc) (string, error) {
	var str string

	jshead := make(map[string]interface{})
	err := json.Unmarshal([]byte(ml.Header), &jshead)
	if err != nil {
		return "", errors.New("failed header to json")
	}

	var md = new(tools.MarkDown)
	md.Title(ml.Name+ml.Version, 1).NewLines(2)
	md.Title("API-URL:", 2).NewLines(2).Text(ml.Method + ":").CodeOneLine(ml.Address).NewLines(2)
	md.Title("Request Headers:", 2).NewLines(2).Code(ml.Header, "json").NewLine().Table(jshead).NewLines(2)
	md.Title("Request Body:", 2).NewLines(2).Code(ml.Body, "json").NewLines(2)
	md.Title("说明:", 2).NewLines(2).Text(ml.Describe).NewLines(2)
	md.Title("备注:", 2).NewLines(2).Text(ml.Remark).NewLines(2)
	str = md.Content
	return str, nil
}
