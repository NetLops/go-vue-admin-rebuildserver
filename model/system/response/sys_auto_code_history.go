package response

import "time"

type AutoCodeHistory struct {
	ID           uint      `json:"ID" gorm:"column:id"`
	CreateAt     time.Time `json:"CreateAt" gorm:"column:create_at"`
	UpdateAt     time.Time `json:"UpdateAt" gorm:"column:update_at"`
	TableName    string    `json:"tableName" gorm:"column:table_name"`
	StructName   string    `json:"structName" gorm:"column:struct_name"`
	StructCNName string    `json:"structCNName" gorm:"column:struct_cn_name"`
	Flag         int       `json:"flag" gorm:"column:flag"`
}
