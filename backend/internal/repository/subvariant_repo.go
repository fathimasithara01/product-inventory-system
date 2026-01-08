package repository

import (
	"context"

	"github.com/fathimasithara01/product-inventory-system/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SubVariantRepository interface {
	Create(ctx context.Context, tx *gorm.DB, subVariants []model.SubVariant) error
	FindByIDForUpdate(ctx context.Context, tx *gorm.DB, id string) (*model.SubVariant, error)
	UpdateStock(ctx context.Context, tx *gorm.DB, subVariant *model.SubVariant) error
}

type subVariantRepo struct {
	db *gorm.DB
}

func NewSubVariantRepository(db *gorm.DB) SubVariantRepository {
	return &subVariantRepo{db: db}
}

func (r *subVariantRepo) Create(ctx context.Context, tx *gorm.DB, subVariants []model.SubVariant) error {
	if len(subVariants) == 0 {
		return nil
	}
	return tx.WithContext(ctx).Create(&subVariants).Error
}

func (r *subVariantRepo) FindByIDForUpdate(ctx context.Context, tx *gorm.DB, id string) (*model.SubVariant, error) {
	var sub model.SubVariant
	if err := tx.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&sub, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &sub, nil
}

func (r *subVariantRepo) UpdateStock(ctx context.Context, tx *gorm.DB, subVariant *model.SubVariant) error {
	return tx.WithContext(ctx).
		Model(&model.SubVariant{}).
		Where("id = ?", subVariant.ID).
		Update("stock", subVariant.Stock).Error
}
