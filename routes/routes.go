package routes

import (
	"inventory-system/handlers"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes mendaftarkan semua endpoint utama
func SetupRoutes(app *fiber.App) {
	// Group: /items
	items := app.Group("/items")
	items.Get("/", handlers.GetItems)
	items.Post("/", handlers.CreateItem)
	items.Get("/:id", handlers.GetItemByID)
	items.Put("/:id", handlers.UpdateItem)
	items.Delete("/:id", handlers.DeleteItem)

	// Group: /stock
	stock := app.Group("/stock")
	stock.Post("/in", handlers.StockIn)
	stock.Post("/out", handlers.StockOut)
}
