package controller

import (
	"github.com/gofiber/fiber/v2"
	"ocg.com/product/model/entity"
	"ocg.com/product/service"
)

func CreateCategory(c *fiber.Ctx) error {
	category := new(entity.Category)
	err := c.BodyParser(&category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	return c.JSON(service.CreateCategory(category))
}

func UpdateCategory(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	category := new(entity.Category)
	err = c.BodyParser(&category)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	return c.JSON(service.UpdateCategoryById(int64(id), category))
}

func DeleteCategoryById(c *fiber.Ctx) error  {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	service.DeleteCategoryById(int64(id))
	return c.JSON(fiber.Map{
		"message" : "deleted",
		"category_id" : id,
	})
}
