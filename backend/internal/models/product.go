package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductCode  string    `gorm:"unique;not null"`
	ProductName  string    `gorm:"not null"`
	ProductImage string
	CreatedDate  time.Time `gorm:"autoCreateTime"`
	UpdatedDate  time.Time
	CreatedUser  uuid.UUID `gorm:"type:uuid"`
	IsFavourite  bool      `gorm:"default:false"`
	Active       bool      `gorm:"default:true"`
	HSNCode      string
	TotalStock   decimal.Decimal `gorm:"type:numeric(20,8);default:0"`
}
