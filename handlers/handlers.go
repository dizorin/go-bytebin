package handlers

import (
	"github.com/dizorin/go-bytebin/database"
	"github.com/dizorin/go-bytebin/utils"
	"github.com/gofiber/fiber/v2"
)

// PasteCreate returns paste page
func PasteCreate(c *fiber.Ctx) error {
	// Получить текст из body запроса
	text := c.FormValue("text")

	// Сгенерировать URL для пасты
	id := utils.GenerateID()
	url := "https://example.com/paste/" + id

	// Сохранить текст в базе данных
	database.DB[id] = text

	// Отправить JSON-ответ с URL
	return c.JSON(fiber.Map{
		"url": url,
	})
}

func PasteGet(c *fiber.Ctx) error {
	// Получить ID пасты из URL
	id := c.Params("id")

	// Получить текст пасты из базы данных
	text, found := database.DB[id]

	// Если паста не найдена
	if !found {
		return c.Status(404).SendString("Not found")
	}

	// Отправить текст пасты
	return c.SendString(text)
}

// NotFound returns custom 404 page
//func NotFound(c *fiber.Ctx) error {
//	return c.Status(404).SendFile("./public/404.html")
//}
