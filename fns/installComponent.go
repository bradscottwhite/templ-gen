package fns

import (
	"fmt"
)

func InstallComponent(name string) {
  src := fmt.Sprintf("%s/componentFiles/newComponent.templ", GetSrcPath())
  dst := fmt.Sprintf("%s/src/%s.templ", GetDstPath(), name)
  if err := CreateFile(src, dst, name); err != nil {
    fmt.Println(err)
  }
}
