package libutils

import (
	"errors"
	"os"
)

func CreateFile(file_path string) error {

	path_seperator := '/'
	seperate_index := -1

	for i := len(file_path) - 1; i >= 0; i -= 1 {
		if file_path[i] == byte(path_seperator) {
			seperate_index = i
			break
		}
	}

	if seperate_index > 0 {
		only_folder := file_path[:seperate_index]
		err := os.MkdirAll(only_folder, os.ModePerm)
		if err != nil {
			return err
		}
	}

	_, err := os.Create(file_path)
	if err != nil {
		return err
	}

	return nil
}

func IsExist(file_path string) bool {
	if _, err := os.Stat(file_path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
