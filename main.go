package main

import (
	"log"

	"github.com/ads6495/applimate-api/database"
	"github.com/ads6495/applimate-api/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)

	// user endpoints
	app.Post("/api/user", routes.CreateUser)
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
