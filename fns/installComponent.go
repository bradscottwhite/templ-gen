package fns

import (
	"fmt"
  
  "github.com/iancoleman/strcase"
)

func InstallComponent(name string) {
  src := fmt.Sprintf("%s/files/newComponent.templ", GetSrcPath())
  dst := fmt.Sprintf("%s/src/%s.templ", GetDstPath(), strcase.ToLowerCamel(name))
  if err := CreateFile(src, dst, name); err != nil {
    fmt.Println(err)
  }
}
