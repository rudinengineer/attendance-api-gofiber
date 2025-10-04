package main

import (
	"absensi-api/internal/config"
	"absensi-api/internal/connection"
	"absensi-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Setup Configuration
	configuration := config.Load()

	// DB Connection
	dbconnection := connection.ConnectDB(configuration.Database)

	// Setup Routes
	routes.New(app, dbconnection)

	// Listen Server
	log.Fatal(app.Listen(configuration.Server.Host + ":" + configuration.Server.Port))
}
