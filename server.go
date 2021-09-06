package main

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	"github.com/gofiber/template/html"
)

func main() {
	httpClient := &http.Client{Timeout: 10 * time.Second}

	app := fiber.New(fiber.Config{
		Views: html.New("./templates", ".html"),
	})
	app.Use(recover.New())
	app.Use(helmet.New())
	app.Use(limiter.New())

	v1 := app.Group("v1")
	app.Get("/:ore", HandleGetPrice(httpClient))
	app.Get("/price/:ore", HandleGetPrice(httpClient))
	v1.Get("/price/:ore", HandleGetPrice(httpClient))

	app.Listen("localhost:8282")
}
