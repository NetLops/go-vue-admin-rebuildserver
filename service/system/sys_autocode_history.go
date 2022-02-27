package system

import (
	"fmt"
	"path/filepath"
	"rebuildServer/global"
	"rebuildServer/model/common/request"
	"rebuildServer/model/system"
	"rebuildServer/utils"
	"strings"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

var RepeatErr = errors.New("重复创建")

type AutoCodeHistoryService struct{}

var AutoCodeHistoryServiceApp = new(AutoCodeHistoryService)

// CreateAutoCodeHistory 创建代码生成器历史记录
// RouterPath: RouterPth@RouterString;RouterPath2@RouterString2
func (autoCodeHistoryService *AutoCodeHistoryService) CreateAutoCodeHistory(meta, structName, structCNName, autoCodePath string, injectionMeta string, tableName string, apiIds string) error {
	return global.GVA_DB.Create(&system.SysAutoCodeHistory{
		RequestMeta:   meta,
		AutoCodePath:  autoCodePath,
		InjectionMeta: injectionMeta,
		StructName:    structName,
		StructCNName:  structCNName,
		TableName:     tableName,
		ApiIDs:        apiIds,
	}).Error
}

// First 根据id 获取代码生成器历史的数据
func (autoCodeHistoryService *AutoCodeHistoryService) First(info *request.GetById) (string, error) {
	var meta string
	return meta, global.GVA_DB.Model(system.SysAutoCodeHistory{}).Select("request_meta").Where("id = ?", info.Uint()).First(&meta).Error
}

// Repeat 检测重复
func (autoCodeHistoryService *AutoCodeHistoryService) Repeat(structName string) bool {
	var count int64
	global.GVA_DB.Model(&system.SysAutoCodeHistory{}).Where("struct_name = ? and flag = 0", structName).Count(&count)
	return count > 0
}

// RollBack 回滚
func (autoCodeService *AutoCodeService) RollBack(info *request.GetById) error {
	md := system.SysAutoCodeHistory{}
	if err := global.GVA_DB.Where("id = ?", info.Uint()).First(&md).Error; err != nil {
		return err
	}
	// 清除API列表
	err := ApiServiceApp.DeleteApiByIds(strings.Split(md.ApiIDs, ";"))
	if err != nil {
		global.GVA_LOG.Error("ClearTag DeleteApiByIds:", zap.Error(err))
	}
	// 获取全部表名
	dbNames, err := AutoCodeServiceApp.Database().getTables(global.GVA_CONFIG.Mysql.Dbname)
	if err != nil {
		global.GVA_LOG.Error("ClearTag GetTables:", zap.Error(err))
	}
	// 删除表
	for _, name := range dbNames {
		if strings.Contains(strings.ToUpper(strings.Replace(name.tableName, "_", "", -1)), strings.ToUpper(md.TableName)) { // Replace方法中-1 代表不限制覆盖次数
			// 删除表
			if err = AutoCodeServiceApp.DropTable(name.TableName); err != nil {
				global.GVA_LOG.Error("ClearTag DropTable:", zap.Error(err))
			}
		}
	}

	// 删除文件

	for _, path := range strings.Split(md.AutoCodePath, ";") {

		// 增加安全判断补丁：
		_path, err := filepath.Abs(path)
		if err != nil || _path != path {
			continue
		}

		// 迁移
		nPath := filepath.Join(global.GVA_CONFIG.AutoCode.Root, "rm_file", time.Now().Format("20060102"), filepath.Base(filepath.Dir(filepath.Dir(path))), filepath.Base(filepath.Dir(path)), filepath.Base(path))
		// 判断目标文件是否存在
		for utils.FileExist(nPath) {
			fmt.Println("文件已存在;", nPath)
			nPath += fmt.Sprintf("_%d", time.Now().Nanosecond())
		}
		err = utils.FileMove(path, nPath)
		if err != nil {
			fmt.Println(">>>>>>>>>>>>>>>>", err)
		}
		// _ = utils.DelFile(path)
	}
	// 清除注入
	for _, v := range strings.Split(md.InjectionMeta, ";") {
		// RouterPath@functionName@RouterString
		meta := strings.Split(v, "@")
		if len(meta) == 3 {
			_ = utils.AutoClearCode(meta[0], meta[2])
		}
	}
	md.Flag = 1
	return global.GVA_DB.Save(&md).Error
}
