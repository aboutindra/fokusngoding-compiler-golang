package main

import (
	"fc-golang/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var r router.Router

func init() {
	r = router.Router{}
}

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Post("/golang/v1/compiler/run", r.Exec)
	app.Post("/golang/v1/compiler/unit-test", r.Exec)

	app.Listen("0.0.0.0:4000")
}
