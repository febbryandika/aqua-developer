package repository

import (
	"e-commerce/entity"

	"gorm.io/gorm"
)

type ICheckoutRepository interface {
	CreateCheckout(id int64) (entity.Checkout, error)
	GetCart() ([]entity.Cart, error)
	ProductCheckout(checkout []entity.CheckoutProduct) error
	ClearCart(id int64) error
}

type CheckoutRepository struct {
	db *gorm.DB
}

func NewCheckoutRepository(db *gorm.DB) *CheckoutRepository {
	return &CheckoutRepository{db: db}
}

func (c CheckoutRepository) CreateCheckout(id int64) (entity.Checkout, error) {
	checkout := entity.Checkout{User_ID: id}
	if err := c.db.Create(&checkout).Error; err != nil {
		return entity.Checkout{}, err
	}
	return checkout, nil
}

func (c CheckoutRepository) GetCart() ([]entity.Cart, error) {
	var cart []entity.Cart
	if err := c.db.Find(&cart).Error; err != nil {
		return nil, err
	}
	return cart, nil
}

func (c CheckoutRepository) ProductCheckout(checkout []entity.CheckoutProduct) error {
	if err := c.db.Create(&checkout).Error; err != nil {
		return err
	}
	return nil
}

func (c CheckoutRepository) ClearCart(id int64) error {
	if err := c.db.Where("user_id = ?", id).Delete(&entity.Cart{}).Error; err != nil {
		return err
	}
	return nil
}
