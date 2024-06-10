package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	OrderStatus string `gorm:"default:'Not Processed" json:"orderStatus"`
	OrderByID   uint   `json:"orderByID" `
	OrderBy     User   `gorm:"foreignKey:OrderByID;"`
}

type ProductOrder struct {
	gorm.Model
	OrderID   uint    `json:"orderID" `
	Order     Order   `gorm:"foreignKey:OrderID;"`
	ProductID uint    `json:"productID" `
	Product   Product `gorm:"foreignKey:ProductID;"`
	Quantity  string  `json:"quantity" `
	Color     string  `json:"color" `
}

const (
	NotProcessed   = "Not Processed"
	CashOnDelivery = "Cash on Delivery"
	Processing     = "Processing"
	Dispatched     = "Dispatched"
	Cancelled      = "Cancelled"
	Delivered      = "Delivered"
)
