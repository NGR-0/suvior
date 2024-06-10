package models

import (
	"time"

	"gorm.io/gorm"
)

type Coupon struct {
	gorm.Model
	Name     string `json:"name" `
	Expiry   time.Time
	Discount string `json:"discount" `
}
