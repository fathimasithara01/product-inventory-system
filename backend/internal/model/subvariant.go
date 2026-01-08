package model

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/shopspring/decimal"
)

type SubVariant struct {
	ID        uuid.UUID       `gorm:"type:uuid;primaryKey"`
	ProductID uuid.UUID       `gorm:"type:uuid;index;not null"`
	OptionIDs pq.StringArray  `gorm:"type:text[];not null"`
	SKU       string          `gorm:"uniqueIndex;not null"`
	Stock     decimal.Decimal `gorm:"type:numeric(20,8);default:0"`
}
