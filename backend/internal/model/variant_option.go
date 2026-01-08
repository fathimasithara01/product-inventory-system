package model

import "github.com/google/uuid"

type VariantOption struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	VariantID uuid.UUID `gorm:"type:uuid;index;not null"`
	Value     string    `gorm:"not null"`
}
