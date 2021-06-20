package repository

import (
	"errors"
	"ocg.com/product/model/entity"
)

type CategoryRepo struct {
	categories map[int64]*entity.Category
	autoID     int64
}

var Categories CategoryRepo

func init() {
	Categories = CategoryRepo{autoID: 0}
	Categories.categories = make(map[int64]*entity.Category)
}

func (r *CategoryRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *CategoryRepo) CreateCategory(category *entity.Category) *entity.Category {
	nextID := r.getAutoID()
	category.Id = nextID
	r.categories[nextID] = category
	return category
}

func (r *CategoryRepo) UpdateCategoryById(id int64, category *entity.Category) (*entity.Category, error) {
	if _, ok := r.categories[id]; ok {
		category.Id = id
		r.categories[id] = category
		return category, nil
	} else {
		return nil, errors.New("category not found")
	}
}

func (r *CategoryRepo) DeleteCategoryById(id int64) error {
	if _, ok := r.categories[id]; ok {
		delete(r.categories, id)
		return nil
	} else {
		return errors.New("category not found")
	}
}
