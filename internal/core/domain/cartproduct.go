package domain

type CartProduct struct {
	CartID    uint
	ProductID uint
	Quantity  uint64

	Cart    Cart
	Product Product
}
