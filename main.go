package main

import (
	"flag"
	"github.com/dizorin/go-bytebin/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

func main() {
	// Parse command-line flags
	flag.Parse()

	// Connected with database
	//database.Connect()

	// Create fiber app
	app := fiber.New(fiber.Config{
		AppName: "Bytebin",
		Prefork: *prod, // go run app.go -prod
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())

	// Create a /api endpoint
	api := app.Group("/api/paste")

	// Bind handlers
	api.Get("/:id", handlers.PasteGet)
	api.Post("/", handlers.PasteCreate)

	// Setup static files
	app.Static("/", "./public")

	// Handle not founds
	//app.Use(handlers.NotFound)

	log.Fatal(app.Listen(*port))
}
