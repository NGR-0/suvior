package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title       string      `json:"title" `
	Description string      `json:"description" `
	CategoryID  uint        `json:"categoryID" `
	Category    BlogCat     `gorm:"foreignKey:CategoryID;"`
	NumViews    string      `json:"numViews" `
	IsLiked     bool        `json:"isLiked" `
	AuthorID    uint        `json:"authorID" `
	Author      User        `gorm:"foreignKey:AuthorID;"`
	Images      []ImageBlog `gorm:"foreignKey:BlogID;"`
}
type BlogCat struct {
	gorm.Model
	Name string `json:"name" `
}
