package service

import (
	"ocg.com/product/model/entity"
	"ocg.com/product/repository"
)

func GetAllReview() []*entity.Review {
	return repository.Reviews.FindAllReview()
}

func CreateReview(review *entity.Review) *entity.Review {
	newReview := repository.Reviews.CreateReview(review)
	productId := newReview.ProductId
	product := GetProductById(productId)
	UpdateProductRating(product)
	return newReview
}

func DeleteReviewById(Id int64) (error, *entity.Product) {
	review, _ := repository.Reviews.FindReviewById(Id)
	productId := review.ProductId
	product := GetProductById(productId)
	return repository.Reviews.DeleteReviewById(Id), UpdateProductRating(product)
}

func UpdateReView(Id int64, newReview *entity.Review) *entity.Review {
	err := repository.Reviews.DeleteReviewById(Id)
	if err != nil {
		return nil
	}
	review := CreateReview(newReview)
	review.Id = Id
	product := GetProductById(review.ProductId)
	UpdateProductRating(product)
	return review
}

func GetReviewByProductId(productId int64) map[int64][]*entity.Review {
	return repository.Reviews.GetReviewByProductId(productId)
}
