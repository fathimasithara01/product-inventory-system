package models

import "github.com/google/uuid"

type VariantOption struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	VariantID uuid.UUID `gorm:"type:uuid;not null;index"`
	Value     string    `gorm:"not null"` // eg: Red, Blue, XL
}


