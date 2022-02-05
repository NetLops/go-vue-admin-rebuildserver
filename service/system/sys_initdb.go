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

func (initDBService *InitDBService) InitDB(conf request.InitDB) error {
	switch conf.DBType {
	case "mysql":

	}
}

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
