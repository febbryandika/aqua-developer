package usecase

import (
	"e-commerce/entity"
	"e-commerce/entity/response"
	"e-commerce/repository"
	"e-commerce/token"

	"github.com/jinzhu/copier"
)

type IUserUseCase interface {
	RegisterUser(req response.CreateUserResponse) error
	LoginAdmin(req response.LoginUserResponse) (response.AccessTokenResponse, error)
	LoginUser(req response.LoginUserResponse) (response.AccessTokenResponse, error)
}

type UserUseCase struct {
	userRepository repository.IUserRepository
}

func NewUserUseCase(userRepository repository.IUserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (u UserUseCase) RegisterUser(req response.CreateUserResponse) error {
	user := entity.User{}
	copier.Copy(&user, &req)

	if err := u.userRepository.Register(user); err != nil {
		return err
	}
	return nil
}

func (u UserUseCase) LoginAdmin(req response.LoginUserResponse) (response.AccessTokenResponse, error) {
	user, err := u.userRepository.Login(req.Username, req.Password)
	if err != nil {
		return response.AccessTokenResponse{}, err
	}

	var at response.AccessTokenResponse
	if user.Username == "admin" {
		token, _ := token.GenerateJWT(user.Username, true)
		at.AccessToken = token
		return at, nil
	}
	return response.AccessTokenResponse{}, err
}

func (u UserUseCase) LoginUser(req response.LoginUserResponse) (response.AccessTokenResponse, error) {
	user, err := u.userRepository.Login(req.Username, req.Password)
	if err != nil {
		return response.AccessTokenResponse{}, err
	}

	var at response.AccessTokenResponse
	if user.Username == req.Username {
		token, _ := token.GenerateJWT(user.Username, true)
		at.AccessToken = token
		return at, nil
	}
	return response.AccessTokenResponse{}, err
}
