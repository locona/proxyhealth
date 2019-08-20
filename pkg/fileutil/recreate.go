package fileutil

import (
	"os"
	"path/filepath"
)

func Recreate(fileName string) (*os.File, string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return nil, "", err
	}

	filePath := filepath.Join(pwd, fileName)
	_, err = os.Stat(filePath)
	if err == nil {
		_err := os.Remove(filePath)
		if _err != nil {
			return nil, "", _err
		}
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, "", err
	}

	return file, filePath, nil
}
