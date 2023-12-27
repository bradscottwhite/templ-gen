package fns

import (
	"io/ioutil"
	"os"
  "strings"
  "fmt"
)

// File copies a single file from src to dst and replace '%NAME%' variable with name
func CreateFile(src, dst, name string) error {	
  var err error
	var srcinfo os.FileInfo

  txt, err := ioutil.ReadFile(src)
  if err != nil {
    fmt.Println(err)
		return err
  }

  err = ioutil.WriteFile(dst, []byte(strings.Replace(string(txt), "%NAME%", name, -1)), 0644)
  if err != nil {
    fmt.Println(err)
		return err
  }

	if srcinfo, err = os.Stat(src); err != nil {
    fmt.Println(err)
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}
