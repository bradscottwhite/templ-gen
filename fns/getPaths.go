package fns

import (
  "os"
  "fmt"
  "path"
)

func GetSrcPath() string {
  exe, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}

  return path.Dir(exe)
}

func GetDstPath() string {
  dir, err := os.Getwd()
  if err != nil {
    fmt.Println(err)
  }

  return dir
}
