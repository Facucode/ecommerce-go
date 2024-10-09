package ports

import "ecommerce-go/internal/core/domain"

type Repository interface {
	FindAllProducts() ([]domain.Product, error)
	SaveProduct(domain.Product) error
	DeleteProductWithoutStock() int64
	//AddProductInCart(uint, uint) error
	//FindProductById(string) (domain.Product, error)
	//FindUserById(string) (domain.User, error)
	//FindAllUsers() ([]domain.User, error)
}
