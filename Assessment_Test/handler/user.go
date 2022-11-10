package handler

import (
	"e-commerce/entity/response"
	"e-commerce/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUserCase usecase.IUserUseCase
}

func NewUserHandler(userUseCase usecase.IUserUseCase) *UserHandler {
	return &UserHandler{
		userUserCase: userUseCase,
	}
}

func (h UserHandler) RegisterNewUser(c echo.Context) error {
	user := response.CreateUserResponse{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to register user",
			Data:    nil,
		})
	}
	if err := h.userUserCase.RegisterUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to register user",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "New user created successfully",
		Data:    nil,
	})
}

func (h UserHandler) LoginAdmin(c echo.Context) error {
	req := response.LoginUserResponse{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to login user",
			Data:    nil,
		})
	}
	admin, err := h.userUserCase.LoginAdmin(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to login admin",
			Data:    nil,
		})
	}
	if admin.AccessToken == "" {
		return c.JSON(http.StatusUnauthorized, response.BaseResponse{
			Code:    http.StatusUnauthorized,
			Message: "User is not authorized",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Login successful",
		Data:    admin.AccessToken,
	})
}

func (h UserHandler) LoginUser(c echo.Context) error {
	req := response.LoginUserResponse{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Failed to login user",
			Data:    nil,
		})
	}
	user, err := h.userUserCase.LoginUser(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to login admin",
			Data:    nil,
		})
	}
	if user.AccessToken == "" {
		return c.JSON(http.StatusUnauthorized, response.BaseResponse{
			Code:    http.StatusUnauthorized,
			Message: "User is not authorized",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Login successful",
		Data:    user.AccessToken,
	})
}
