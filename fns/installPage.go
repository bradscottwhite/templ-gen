package fns

import (
	"fmt"
  
  "github.com/stoewer/go-strcase"
)

func InstallPage(name string) {
  src := fmt.Sprintf("%s/files/newPage.templ", GetSrcPath())
  dst := fmt.Sprintf("%s/src/%sPage.templ", GetDstPath(), strcase.LowerCamelCase(name))
  if err := CreateFile(src, dst, strcase.UpperCamelCase(name)); err != nil {
    fmt.Println(err)
  }
  
  // Set Router
  dst = fmt.Sprintf("%s/main.go", GetDstPath())
  if err := CreateNewRoute(dst, name); err != nil {
    fmt.Println(err)
  }
}
