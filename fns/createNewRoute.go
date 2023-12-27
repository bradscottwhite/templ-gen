package fns

import (
	"io/ioutil"
	"os"
  "strings"
  "fmt"

  "github.com/iancoleman/strcase"
)

// File copies the Main.go file contents from dst and appends a new route
func CreateNewRoute(dst, name string) error {
  // Make name lowercamelcase:
  name = strcase.ToLowerCamel(name)
	
  var err error
	var srcinfo os.FileInfo


  txt, err := ioutil.ReadFile(dst)
  if err != nil {
    fmt.Println(err)
		return err
  }

  errLogTxt := "log.Fatal(app.Listen(\":3000\"))"
  newTxt := fmt.Sprintf("component = components.%sPage()\n\tapp.Get(\"/%s\", adaptor.HTTPHandler(templ.Handler(component)))\n\n\t%s", name, name, errLogTxt)

  err = ioutil.WriteFile(dst, []byte(strings.Replace(string(txt), errLogTxt, newTxt, -1)), 0644)
  if err != nil {
    fmt.Println(err)
		return err
  }

	if srcinfo, err = os.Stat(dst); err != nil {
    fmt.Println(err)
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}
