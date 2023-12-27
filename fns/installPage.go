package fns

import (
	"fmt"
  
  "github.com/iancoleman/strcase"
)

func InstallPage(name string) {
  src := fmt.Sprintf("%s/files/newPage.templ", GetSrcPath())
  dst := fmt.Sprintf("%s/src/%sPage.templ", GetDstPath(), strcase.ToLowerCamel(name))
  if err := CreateFile(src, dst, name); err != nil {
    fmt.Println(err)
  }
  
  // Set Router
  dst = fmt.Sprintf("%s/main.go", GetDstPath())
  if err := CreateNewRoute(dst, name); err != nil {
    fmt.Println(err)
  }
}
