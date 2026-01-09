package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	ProductID    int64     `gorm:"autoIncrement;uniqueIndex"`
	ProductCode  string    `gorm:"uniqueIndex;not null"`
	ProductName  string    `gorm:"not null"`
	ProductImage string
	CreatedDate  time.Time `gorm:"autoCreateTime"`
	UpdatedDate  time.Time `gorm:"autoUpdateTime"`
	CreatedUser  uuid.UUID `gorm:"type:uuid;not null"`
	IsFavourite  bool      `gorm:"default:false"`
	Active       bool      `gorm:"default:true"`
	HSNCode      string
	TotalStock   decimal.Decimal `gorm:"type:numeric(20,8);default:0"`

	Variants    []Variant    `gorm:"foreignKey:ProductID"`
	SubVariants []SubVariant `gorm:"foreignKey:ProductID"`
}
