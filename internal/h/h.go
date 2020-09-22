package h

import (
	"fmt"
	"io/ioutil"
	"os"
)

func WriteFile(path string, b []byte) error {
	if !IsExist(path) {
		return fmt.Errorf("path not exist")
	}
	if err := ioutil.WriteFile(path, b, 0644); err != nil {
		return err
	}
	return nil
}

func IsExist(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
