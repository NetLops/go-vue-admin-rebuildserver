package system

import (
	"database/sql"
	"fmt"
	adapter "github.com/casbin/gorm-adapter/v3"
	"rebuildServer/global"
	"rebuildServer/model/example"
	"rebuildServer/model/system"
	"rebuildServer/model/system/request"
)

type InitDBService struct {
}

// InitDB
//
// Description: 创建数据库并初始化 总入口
//
// receiver: initDBService
//
// param: conf request.InitDB
//
// return: error
func (initDBService *InitDBService) InitDB(conf request.InitDB) error {
	switch conf.DBType {
	case "mysql":
		return initDBService.initMysqlDB(conf)
	case "pgsql":
		return initDBService.initPgsqlDB(conf)
	default:
		return initDBService.initMysqlDB(conf)
	}
}

// InitTables
//
// Description: 初始化表
//
// receiver: initDBService
//
//
// return: error
func (initDBService *InitDBService) InitTables() error {
	return global.GVA_DB.AutoMigrate(
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.SysAuthority{},
		system.JwtBlacklist{},
		system.SysDictionary{},
		system.SysAutoCodeHistory{},
		system.SysOperationRecord{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},

		adapter.CasbinRule{},

		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},
	)
}

// createDatabase
//
// Description: 创建数据库(mysql)
//
// receiver: initDBService
//
// param: dsn string
// param: driver string
// param: createSql string
//
// return: error
func (initDBService *InitDBService) createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err2 := db.Close()
		if err2 != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}
