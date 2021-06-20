package repository

import (
	"errors"
	"ocg.com/product/model/entity"
)

type ReviewRepo struct {
	reviews map[int64]*entity.Review
	autoID  int64
}

var Reviews ReviewRepo

func init() {
	Reviews = ReviewRepo{autoID: 0}
	Reviews.reviews = make(map[int64]*entity.Review)
}

func (r *ReviewRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *ReviewRepo) CreateReview(review *entity.Review) *entity.Review {
	nextID := r.getAutoID()
	review.Id = nextID
	r.reviews[nextID] = review
	return review
}

func (r *ReviewRepo) FindAllReview() (result []*entity.Review) {
	mapReviews := r.reviews
	for _, review := range mapReviews {
		result = append(result, review)
	}
	return result
}

func (r *ReviewRepo) FindReviewById(Id int64) (*entity.Review, error) {
	if review, ok := r.reviews[Id]; ok {
		return review, nil
	} else {
		return nil, errors.New("review not fount")
	}
}

func (r *ReviewRepo) UpdateReviewById(review *entity.Review) error {
	if _, ok := r.reviews[review.Id]; ok {
		r.reviews[review.Id] = review
		return nil
	} else {
		return errors.New("review not found")
	}
}

func (r *ReviewRepo) GetReviewByProductId(productId int64) (result map[int64][]*entity.Review) {
	result = make(map[int64][]*entity.Review)
	reviews := r.FindAllReview()
	for _, review := range reviews {
		if productId == review.ProductId {
			result[productId] = append(result[productId], review)
		}
	}
	return result
}

func (r *ReviewRepo) DeleteReviewById(Id int64) error {
	if _, ok := r.reviews[Id]; ok {
		delete(r.reviews, Id)
		return nil
	} else {
		return errors.New("review not found")
	}
}
