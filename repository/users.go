package repository

import (
	"goLang-userManage79/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user model.Users) (model.Users, error)
	GetUsers(offset int, limit int) ([]model.Users, error)
	GetUser(ID int) (model.Users, error)
	UpdateUser(user model.Users) (model.Users, error)
	DeleteUser(ID int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUsers(offset int, limit int) ([]model.Users, error) {
	var users []model.Users
	err := r.db.Offset(offset).Limit(limit).Find(&users).Error
	return users, err
}

func (r *userRepository) CreateUser(user model.Users) (model.Users, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *userRepository) GetUser(ID int) (model.Users, error) {
	var user model.Users
	err := r.db.Find(&user, ID).Error
	return user, err
}

func (r *userRepository) UpdateUser(user model.Users) (model.Users, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *userRepository) DeleteUser(ID int) error {
	err := r.db.Delete(&model.Users{}, ID).Error
	return err
}
