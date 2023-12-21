package main

import (
  "log"
  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/a-h/templ"
  //components "templ-demo/components"
  components "templ-demo/src"
)

func main() {
  app := fiber.New()

  app.Static("/dist", "./dist")
	
  component := components.HomePage()

  app.Get("/", adaptor.HTTPHandler(templ.Handler(component)))

  log.Fatal(app.Listen(":3000"))
}
