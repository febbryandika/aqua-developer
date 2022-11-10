package usecase

import (
	"e-commerce/entity"
	"e-commerce/entity/response"
	"e-commerce/repository"

	"github.com/jinzhu/copier"
)

type ICartUseCase interface {
	AddProductToCart(req response.AddProductToCartResponse) error
	GetCartList() ([]response.GetCartProductResponse, error)
}

type CartUseCase struct {
	cartRepository repository.ICartRepository
}

func NewCartUseCase(cartRepository repository.ICartRepository) *CartUseCase {
	return &CartUseCase{
		cartRepository: cartRepository,
	}
}

func (c CartUseCase) AddProductToCart(req response.AddProductToCartResponse) error {
	var cart entity.Cart
	copier.Copy(&cart, &req)

	quantity, err := c.cartRepository.GetQuantity(cart)
	if err != nil {
		return err
	}

	quantity += req.Quantity

	if err := c.cartRepository.AddToCart(cart, quantity); err != nil {
		return err
	}
	return nil
}

func (c CartUseCase) GetCartList() ([]response.GetCartProductResponse, error) {
	cart, err := c.cartRepository.GetCart()
	if err != nil {
		return nil, err
	}

	cartResponse := []response.GetCartProductResponse{}
	copier.Copy(&cartResponse, &cart)
	return cartResponse, nil
}
