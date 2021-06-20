package repository

import (
	"errors"
	"ocg.com/product/model/entity"
	"time"
)

type UserRepo struct {
	users map[int64]*entity.User
	autoID int64
}

var Users UserRepo

func init(){
	Users = UserRepo{autoID: 0}
	Users.users = make(map[int64]*entity.User)
}

func (r *UserRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}

func (r *UserRepo) CreateUser(user *entity.User) *entity.User {
	nextID := r.getAutoID()
	user.Id = nextID
	user.CreatedAt = time.Now().Unix()
	user.ModifiedAt = time.Now().Unix()
	r.users[nextID] = user
	return user
}

func (r *UserRepo) FindAllUser() map[int64]*entity.User {
	return r.users
}

func (r *UserRepo) FindUserById(id int64) (*entity.User, error) {
	if user, ok := r.users[id]; ok {
		return user, nil
	} else {
		return nil, errors.New("user not found")
	}
}

func (r *UserRepo) UpdateUserById(id int64, user *entity.User) (*entity.User, error) {
	if _, ok := r.users[id]; ok {
		user.Id = id
		user.CreatedAt = r.users[id].CreatedAt
		user.ModifiedAt = time.Now().Unix()
		r.users[id] = user
		return user, nil
	} else {
		return nil, errors.New("user not found")
	}
}

func (r *UserRepo) DeleteUserById(id int64) error {
	if _, ok := r.users[id]; ok {
		delete(r.users, id)
		return nil
	} else {
		return errors.New("user not found")
	}
}