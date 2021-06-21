package routes

import (
	"github.com/gofiber/fiber/v2"
	"ocg.com/product/controller"
)

func ConfigUserRouter(router *fiber.Router) {
	(*router).Get("/users", controller.GetAllUser)
	(*router).Post("/users", controller.CreateUser)
	(*router).Get("/users/:id", controller.GetUserById)
	(*router).Delete("/users/:id", controller.DeleteUserById)
	(*router).Put("/users/:id", controller.UpdateUserById)
}

func ConfigProductRouter(router *fiber.Router) {
	(*router).Get("/products", controller.GetAllProduct)
	(*router).Post("/products", controller.CreateProduct)
	(*router).Get("/products/:id", controller.GetProductById)
	(*router).Put("/products/:id", controller.UpdateProductById)
	(*router).Delete("/products/:id", controller.DeleteProductById)
}

func ConfigReviewRouter(router *fiber.Router) {
	(*router).Get("/reviews", controller.GetAllReview)
	(*router).Post("/reviews", controller.CreateNewReView)
	(*router).Get("/products/:id/reviews", controller.GetReviewByProductId)
	(*router).Put("reviews/:id", controller.UpdateReviewById)
	(*router).Delete("reviews/:id", controller.DeleteReviewById)
}

func ConfigCategoryRouter(router *fiber.Router) {
	(*router).Post("/categories", controller.CreateCategory)
	(*router).Delete("/categories/:id", controller.DeleteCategoryById)
	(*router).Put("/categories/:id", controller.UpdateCategory)
}
