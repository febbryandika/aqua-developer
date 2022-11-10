package repository

import (
	"e-commerce/entity"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Register(user entity.User) error
	Login(username string, password string) (entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (u UserRepository) Register(user entity.User) error {
	if err := u.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (u UserRepository) Login(username string, password string) (entity.User, error) {
	var user entity.User
	if err := u.db.Where("username = ? AND password = ?", username, password).Find(&user).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}
