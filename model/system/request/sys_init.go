package request

import (
	"fmt"
	"rebuildServer/config"
)

type InitDB struct {
	DBType   string `json:"DBType"`                      // 数据库类型
	Host     string `json:"host"`                        // 服务器地址
	Port     string `json:"port"`                        // 数据库连接端口
	UserName string `json:"userName" binding:"required"` // 数据库用户名
	Password string `json:"password"`                    // 数据库密码
	DBName   string `json:"DBName" binding:"required"`   // 数据库名
}

// MysqlEmptyDsn
//
// Description: mysql 空数据库 建表链接
//
// receiver: i
//
//
// return: string
func (i *InitDB) MysqlEmptyDsn() string {
	if i.Host == "" {
		i.Host = "127.0.0.1"
	}
	if i.Port == "" {
		i.Port = "3306"
	}
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/", i.UserName, i.Password, i.Host, i.Port)
}

// PgsqlEmptyDsn
//
// Description: pgsql 空数据库 建库连接
//
// receiver: i
//
//
func (i *InitDB) PgsqlEmptyDsn() string {
	if i.Host == "" {
		i.Host = "127.0.0.1"
	}
	if i.Port == "" {
		i.Port = "3306"
	}
	return "host=" + i.Host + " user=" + i.UserName + " password=" + i.Password + " port=" + i.Port + " " + "sslmode=disable TimeZone=Asia/Shanghai"
}

// ToMysqlConfig
//
// Description: 转换config.Mysql
//
// receiver: i
//
//
func (i *InitDB) ToMysqlConfig() config.MySql {
	return config.MySql{
		Path:     i.Host,
		Port:     i.Port,
		Dbname:   i.DBName,
		Username: i.UserName,
		Password: i.Password,
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}
}

// ToPgsqlConfig
//
// Description: 转换config.Pgsql
//
// receiver: i
//
//
// return: config.Pgsql
func (i *InitDB) ToPgsqlConfig() config.Pgsql {
	return config.Pgsql{
		Path:     i.Host,
		Port:     i.Port,
		Dbname:   i.DBName,
		Username: i.UserName,
		Password: i.Password,
		Config:   "sslmode=disable TimeZone=Asia/Shanghai",
	}
}
