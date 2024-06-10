package models

import "gorm.io/gorm"

type Enq struct {
	gorm.Model
	Name    string `json:"name" `
	Email   string `json:"email" `
	Mobile  string `json:"mobile" `
	Comment string `json:"comment" `
	Status  string `json:"status" `
}

const (
	Submitted  = "Submitted"
	Contacted  = "Contacted"
	InProgress = "In Progress"
	Resolved   = "Resolved"
)
