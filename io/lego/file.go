package lego

import (
	"io/ioutil"
	"os"
)

func ReadFile(path string) (string, error) {
  dat, err := ioutil.ReadFile(path)
  return string(dat), err
}

func WriteFile(path string, value string) error {
  data := []byte(value)
  return ioutil.WriteFile(path, data, 0644)
}

func ListFiles(path string) ([]string, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

  stringFiles := make([]string, len(files))
  for i, file := range files {
    stringFiles[i] = file.Name()
  }
  return stringFiles, nil
}

func FileExists(path string) bool {
  _, err := os.Stat(path)
  return os.IsNotExist(err)
}