package handler

import (
	"e-commerce/entity/response"
	"e-commerce/usecase"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CheckoutHandler struct {
	checkoutUseCase usecase.ICheckoutUseCase
}

func NewCheckoutHandler(checkoutUseCase usecase.ICheckoutUseCase) *CheckoutHandler {
	return &CheckoutHandler{
		checkoutUseCase: checkoutUseCase,
	}
}

func (h CheckoutHandler) CheckoutProduct(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(file.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}
	id := c.Param("id")
	userId, _ := strconv.ParseInt(id, 0, 64)
	if err := h.checkoutUseCase.CheckoutCart(userId); err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed to checkout",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully checkout cart",
		Data:    nil,
	})
}
