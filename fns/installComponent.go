package fns

import (
	"fmt"
  
  "github.com/stoewer/go-strcase"
)

func InstallComponent(name string) {
  src := fmt.Sprintf("%s/files/newComponent.templ", GetSrcPath())
  dst := fmt.Sprintf("%s/src/%s.templ", GetDstPath(), strcase.LowerCamelCase(name))
  if err := CreateFile(src, dst, strcase.UpperCamelCase(name)); err != nil {
    fmt.Println(err)
  }
}
