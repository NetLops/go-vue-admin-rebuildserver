package utils

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

//
//  ClearTable
//  Description: 清楚数据库表数据
//  param db 数据库对象
//  param tableName 表名
//  param compareField 比较字段
//  param interval 间隔
//  return error
//
func ClearTable(db *gorm.DB, tableName string, compareField string, interval string) error {
	if db == nil {
		return errors.New("db cannot be empty")
	}
	duration, err := time.ParseDuration(interval)
	if err != nil {
		return err
	}
	if duration < 0 {
		return errors.New("parse duration < 0")
	}
	return db.Debug().Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", tableName, compareField), time.Now().Add(-duration)).Error
}
