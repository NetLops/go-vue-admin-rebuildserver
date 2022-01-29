package utils

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"os"
	"rebuildServer/global"
	"reflect"
)

//
//  PathExists
//  @Description: 文件目录是否存在
//  @param path
//  @return bool
//  @return error
//
func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}

	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//
//  CreateDir
//  @Description: 批量创建文件夹
//  @param dirs
//  @return err
//
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if err != nil {
			return nil
		}
		if !exist {
			fmt.Println("测试", v, reflect.TypeOf(v))
			global.GVA_LOG.Debug("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); err != nil {
				global.GVA_LOG.Error("create directory"+v, zap.Any("error", err))
				return err
			}
		}
	}
	return err
}
