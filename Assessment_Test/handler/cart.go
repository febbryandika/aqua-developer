package handler

import (
	"e-commerce/entity/response"
	"e-commerce/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartUseCase usecase.ICartUseCase
}

func NewCartHandler(cartUseCase usecase.ICartUseCase) *CartHandler {
	return &CartHandler{
		cartUseCase: cartUseCase,
	}
}

func (h CartHandler) AddToCart(c echo.Context) error {
	cart := response.AddProductToCartResponse{}
	if err := c.Bind(&cart); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Add product to cart failed",
			Data:    nil,
		})
	}
	if err := h.cartUseCase.AddProductToCart(cart); err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to add product to cart",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully add product to cart",
		Data:    nil,
	})
}

func (h CartHandler) GetCart(c echo.Context) error {
	cart, err := h.cartUseCase.GetCartList()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Get cart failed",
			Data:    nil,
		})
	}
	if len(cart) <= 0 {
		return c.JSON(http.StatusNotFound, response.BaseResponse{
			Code:    http.StatusNotFound,
			Message: "Empty cart",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully get cart",
		Data:    cart,
	})
}
