package usecase

import (
	"user-management/entity"
	"user-management/entity/response"
	"user-management/repository"

	"github.com/jinzhu/copier"
)

type IUseCase interface {
	CreateUser(user response.CreateUserRequest) error
	GetListUser() ([]response.GetUserResponse, error)
	UpdateUser(user response.UpdateUserResponse) error
	DeleteUser(id int64) error
}

type UserUseCase struct {
	userRepository repository.IUserRepository
}

func NewUserUseCase(userRepository repository.IUserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (u UserUseCase) CreateUser(req response.CreateUserRequest) error {
	user := entity.User{}
	copier.Copy(&user, &req)

	if err := u.userRepository.Create(user); err != nil {
		return err
	}
	return nil
}

func (u UserUseCase) GetListUser() ([]response.GetUserResponse, error) {
	users, err := u.userRepository.GetAll()
	if err != nil {
		return nil, err
	}
	userResponse := []response.GetUserResponse{}
	copier.Copy(&userResponse, &users)
	return userResponse, nil
}

func (u UserUseCase) UpdateUser(req response.UpdateUserResponse, id int64) error {
	if err := u.userRepository.GetById(id); err != nil {
		return err
	}

	user := entity.User{ID: id}
	copier.Copy(&user, &req)

	if err := u.userRepository.Update(user, id); err != nil {
		return err
	}
	return nil
}

func (u UserUseCase) DeleteUser(id int64) error {
	if err := u.userRepository.GetById(id); err != nil {
		return err
	}
	if err := u.userRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
