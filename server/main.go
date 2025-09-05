package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/shimodii/sina-uni-project/config"
	"github.com/shimodii/sina-uni-project/controller"
)

func main() {
  fmt.Println("server is starting ...")
  app := fiber.New()

  config.UploadsPath()

  app.Get("/", controller.Root)
  app.Post("/upload", controller.Upload)

  fmt.Println("server is running on port 3030")
  app.Listen(":3030")
}
