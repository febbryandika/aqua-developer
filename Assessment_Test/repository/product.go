package repository

import (
	"e-commerce/entity"

	"gorm.io/gorm"
)

type IProductRepository interface {
	Create(product entity.Product) error
	GetProducts() ([]entity.Product, error)
	GetById(id int64) error
	GetProduct(id int64) (*entity.Product, error)
	FilterCategory(category string) ([]entity.Product, error)
	FilterPrice(min, max int64) ([]entity.Product, error)
}

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p ProductRepository) Create(product entity.Product) error {
	if err := p.db.Create(&product).Error; err != nil {
		return err
	}
	return nil
}

func (p ProductRepository) GetProducts() ([]entity.Product, error) {
	var products []entity.Product
	if err := p.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (u ProductRepository) GetById(id int64) error {
	var product entity.Product
	if err := u.db.First(&product, id).Error; err != nil {
		return err
	}
	return nil
}

func (p ProductRepository) GetProduct(id int64) (*entity.Product, error) {
	var product entity.Product
	if err := p.db.Find(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (p ProductRepository) FilterCategory(category string) ([]entity.Product, error) {
	var product []entity.Product
	if err := p.db.Where("category = ?", category).Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (p ProductRepository) FilterPrice(min, max int64) ([]entity.Product, error) {
	var product []entity.Product
	if err := p.db.Where("price BETWEEN ? AND ?", min, max).Find(&product).Error; err != nil {
		return nil, err
	}
	return product, nil
}
