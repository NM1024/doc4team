package models

import "time"

type baseModel struct {
	CreatedAt time.Time `xorm:"created DATETIME" json:"createdAt"`
	UpdatedAt time.Time `xorm:"updated DATETIME" json:"updatedAt"`
	DeletedAt time.Time `xorm:"deleted DATETIME" json:"deletedAt"`
}
