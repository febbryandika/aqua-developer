package handler

import (
	"e-commerce/entity/response"
	"e-commerce/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUseCase usecase.IProductUseCase
}

func NewProductHandler(productUseCase usecase.IProductUseCase) *ProductHandler {
	return &ProductHandler{
		productUseCase: productUseCase,
	}
}

func (h ProductHandler) CreateProduct(c echo.Context) error {
	product := response.CreateProductResponse{}
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Create product failed",
			Data:    nil,
		})
	}
	if err := h.productUseCase.CreateProduct(product); err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Create product failed",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusCreated,
		Message: "Product successfully created",
		Data:    nil,
	})
}

func (h ProductHandler) GetAllProducts(c echo.Context) error {
	category := c.QueryParam("category")
	min := c.QueryParam("min")
	max := c.QueryParam("max")
	if len(category) == 0 && len(min) == 0 && len(max) == 0 {
		products, err := h.productUseCase.GetProductsList()
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BaseResponse{
				Code:    http.StatusBadRequest,
				Message: "Get products list failed",
				Data:    nil,
			})
		}
		if len(products) <= 0 {
			return c.JSON(http.StatusNotFound, response.BaseResponse{
				Code:    http.StatusNotFound,
				Message: "Products not found",
				Data:    nil,
			})
		}
		return c.JSON(http.StatusCreated, response.BaseResponse{
			Code:    http.StatusOK,
			Message: "Successfully get all products",
			Data:    products,
		})
	}
	if len(min) == 0 && len(max) == 0 {
		product, err := h.productUseCase.GetFilterCategory(category)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BaseResponse{
				Code:    http.StatusBadRequest,
				Message: "Get filtered product failed",
				Data:    nil,
			})
		}

		if len(product) <= 0 {
			return c.JSON(http.StatusNotFound, response.BaseResponse{
				Code:    http.StatusNotFound,
				Message: "Product not found",
				Data:    nil,
			})
		}

		return c.JSON(http.StatusOK, response.BaseResponse{
			Code:    http.StatusOK,
			Message: "Successfully get filtered product",
			Data:    product,
		})
	} else {
		minPrice, _ := strconv.ParseInt(min, 0, 64)
		maxPrice, _ := strconv.ParseInt(max, 0, 64)
		product, err := h.productUseCase.GetFilterPrice(minPrice, maxPrice)
		if err != nil {
			return c.JSON(http.StatusBadRequest, response.BaseResponse{
				Code:    http.StatusBadRequest,
				Message: "Get filtered product failed",
				Data:    nil,
			})
		}

		if len(product) <= 0 {
			return c.JSON(http.StatusNotFound, response.BaseResponse{
				Code:    http.StatusNotFound,
				Message: "Product not found",
				Data:    nil,
			})
		}

		return c.JSON(http.StatusOK, response.BaseResponse{
			Code:    http.StatusOK,
			Message: "Successfully get filtered product",
			Data:    product,
		})
	}
}

func (h ProductHandler) GetProductDetail(c echo.Context) error {
	id := c.Param("id")
	productId, _ := strconv.ParseInt(id, 0, 64)
	product, err := h.productUseCase.GetProduct(productId)
	if productId <= 0 {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid product ID",
			Data:    nil,
		})
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Get product detail failed",
			Data:    nil,
		})
	}
	return c.JSON(http.StatusCreated, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Successfully get product detail",
		Data:    product,
	})
}
