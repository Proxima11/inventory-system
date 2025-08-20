package handlers

import (
	"inventory-system/config"
	"inventory-system/models"

	"github.com/gofiber/fiber/v2"
)

// POST /stock/in
func StockIn(c *fiber.Ctx) error {
	var input models.Transaction
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if input.Quantity <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Quantity must be greater than zero"})
	}

	var item models.Item
	if err := config.DB.First(&item, input.ItemID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
	}

	item.Quantity += input.Quantity
	config.DB.Save(&item)

	tx := models.Transaction{
		ItemID:   item.ID,
		Type:     models.StockIn,
		Quantity: input.Quantity,
		Note:     input.Note,
	}
	config.DB.Create(&tx)

	return c.JSON(fiber.Map{
		"message": "Stock added",
		"item":    item,
	})
}

// POST /stock/out
func StockOut(c *fiber.Ctx) error {
	var input models.Transaction
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	if input.Quantity <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Quantity must be greater than zero"})
	}

	var item models.Item
	if err := config.DB.First(&item, input.ItemID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Item not found"})
	}

	if item.Quantity < input.Quantity {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Not enough stock"})
	}

	item.Quantity -= input.Quantity
	config.DB.Save(&item)

	tx := models.Transaction{
		ItemID:   item.ID,
		Type:     models.StockOut,
		Quantity: input.Quantity,
		Note:     input.Note,
	}
	config.DB.Create(&tx)

	return c.JSON(fiber.Map{
		"message": "Stock removed",
		"item":    item,
	})
}
