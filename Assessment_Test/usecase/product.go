package usecase

import (
	"e-commerce/entity"
	"e-commerce/entity/response"
	"e-commerce/repository"

	"github.com/jinzhu/copier"
)

type IProductUseCase interface {
	CreateProduct(req response.CreateProductResponse) error
	GetProductsList() ([]response.GetAllProductResponse, error)
	GetProduct(id int64) (*response.GetDetailedProductResponse, error)
	GetFilterCategory(category string) ([]response.GetAllProductResponse, error)
	GetFilterPrice(min, max int64) ([]response.GetAllProductResponse, error)
}

type ProductUseCase struct {
	productRepository repository.IProductRepository
}

func NewProductUseCase(productRepository repository.IProductRepository) *ProductUseCase {
	return &ProductUseCase{
		productRepository: productRepository,
	}
}

func (p ProductUseCase) CreateProduct(req response.CreateProductResponse) error {
	product := entity.Product{}
	copier.Copy(&product, &req)

	if err := p.productRepository.Create(product); err != nil {
		return err
	}
	return nil
}

func (p ProductUseCase) GetProductsList() ([]response.GetAllProductResponse, error) {
	products, err := p.productRepository.GetProducts()
	if err != nil {
		return nil, err
	}

	productResponse := []response.GetAllProductResponse{}
	copier.Copy(&productResponse, &products)
	return productResponse, nil
}

func (p ProductUseCase) GetProduct(id int64) (*response.GetDetailedProductResponse, error) {
	if err := p.productRepository.GetById(id); err != nil {
		return nil, err
	}

	product, err := p.productRepository.GetProduct(id)
	if err != nil {
		return nil, err
	}

	productResponse := response.GetDetailedProductResponse{}
	copier.Copy(&productResponse, &product)
	return &productResponse, nil
}

func (p ProductUseCase) GetFilterCategory(category string) ([]response.GetAllProductResponse, error) {
	product, err := p.productRepository.FilterCategory(category)
	if err != nil {
		return nil, err
	}
	productResponse := []response.GetAllProductResponse{}
	copier.Copy(&productResponse, &product)
	return productResponse, nil
}

func (p ProductUseCase) GetFilterPrice(min, max int64) ([]response.GetAllProductResponse, error) {
	product, err := p.productRepository.FilterPrice(min, max)
	if err != nil {
		return nil, err
	}
	productResponse := []response.GetAllProductResponse{}
	copier.Copy(&productResponse, &product)
	return productResponse, nil
}
