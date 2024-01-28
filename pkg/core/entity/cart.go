package entity

import (
	"gorm.io/gorm"
)

const (
	CartOpen   = "open"
	CartClosed = "closed"
)

type Cart struct {
	gorm.Model
	Total     float64
	SessionID string
	Status    string
}
