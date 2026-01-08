package model

import "github.com/google/uuid"

type Variant struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID uuid.UUID `gorm:"type:uuid;index;not null"`
	Name      string    `gorm:"not null"`

	Options []VariantOption `gorm:"foreignKey:VariantID"`
}
