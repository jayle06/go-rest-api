package service

import (
	"ocg.com/product/model/entity"
	"ocg.com/product/repository"
)

func CreateUser(user *entity.User) *entity.User {
	return repository.Users.CreateUser(user)
}

func GetAllUser() map[int64]*entity.User {
	return repository.Users.FindAllUser()
}

func GetUserById(id int64) *entity.User {
	user, _ := repository.Users.FindUserById(id)
	return user
}

func DeleteUserById(id int64) {
	_ = repository.Users.DeleteUserById(id)
}

func UpdateUserById(id int64, user *entity.User) *entity.User {
	updatedUser, _ := repository.Users.UpdateUserById(id, user)
	return updatedUser
}
