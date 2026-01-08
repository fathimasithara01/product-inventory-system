package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type StockTransaction struct {
	ID              uuid.UUID       `gorm:"type:uuid;primaryKey"`
	ProductID       uuid.UUID       `gorm:"type:uuid;index;not null"`
	SubVariantID    uuid.UUID       `gorm:"type:uuid;index;not null"`
	Quantity        decimal.Decimal `gorm:"type:numeric(20,8);not null"`
	TransactionType string          `gorm:"type:varchar(20);not null"` // IN | OUT
	TransactionDate time.Time       `gorm:"autoCreateTime"`
}
