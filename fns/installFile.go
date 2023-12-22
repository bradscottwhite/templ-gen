package fns

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
  "strings"
)

func InstallFile() {
  srcPath := GetSrcPath()
  dstPath := GetDstPath()

  src := fmt.Sprintf("%s/files", srcPath)
  dst := dstPath
  if err := copyDir(src, dst); err != nil {
    fmt.Println(err)
  }
  
  // Change go.mod file to include NAME variable
  dir := strings.Split(dstPath, "/")
  dirName := dir[len(dir)-1]
  src = fmt.Sprintf("%s/files/go.mod", srcPath)
  dst = fmt.Sprintf("%s/go.mod", dstPath)
  if err := CreateFile(src, dst, dirName); err != nil {
    fmt.Println(err)
  }
}

// File copies a single file from src to dst
func copyFile(src, dst string) error {
	var err error
	var srcfd *os.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = os.Open(src); err != nil {
    fmt.Println(err)
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
    fmt.Println(err)
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
    fmt.Println(err)
		return err
	}
	if srcinfo, err = os.Stat(src); err != nil {
    fmt.Println(err)
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}

// copyDir copies a whole directory recursively
func copyDir(src string, dst string) error {
	var err error
	var fds []os.FileInfo
	var srcinfo os.FileInfo

	if srcinfo, err = os.Stat(src); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = ioutil.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = copyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = copyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}
