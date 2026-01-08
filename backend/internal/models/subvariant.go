package models

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/shopspring/decimal"
)

type SubVariant struct {
	ID        uuid.UUID       `gorm:"type:uuid;primaryKey"`
	ProductID uuid.UUID       `gorm:"type:uuid;not null;index"`
	OptionIDs pq.StringArray  `gorm:"type:text[]"`
	SKU       string          `gorm:"unique;not null"`
	Stock     decimal.Decimal `gorm:"type:numeric(20,8);default:0"`
}

