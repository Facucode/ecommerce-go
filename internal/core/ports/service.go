package ports

import (
	"ecommerce-go/internal/core/domain"
)

type EcommerceService interface {
	GetProducts() ([]domain.Product, error)
	//GetProduct(string) (domain.Product, error)
	PostProduct(domain.Product) error
	DeleteProductWithoutStock() int64
	//GetUsers() ([]domain.User, error)
}
