package handler

import (
	"user-management/entity/response"
	"user-management/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(usercase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usercase,
	}
}

func (h UserHandler) CreateUser(ctx *fiber.Ctx) error {
	userRequest := response.CreateUserRequest{}
	if err := ctx.BodyParser(&userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
	}
	if err := h.userUseCase.CreateUser(userRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Invalid request body",
			Data:    nil,
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(response.BaseResponse{
		Code:    fiber.StatusCreated,
		Message: "User created successfully",
		Data:    nil,
	})
}

func (h UserHandler) GetList(ctx *fiber.Ctx) error {
	users, err := h.userUseCase.GetListUser()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Get user list failed",
			Data:    nil,
		})
	}
	if len(users) <= 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.BaseResponse{
			Code:    fiber.StatusBadRequest,
			Message: "Users not found",
			Data:    nil,
		})
	}
	return ctx.Status(fiber.StatusCreated).JSON(response.BaseResponse{
		Code:    fiber.StatusCreated,
		Message: "Successfully get all users",
		Data:    users,
	})
}
