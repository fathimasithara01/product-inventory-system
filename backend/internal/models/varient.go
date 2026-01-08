package models

import "github.com/google/uuid"

type Variant struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID uuid.UUID `gorm:"type:uuid;not null;index"`
	Name      string    `gorm:"not null"` // eg: Color, Size
}
