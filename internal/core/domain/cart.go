package domain

type Cart struct {
	ID          uint
	TotalAmount uint16
	UserID      uint
	User        User
}
