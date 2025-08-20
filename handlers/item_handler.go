package handlers

import (
	"inventory-system/config"
	"inventory-system/models"

	"github.com/gofiber/fiber/v2"
)

// GET /items
func GetItems(c *fiber.Ctx) error {
	var items []models.Item
	config.DB.Find(&items)
	return c.JSON(items)
}

// POST /items
func CreateItem(c *fiber.Ctx) error {
	item := new(models.Item)
	if err := c.BodyParser(item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}
	config.DB.Create(&item)
	return c.Status(fiber.StatusCreated).JSON(item)
}

// GET /items/:id
func GetItemByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Item
	result := config.DB.First(&item, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Item not found",
		})
	}
	return c.JSON(item)
}

// PUT /items/:id
func UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Item not found",
		})
	}

	updateData := new(models.Item)
	if err := c.BodyParser(updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	config.DB.Model(&item).Updates(updateData)
	return c.JSON(item)
}

// DELETE /items/:id
func DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Item
	if err := config.DB.First(&item, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Item not found",
		})
	}

	config.DB.Delete(&item)
	return c.JSON(fiber.Map{
		"message": "Item deleted",
	})
}
