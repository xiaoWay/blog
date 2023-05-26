package utils

import (
	"errors"
	"os"
)

// 判断文件目录是否存在
func PathExists(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	// 存在文件夹
	if fileInfo.IsDir() {
		return true, nil
	}
	// 不是文件夹, 则存在同名文件
	return false, errors.New("存在同名文件")
}

//func CreateDir(dirs ...string) (err error) {
//	for _, v := range dirs {
//		exist, err := PathExists(v)
//		if err != nil {
//			return err
//		}
//		if !exist {
//			global.GVA_LOG.Debug("create directory" + v)
//			if err := os.MkdirAll(v, os.ModePerm); err != nil {
//				global.GVA_LOG.Error("create directory"+v, zap.Any(" error:", err))
//				return err
//			}
//		}
//	}
//	return err
//}
