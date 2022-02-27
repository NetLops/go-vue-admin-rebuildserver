package system

import (
	"rebuildServer/global"
	"rebuildServer/model/system/response"
)

type Database interface {
	GetDB() (data []response.Db, err error)
	GetTables(dbNaem string) (data []response.Table, err error)
	GetColumn(tableName string, dbName string) (data []response.Column, err error)
}

func (autoCodeService *AutoCodeService) Database() Database {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return AutoCodeMysql
	case "pgsql":
		return AutoCodePgsql
	default:
		return AutoCodeMysql
	}
}
