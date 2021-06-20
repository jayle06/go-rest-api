package controller

import (
	"github.com/gofiber/fiber/v2"
	"ocg.com/product/model/entity"
	"ocg.com/product/service"
)

func GetAllReview(c *fiber.Ctx) error {
	return c.JSON(service.GetAllReview())
}

func CreateNewReView(c *fiber.Ctx) error {
	review := new(entity.Review)
	err := c.BodyParser(&review)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	newReview := service.CreateReview(review)
	return c.JSON(newReview)
}

func UpdateReviewById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	review := new(entity.Review)
	err := c.BodyParser(&review)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cannot parse JSON",
			"error":   err,
		})
	}
	updateReview := service.UpdateReView(int64(id), review)
	return c.JSON(updateReview)

}

func GetReviewByProductId(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	return c.JSON(service.GetReviewByProductId(int64(id)))
}

func DeleteReviewById(c *fiber.Ctx) error {
	id, _ := c.ParamsInt("id")
	err, _ := service.DeleteReviewById(int64(id))
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message":        "Review was deleted",
		"product_rating": "rating was updated",
	})
}
