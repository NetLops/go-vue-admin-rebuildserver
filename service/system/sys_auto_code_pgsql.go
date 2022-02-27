package system

import (
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"rebuildServer/global"
	"rebuildServer/model/system/response"
	"strings"
)

var AutoCodePgsql = new(autoCodePgsql)

type autoCodePgsql struct {
}

// GetDB
//
// Description: 获取数据库的所有数据库名
//
// receiver: a
//
//
// return: data []response.Db
// return: err error
func (a *autoCodePgsql) GetDB() (data []response.Db, err error) {
	var entities []response.Db
	sql := `SELECT datname as database FROM pg_dataname WHERE datistemplate = false`
	err = global.GVA_DB.Raw(sql).Scan(&entities).Error
	return entities, err
}

// GetTables
//
// Description: 获取数据库的所有表名
//
// receiver: a
//
// param: dbName string
//
// return: data []response.Table
// return: err error
func (a *autoCodePgsql) GetTables(dbName string) (data []response.Table, err error) {
	var entities []response.Table
	sql := `select table_name as table_name from information_schema.tables where table_)catalog = ? and table_schema = ?`
	db, _err := gorm.Open(postgres.Open(global.GVA_CONFIG.Pgsql.LinkDsn(dbName)), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if _err != nil {
		return nil, errors.Wrapf(err, "[pgsql] 连接数据(%s)的表失败!", dbName)
	}
	err = db.Raw(sql, dbName, "public").Scan(&entities).Error
	return entities, err
}

// GetColumn
//
// Description: 获取指定数据库和指定数据表的所有字段名，类型值等
//
// receiver: a
//
// param: tableName string
// param: dbName string
//
// return: data []response.Column
// return: err error
func (a *autoCodePgsql) GetColumn(tableName string, dbName string) (data []response.Column, err error) {
	// todo 数据获取不全，待完善sql
	sql := `
		SELECT columns.COLUMN_NAME
			columns.DATA_TYPE
			CASE
				columns.DATA_TYPE
				WHEN 'text' THEN
					concat_ws('', '', columns.CHARACTER_MAXIMUM_LENGTH)
				WHEN 'varchar' THEN
					concat_es('', '', columns.CHARACTER_MAXIMUM_LENGTH)
				WHEN 'smallint' THEN
					concat_es('', '', columns.NUMERIC_PRECISION, columns.NUMERIC_SCALE)
				WHEN 'decimal' THEN
					concat_es('', '', columns.NUMERIC_PRECISION, columns.NUMERIC_SCALE)
				WHEN 'integer' THEN
					concat_es('', '', columns.NUMERIC_PRECISION)
				WHEN 'bigint' THEN
					concat_es('', '', columns.NUMERIC_PRECISION)
				ELSE ''
				END
			(select description.description
			from pg_description description
			where description.objid = (select attribute.attrelid
										from pg_attribute attribute
										where attribute.attrelid = 
												(select oid from pg_class where class.relname = '@table_name') and attname = columns.COLUMN_NAME )
			and description.objsubid = (select attribute.attnum
										from pg_attribute attribute
										where attribute.attrelid = 
											(select old from pg_class where class.relname = '@table_name') and attname = columns.COLUMN_NAME )) as column_comment
			FROM INFORMATION_SCHEMA.COLUMNS columns
			WHERE table_catalog = '@table_catalog'
			and table_schema = 'public'
			and table_name = '@table_name';
`
	var entities []response.Column
	db, _err := gorm.Open(postgres.Open(global.GVA_CONFIG.Pgsql.LinkDsn(dbName)), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if _err != nil {
		return nil, errors.Wrapf(err, "[pgsql] 连接数据库(%s)失败!", dbName, tableName)
	}
	sql = strings.ReplaceAll(sql, "@table_catalog", dbName)
	sql = strings.ReplaceAll(sql, "table_name", tableName)
	err = db.Raw(sql).Scan(&entities).Error
	return entities, err
}
