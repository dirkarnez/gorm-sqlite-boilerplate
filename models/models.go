package models

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)
type User struct {
	gorm.Model
	Name   string
	Saving decimal.Decimal `gorm:"type:decimal(13,6);"`
}
