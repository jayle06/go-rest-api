package service

import (
	"ocg.com/product/model/entity"
	"ocg.com/product/repository"
)

func CreateProduct(product *entity.Product) *entity.Product {
	return repository.Products.CreateProduct(product)
}

func GetAllProduct() map[int64]*entity.Product {
	return repository.Products.FindAllProduct()
}

func GetProductById(id int64) *entity.Product {
	product, _ := repository.Products.FindProductById(id)
	return product
}

func DeleteProductById(id int64) {
	_ = repository.Products.DeleteProductById(id)
}

func UpdateProductById(id int64, product *entity.Product) *entity.Product {
	updatedProduct, _ := repository.Products.UpdateProductById(id, product)
	return UpdateProductRating(updatedProduct)
}

func UpdateProductRating(product *entity.Product) *entity.Product {
	var sum int64 = 0
	var count int64 = 0
	reviews := repository.Reviews.FindAllReview()
	for _, review := range reviews {
		if product.Id == review.ProductId {
			sum += review.Rating
			count++
		}
	}
	rating := float64(sum) / float64(count)
	product.Rating = rating

	return product
}
