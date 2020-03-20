package lib

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
)

// Empty
type Empty struct{}

// GetPkgName get package name
func GetPkgName() string {
	return reflect.TypeOf(Empty{}).PkgPath()
}

// Exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// IsDir
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()

}

// Mkdir
func Mkdir(fpath string) error {
	err := os.MkdirAll(fpath, os.ModePerm)
	return err
}

// NewFile
func NewFile(filename string) (*os.File, error) {
	dir := Dir(filename)
	var err error
	if !Exists(dir) {
		err = Mkdir(dir)
		if err != nil {
			return nil, err
		}
	}
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// CreateFile
func CreateFile(filename string, src ...io.Reader) error {
	dir := Dir(filename)
	var err error
	if !Exists(dir) {
		err = Mkdir(dir)
		if err != nil {
			return err
		}
	}
	if len(src) > 0 {
		out, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			return err
		}
		defer out.Close()
		_, err = io.Copy(out, src[0])
		return err
	}
	// create file
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func Dir(filename string) string {
	return filepath.Dir(filename)
}

func Ext(path string) string {
	return filepath.Ext(path)
}

func GetFileContent(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func Copy(src string, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	dir := Dir(dst)
	if !Exists(dir) {
		err := Mkdir(dir)
		if err != nil {
			return err
		}
	}
	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}
	err = dstFile.Sync()
	if err != nil {
		return err
	}
	srcFile.Close()
	dstFile.Close()
	return nil
}
