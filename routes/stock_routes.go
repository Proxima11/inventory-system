package routes

import (
	"inventory-system/handlers"

	"github.com/gofiber/fiber/v2"
)

func RegisterStockRoutes(app *fiber.App) {
	stock := app.Group("/stock")
	stock.Post("/in", handlers.StockIn)
	stock.Post("/out", handlers.StockOut)
}
