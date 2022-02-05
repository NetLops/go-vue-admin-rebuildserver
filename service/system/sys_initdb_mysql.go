package system

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"rebuildServer/config"
	"rebuildServer/global"
	model "rebuildServer/model/system"
	"rebuildServer/model/system/request"
	"rebuildServer/source/system"
	"rebuildServer/utils"
)

// writeMysqlConfig
//
// Description: mysql回写配置
//
// receiver: InitDBService
//
// param: sql config.MySql
//
// return: error
func (InitDBService *InitDBService) writeMysqlConfig(mysql config.MySql) error {
	global.GVA_CONFIG.Mysql = mysql
	cs := utils.StructToMap(global.GVA_CONFIG)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	global.GVA_VP.Set("jwt.signing-key", uuid.NewV4().String())
	return global.GVA_VP.WriteConfig()
}

func (initDBService *InitDBService) initMysqlDB(conf request.InitDB) error {
	dsn := conf.MysqlEmptyDsn()
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAUlT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DBName)
	if err := initDBService.createDatabase(dsn, "mysql", createSql); err != nil {
		return err
	}
	mysqlConfig := conf.ToMysqlConfig()
	if mysqlConfig.Dbname == "" { // 如果没有数据库名，则跳出初始化数据
		return nil
	}
	if db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       mysqlConfig.Dsn(), // DNS data source name
		DefaultStringSize:         191,               // string 类型字段的默认长度
		SkipInitializeWithVersion: true,              // 根据版本自动配置
	}), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		global.GVA_DB = db
	}

	if err := initDBService.InitTables(); err != nil {
		global.GVA_DB = nil
		return err
	}

}

func (initDBService *InitDBService) initMysqlData() error {
	return model.MysqlDataInitialize(
		system.Api,
		system.User,
	)
}
