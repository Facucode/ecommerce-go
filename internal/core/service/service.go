package service

import (
	"ecommerce-go/internal/core/domain"
	"ecommerce-go/internal/core/ports"
)

type ecommerce_service struct {
	repository ports.Repository
}

func (e ecommerce_service) GetProducts() ([]domain.Product, error) {
	products, err := e.repository.FindAllProducts()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (e ecommerce_service) PostProduct(product domain.Product) error {
	err := e.repository.SaveProduct(product)
	if err != nil {
		return err
	}
	return nil
}

func (e ecommerce_service) DeleteProductWithoutStock() int64 {
	rowsAffected := e.repository.DeleteProductWithoutStock()

	return rowsAffected
}

func NewEcommerceService(repository ports.Repository) ports.EcommerceService {
	ecommerce_service := ecommerce_service{
		repository: repository,
	}
	return &ecommerce_service
}

/*
func (e ecommerce_service) AddProductInCart(product domain.Product, cart domain.Cart) error {
	productId := product.ID
	cartId := cart.ID
	err := e.repository.AddProductInCart(productId, cartId)
	if err != nil {

	}
}
*/
