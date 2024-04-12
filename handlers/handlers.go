package handlers

import (
	"github.com/dizorin/go-bytebin/database"
	"github.com/dizorin/go-bytebin/utils"
	"github.com/gofiber/fiber/v2"
)

// PasteCreate returns paste page
func PasteCreate(c *fiber.Ctx) error {
	text := string(c.Body())
	key := utils.GenerateID()
	database.DB[key] = text

	return c.JSON(fiber.Map{
		"key": key,
	})
}

func PasteGet(c *fiber.Ctx) error {
	id := c.Params("id")
	text, found := database.DB[id]

	if !found {
		return c.Status(404).SendString("Not found")
	}

	return c.SendString(text)
}

// NotFound returns custom 404 page
//func NotFound(c *fiber.Ctx) error {
//	return c.Status(404).SendFile("./public/404.html")
//}
