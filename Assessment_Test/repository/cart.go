package repository

import (
	"e-commerce/entity"

	"gorm.io/gorm"
)

type ICartRepository interface {
	AddToCart(cart entity.Cart, qty int64) error
	GetQuantity(cart entity.Cart) (int64, error)
	GetCart() ([]entity.Cart, error)
}

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (c CartRepository) AddToCart(cart entity.Cart, qty int64) error {
	result := c.db.FirstOrCreate(&cart, entity.Cart{Product_ID: cart.Product_ID})
	if err := result.Error; err != nil {
		return err
	}

	if result.RowsAffected == 0 {
		if err := c.db.Model(&cart).Updates(entity.Cart{Quantity: qty}).Error; err != nil {
			return err
		}
		return nil
	}

	if err := c.db.Model(&cart).Updates(entity.Cart{Quantity: cart.Quantity}).Error; err != nil {
		return err
	}
	return nil
}

func (c CartRepository) GetQuantity(cart entity.Cart) (int64, error) {
	if err := c.db.Find(&cart, "product_id = ?", cart.Product_ID).Error; err != nil {
		return 0, err
	}
	return cart.Quantity, nil
}

func (c CartRepository) GetCart() ([]entity.Cart, error) {
	var cartProduct []entity.Cart
	if err := c.db.Joins("Product").Find(&cartProduct).Error; err != nil {
		return nil, err
	}
	return cartProduct, nil
}
