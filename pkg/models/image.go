package models

import (
	"gorm.io/gorm"
)

type ImageProd struct {
	gorm.Model
	ProductID uint   `json:"productID" `
	Url       string `json:"url" `
}
type ImageBlog struct {
	gorm.Model
	BlogID uint   `json:"blogID" `
	Url    string `json:"url" `
}
