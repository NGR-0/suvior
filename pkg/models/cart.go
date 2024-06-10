package models

import (
	"gorm.io/gorm"
)

type Cart struct {
	gorm.Model
	Products           []CartItem `gorm:"foreignKey:CartID;"`
	CartTotal          string     `json:"cartTotal" `
	TotalAfterDiscount string     `json:"totalAfterDiscount" `
	OrderByID          uint       `json:"orderByID" `
	OrderBy            User       `gorm:"foreignKey:OrderByID;"`
}
type CartItem struct {
	gorm.Model
	CartID    uint    `json:"cartID" `
	ProductID uint    `json:"productID" `
	Products  Product `gorm:"foreignKey:ProductID;"`
	Color     string  `json:"color" `
	Price     string  `json:"price" `
	Quantity  string  `json:"quantity" `
}
