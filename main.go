package main

import (
	"flag"
	"github.com/dizorin/go-bytebin/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"regexp"
)

var (
	port = flag.String("port", ":3000", "Port to listen on")
	prod = flag.Bool("prod", false, "Enable prefork in Production")
)

var rToken = regexp.MustCompile("^/[a-zA-Z\\d]{7}$")

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

	// Create a /documents endpoint
	api := app.Group("/documents")

	// Bind handlers
	api.Get("/:id", handlers.PasteGet)
	api.Post("/post", handlers.PasteCreate)

	app.Get("/*", func(c *fiber.Ctx) error {
		if rToken.MatchString(c.Path()) {
			return c.SendFile("./public/index.html")
		} else {
			return c.Next()
		}
	})

	// Setup static files
	app.Static("/", "./public")

	// Handle not founds
	//app.Use(handlers.NotFound)

	log.Fatal(app.Listen(*port))
}
