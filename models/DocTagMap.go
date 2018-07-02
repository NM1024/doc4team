package models

import (
	"time"
)

type DocTagMap struct {
	Id        int64     `xorm:"not null pk autoincr BIGINT(20)"`
	ApiId     int64     `xorm:"not null BIGINT(20)"`
	TagId     int64     `xorm:"not null BIGINT(20)"`
	CreatedAt time.Time `xorm:"DATETIME"`
	UpdatedAt time.Time `xorm:"DATETIME"`
	DeletedAt time.Time `xorm:"DATETIME"`
}

// 设置对应的数据库表名称
func (a *DocTagMap) TableName() string {
	return "doctagmap"
}
