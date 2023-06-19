package fs

import (
	"os"
)

func IsFileExist(filename string) bool {
	_, err := os.Stat(filename)

	return err == nil
}

func GetFile(filename string) (*os.File, error) {
	data, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func WriteFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func DeleteFile(filename string) error {
	err := os.Remove(filename)
	if err != nil {
		return err
	}

	return nil
}
