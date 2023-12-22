package fns

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
)

func InstallFile(pkg string) {
  src := fmt.Sprintf("%s/files", GetSrcPath())
  dst := GetDstPath()
  if err := copyDir(src, dst); err != nil {
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
