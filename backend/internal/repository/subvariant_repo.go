package repository

import (
	"context"
	"fmt"

	"github.com/fathimasithara01/product-inventory-system/internal/model"
	"github.com/google/uuid"
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

	for i := range subVariants {
		if subVariants[i].ID == uuid.Nil {
			subVariants[i].ID = uuid.New()
		}
	}

	return tx.WithContext(ctx).Create(&subVariants).Error
}

func (r *subVariantRepo) FindByIDForUpdate(ctx context.Context, tx *gorm.DB, id string) (*model.SubVariant, error) {
	if _, err := uuid.Parse(id); err != nil {
		return nil, fmt.Errorf("invalid UUID: %s", id)
	}

	var sub model.SubVariant
	if err := tx.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&sub, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &sub, nil
}

func (r *subVariantRepo) UpdateStock(ctx context.Context, tx *gorm.DB, subVariant *model.SubVariant) error {
	if subVariant.ID == uuid.Nil {
		return fmt.Errorf("invalid UUID: %s", subVariant.ID.String())
	}

	return tx.WithContext(ctx).
		Model(&model.SubVariant{}).
		Where("id = ?", subVariant.ID).
		Update("stock", subVariant.Stock).Error
}
