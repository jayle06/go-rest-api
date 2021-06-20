package controller

import (
	"github.com/gofiber/fiber/v2"
	"ocg.com/product/model/entity"
	"ocg.com/product/service"
)

func CreateUser(c *fiber.Ctx) error {
	user := new(entity.User)
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	return c.JSON(service.CreateUser(user))
}

func GetAllUser(c *fiber.Ctx) error {
	return c.JSON(service.GetAllUser())
}

func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	return c.JSON(service.GetUserById(int64(id)))
}

func DeleteUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	service.DeleteUserById(int64(id))
	return c.JSON(fiber.Map{
		"message": "deleted",
		"user_id": id,
	})
}

func UpdateUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	user := new(entity.User)
	err = c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	return c.JSON(service.UpdateUserById(int64(id), user))
}
