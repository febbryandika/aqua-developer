package usecase

import (
	"e-commerce/entity"
	"e-commerce/entity/response"
	"e-commerce/repository"

	"github.com/jinzhu/copier"
)

type ICheckoutUseCase interface {
	CheckoutCart(id int64) error
}

type CheckoutUseCase struct {
	checkoutRepository repository.ICheckoutRepository
}

func NewCheckoutUseCase(checkoutRepository repository.ICheckoutRepository) *CheckoutUseCase {
	return &CheckoutUseCase{
		checkoutRepository: checkoutRepository,
	}
}

func (c CheckoutUseCase) CheckoutCart(id int64) error {
	checkout, err := c.checkoutRepository.CreateCheckout(id)
	if err != nil {
		return err
	}

	carts, err := c.checkoutRepository.GetCart()
	if err != nil {
		return err
	}

	var checkoutProduct []response.CheckoutCartResponse
	for _, cart := range carts {
		data := response.CheckoutCartResponse{Checkout_ID: checkout.ID, Product_ID: cart.Product_ID, Quantity: cart.Quantity}
		checkoutProduct = append(checkoutProduct, data)
	}

	var checkouts []entity.CheckoutProduct
	copier.Copy(&checkouts, checkoutProduct)

	if err := c.checkoutRepository.ProductCheckout(checkouts); err != nil {
		return err
	}

	if err := c.checkoutRepository.ClearCart(id); err != nil {
		return err
	}
	return nil
}
