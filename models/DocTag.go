package models

import (
	"time"
)

type DocTag struct {
	TagId     int64     `xorm:"not null pk autoincr BIGINT(20)"`
	Name      string    `xorm:"VARCHAR(50)"`
	CreatedAt time.Time `xorm:"DATETIME"`
	UpdatedAt time.Time `xorm:"DATETIME"`
	DeletedAt time.Time `xorm:"DATETIME"`
}

// 设置对应的数据库表名称
func (a *DocTag) TableName() string {
	return "doctag"
}
