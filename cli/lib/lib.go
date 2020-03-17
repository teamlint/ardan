package lib

import (
	"os"
	"reflect"
)

type Empty struct{}

// GetPkgName get package name
func GetPkgName() string {
	return reflect.TypeOf(Empty{}).PkgPath()
}

// Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
