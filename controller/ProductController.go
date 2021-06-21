package controller

import (
	"github.com/gofiber/fiber/v2"
	"ocg.com/product/model/entity"
	"ocg.com/product/service"
)

func CreateProduct(c *fiber.Ctx) error {
	product := new(entity.Product)
	err := c.BodyParser(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	return c.JSON(service.CreateProduct(product))
}

func GetAllProduct(c *fiber.Ctx) error {
	return c.JSON(service.GetAllProduct())
}

func GetProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	return c.JSON(service.GetProductById(int64(id)))
}

func DeleteProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	service.DeleteProductById(int64(id))
	return c.JSON(fiber.Map{
		"message":    "deleted",
		"product_id": id,
	})
}

func UpdateProductById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	product := new(entity.Product)
	err = c.BodyParser(&product)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	return c.JSON(service.UpdateProductById(int64(id), product))
}

func GetPriceBeforeUpdateById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	return c.JSON(service.GetPriceBeforeUpdateById(int64(id)))
}
