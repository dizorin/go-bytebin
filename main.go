package main

import (
	"context"
	"github.com/dizorin/go-bytebin/database/cache"
	"github.com/dizorin/go-bytebin/executor"
	"github.com/dizorin/go-bytebin/handlers"
	"github.com/dizorin/go-bytebin/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	// Connected with database
	// database.Connect()

	cache.Setup(ctx)
	executor.SetupScheduler(ctx)

	// Create fiber app
	app := fiber.New(fiber.Config{
		AppName: os.Getenv("APP_NAME"),
		Prefork: utils.GetenvBool("PREFORK"),

		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			log.Error(err)

			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"code":    fiber.StatusInternalServerError,
				"message": "Internal server error",
			})
		},
	})

	// Middleware
	app.Use(recover.New(recover.Config{
		StackTraceHandler: recover.ConfigDefault.StackTraceHandler,
		EnableStackTrace:  true,
	}))
	app.Use(logger.New())

	// Create a /documents endpoint
	api := app.Group("/documents")

	// Bind handlers
	api.Get("/:id", handlers.PasteGet)
	api.Post("/post", handlers.PasteCreate)

	app.Get("/*", func(c *fiber.Ctx) error {
		if utils.RegexToken.MatchString(c.Path()) {
			return c.SendFile("./public/index.html")
		} else {
			return c.Next()
		}
	})

	// Setup static files
	app.Static("/", "./public")

	// Handle not founds
	// app.Use(handlers.NotFound)

	log.Panic(app.Listen(os.Getenv("HOST")))
}
