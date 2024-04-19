package handlers

import (
	"github.com/dizorin/go-bytebin/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// PasteCreate returns paste page
func PasteCreate(c *fiber.Ctx) error {
	key, err := database.Save(c.Body())
	if err != nil {
		log.Error(err)
		return c.Status(500).SendString("Internal server error")
	}
	return c.JSON(fiber.Map{
		"key": key,
	})
}

func PasteGet(c *fiber.Ctx) error {
	id := c.Params("id")
	content, found := database.Load(id)

	if !found {
		return c.Status(404).SendString("Not found")
	}

	return c.Send(content.Body)
}

// NotFound returns custom 404 page
// func NotFound(c *fiber.Ctx) error {
//	return c.Status(404).SendFile("./public/404.html")
//}
