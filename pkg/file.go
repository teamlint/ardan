package pkg

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Empty
type Empty struct{}

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

// putContents 写入文件内容
func putContents(path string, data []byte, flag int, perm os.FileMode) error {
	// 支持目录递归创建
	dir := Dir(path)
	if !Exists(dir) {
		if err := Mkdir(dir); err != nil {
			return err
		}
	}
	// 创建/打开文件
	f, err := os.OpenFile(path, flag, perm)
	if err != nil {
		return err
	}
	defer f.Close()
	n, err := f.Write(data)
	if err != nil {
		return err
	} else if n < len(data) {
		return io.ErrShortWrite
	}
	return nil
}

// Truncate changes the size of the named file.
func Truncate(path string, size int) error {
	return os.Truncate(path, int64(size))
}

// WriteFile (文本)写入文件内容
func WriteFile(path string, content []byte) error {
	return putContents(path, content, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
}

// DeleteFile delete file
func DeleteFile(filename string, delDir bool) error {
	dir := Dir(filename)
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	if delDir {
		err = os.Remove(dir)
		if err != nil {
			return err
		}
	}
	return nil
}

// Move 文件移动/重命名
func Move(src string, dst string) error {
	dir := Dir(dst)
	if !Exists(dir) {
		err := Mkdir(dir)
		if err != nil {
			return err
		}
	}
	return os.Rename(src, dst)
}

// Rename 文件移动/重命名
func Rename(src string, dst string) error {
	return Move(src, dst)
}

// Filename 获取指定文件路径的文件名称
func Filename(path string) string {
	return filepath.Base(path)
}

// TempDir 系统临时目录
func TempDir() string {
	return os.TempDir()
}
