package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

func GetSize(f multipart.File) (int, error) {
	bytes, err := ioutil.ReadAll(f)
	return len(bytes), err
}

func GetExt(fileName string) string {
	return path.Ext(fileName)
}

func CheckExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

func CheckPermission(src string) bool {

	_, err := os.Stat(src)
	return os.IsPermission(err)
}

func IsNotExistMDir(src string) error {
	if noExist := CheckExist(src); noExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil

}

func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	file, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}

	return file, nil
}
