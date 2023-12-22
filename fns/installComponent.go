package fns

import (
	"fmt"
	"io/ioutil"
	"os"
  "strings"

  "github.com/iancoleman/strcase"
)

func InstallComponent(name string) {
  // Make name lowercamelcase:
  name = strcase.ToLowerCamel(name)

  src := fmt.Sprintf("%s/componentFiles/newComponent.templ", GetSrcPath())
  dst := fmt.Sprintf("%s/src/%s.templ", GetDstPath(), name)
  if err := createFile(src, dst, name); err != nil {
    fmt.Println(err)
  }
}

// File copies a single file from src to dst and replace '%NAME%' variable with name
func createFile(src, dst, name string) error {
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
