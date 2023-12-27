package fns

import (
  "os"
  "log"
	"io/ioutil"
  "fmt"
)

func GetFiles(path string) []os.FileInfo {
  files, err := ioutil.ReadDir(fmt.Sprintf("%s/files/base", path))
  if err != nil {
    log.Fatal(err)
  }
  return files
}

func GetSize(path string) (int64, error) {
  file, err := os.Stat(path)
  if err != nil {
    return 0, err
  }
  return file.Size(), nil
}
