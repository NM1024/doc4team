package models

import "time"

type ApiDoc struct {
	ApiDocid   int64  `xorm:"'api_id' not null pk autoincr BIGINT"`
	Name       string `xorm:"comment('名称') VARCHAR(100)"`
	Address    string `xorm:"comment('API地址') VARCHAR(100)"`
	Method     string `xorm:"comment('类型') VARCHAR(50)"`
	Header     string `xorm:"comment('请求头') VARCHAR(1000)"`
	Parameters string `xorm:"VARCHAR(1000)"`
	Body       string `xorm:"VARCHAR(1000)"`
	Response   string `xorm:"VARCHAR(1000)"`
	Describe   string `xorm:"comment('描述') VARCHAR(300)"`
	Version    string `xorm:"comment('版本') VARCHAR(20)"`
	Remark     string `xorm:"comment('备注') VARCHAR(300)"`

	CreatedAt time.Time `xorm:"created DATETIME" json:"createdAt"`
	UpdatedAt time.Time `xorm:"updated DATETIME" json:"updatedAt"`
	DeletedAt time.Time `xorm:"deleted DATETIME" json:"deletedAt"`
}

// 设置对应的数据库表名称
func (a *ApiDoc) TableName() string {
	return "apidoc"
}
