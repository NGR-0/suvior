package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName            string    `json:"firstName" validate:"required,max=50"`
	LastName             string    `json:"lastName" validate:"required,max=50"`
	Email                string    `json:"email" validate:"required,max=50"`
	Mobile               string    `json:"mobile"`
	Password             string    `json:"password" validate:"required,max=50"`
	Role                 string    `json:"role"`
	IsBlocked            bool      `json:"isBlocked"`
	Cart                 []Cart    `gorm:"foreignKey:OrderByID"`
	Address              string    `json:"address" validate:"required,max=100"`
	Wishlist             []Product `gorm:"many2many:user_wishlist;"`
	Likes                []Blog    `gorm:"many2many:user_likes;"`
	Dislikes             []Blog    `gorm:"many2many:user_dislikes;"`
	RefreshToken         string    `json:"refreshToken" validate:"required,max=100"`
	PasswordChangedAt    time.Time
	PasswordResetToken   string `json:"passwordResetToken" validate:"required,max=100"`
	PasswordResetExpires time.Time
}
