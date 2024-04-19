package main

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dizorin/go-bytebin/database"
	"github.com/dizorin/go-bytebin/handlers"

	"github.com/gofiber/fiber/v2"
)

func TestCreateNewPaste(t *testing.T) {
	app := fiber.New()

	api := app.Group("/api/paste")
	api.Get("/:id", handlers.PasteGet)
	api.Post("/", handlers.PasteCreate)

	// Create a POST request to create a new paste
	req := httptest.NewRequest("POST", "/api/paste", strings.NewReader("text=Hello, World!"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Add your assertions for checking the response body and other details
}

func TestRetrieveExistingPaste(t *testing.T) {
	app := fiber.New()

	api := app.Group("/api/paste")
	api.Get("/:id", handlers.PasteGet)
	api.Post("/", handlers.PasteCreate)

	// Create a test paste in the database
	testID := "test123"
	database.db[testID] = "This is a test paste"

	// Create a GET request to retrieve the test paste
	req := httptest.NewRequest("GET", "/api/paste/"+testID, nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Add your assertions for checking the response body and other details
}

func TestRetrieveNonExistentPaste(t *testing.T) {
	app := fiber.New()

	api := app.Group("/api/paste")
	api.Get("/:id", handlers.PasteGet)
	api.Post("/", handlers.PasteCreate)

	// Create a GET request for a non-existent paste
	req := httptest.NewRequest("GET", "/api/paste/nonexistent", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 404 {
		t.Errorf("Expected status 404, got %d", resp.StatusCode)
	}
}
