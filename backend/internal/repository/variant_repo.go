package repository

import (
	"context"

	"github.com/fathimasithara01/product-inventory-system/internal/models"
	"gorm.io/gorm"
)

type VariantRepository interface {
	Create(ctx context.Context, tx *gorm.DB, variant *models.Variant) error
	CreateOptions(ctx context.Context, tx *gorm.DB, options []models.VariantOption) error
}

type variantRepository struct {
	db *gorm.DB
}

func NewVariantRepository(db *gorm.DB) VariantRepository {
	return &variantRepository{db: db}
}

func (r *variantRepository) Create(ctx context.Context, tx *gorm.DB, variant *models.Variant) error {
	return tx.WithContext(ctx).Create(variant).Error
}

func (r *variantRepository) CreateOptions(ctx context.Context, tx *gorm.DB, options []models.VariantOption) error {
	if len(options) == 0 {
		return nil
	}
	return tx.WithContext(ctx).Create(&options).Error
}
