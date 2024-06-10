package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title       string       `json:"title" validate:"required,max=50"`
	Slug        string       `json:"slug" validate:"required,max=50"`
	Description string       `json:"description" validate:"required,max=100"`
	Price       int          `json:"price" `
	Quantity    int          `json:"quantity" `
	CategoryID  string       `json:"categoryID"`
	Category    CategoryProd `gorm:"foreignKey:ProductID;"`
	Brand       string       `json:"brand" `
	Sold        int          `json:"sold" `
	Image       []ImageProd  `gorm:"foreignKey:ProductID;"`
	Color       string       `json:"color" `
	Tags        string       `json:"tags" `
	Ratings     RatingProd   `gorm:"foreignKey:ProductID"`
	TotalRating string       `json:"totalRating" `
}

type RatingProd struct {
	gorm.Model
	ProductID  uint   `json:"productID" `
	Name       string `json:"name" `
	Comment    string `json:"comment" `
	Value      string `json:"value" `
	PostedByID uint   `json:"postedByID" `
	PostedBy   User   `gorm:"foreignKey:PostedByID;"`
}
type CategoryProd struct {
	gorm.Model
	ProductID uint   `json:"productID" `
	Name      string `json:"name" `
}
