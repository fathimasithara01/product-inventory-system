package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type StockTransaction struct {
	ID              uuid.UUID       `gorm:"type:uuid;primaryKey"`
	ProductID       uuid.UUID       `gorm:"type:uuid;not null;index"`
	SubVariantID    uuid.UUID       `gorm:"type:uuid;not null;index"`
	Quantity        decimal.Decimal `gorm:"type:numeric(20,8);not null"`
	TransactionType string          `gorm:"not null"`
	TransactionDate time.Time       `gorm:"autoCreateTime"`
}
