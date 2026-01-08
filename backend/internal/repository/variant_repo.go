package repository

import (
	"context"

	"github.com/fathimasithara01/product-inventory-system/internal/model"
	"gorm.io/gorm"
)

type VariantRepository interface {
	Create(ctx context.Context, tx *gorm.DB, variant *model.Variant) error
	CreateOptions(ctx context.Context, tx *gorm.DB, options []model.VariantOption) error
}

type variantRepo struct {
	db *gorm.DB
}

func NewVariantRepository(db *gorm.DB) VariantRepository {
	return &variantRepo{db: db}
}

func (r *variantRepo) Create(ctx context.Context, tx *gorm.DB, variant *model.Variant) error {
	return tx.WithContext(ctx).Create(variant).Error
}

func (r *variantRepo) CreateOptions(ctx context.Context, tx *gorm.DB, options []model.VariantOption) error {
	if len(options) == 0 {
		return nil
	}
	return tx.WithContext(ctx).Create(&options).Error
}
