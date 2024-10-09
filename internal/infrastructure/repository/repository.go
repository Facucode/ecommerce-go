package repository

import (
	"ecommerce-go/internal/core/domain"
	"ecommerce-go/internal/core/ports"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func (r repository) FindAllProducts() ([]domain.Product, error) {
	var products []domain.Product
	r.db.Find(&products)
	return products, nil
}
func (r repository) SaveProduct(product domain.Product) error {
	r.db.Save(&product)
	return nil
}

func (r repository) DeleteProductWithoutStock() int64 {
	result := r.db.Where("stock = ?", "0").Delete(&domain.Product{})

	return result.RowsAffected

}

func NewRepository(db *gorm.DB) ports.Repository {
	repo := repository{
		db: db,
	}
	return &repo
}

/*
func (r *repository) AddProductInCart(productId uint, cartId uint) error {
	cartProduct := domain.CartProduct{CartID: cartId, ProductID: productId}
	r.db.Create(cartProduct)
}*/
