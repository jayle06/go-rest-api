package service

import (
	"ocg.com/product/model/entity"
	"ocg.com/product/repository"
)

func CreateCategory(category *entity.Category) *entity.Category {
	return repository.Categories.CreateCategory(category)
}

func DeleteCategoryById(id int64) {
	_ = repository.Categories.DeleteCategoryById(id)
}

func UpdateCategoryById(id int64, category *entity.Category) *entity.Category {
	updatedCategory, _ := repository.Categories.UpdateCategoryById(id, category)
	return updatedCategory
}
