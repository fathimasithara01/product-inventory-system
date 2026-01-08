package repository

import (
	"context"

	"github.com/fathimasithara01/product-inventory-system/internal/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SubVariantRepository interface {
	Create(ctx context.Context, tx *gorm.DB, subVariants []models.SubVariant) error
	FindByIDForUpdate(ctx context.Context, tx *gorm.DB, id string) (*models.SubVariant, error)
	UpdateStock(ctx context.Context, tx *gorm.DB, subVariant *models.SubVariant) error
}

type subVariantRepository struct {
	db *gorm.DB
}

func NewSubVariantRepository(db *gorm.DB) SubVariantRepository {
	return &subVariantRepository{db: db}
}

func (r *subVariantRepository) Create(ctx context.Context, tx *gorm.DB, subVariants []models.SubVariant) error {
	return tx.WithContext(ctx).Create(&subVariants).Error
}

func (r *subVariantRepository) FindByIDForUpdate(
	ctx context.Context,
	tx *gorm.DB,
	id string,
) (*models.SubVariant, error) {

	var subVariant models.SubVariant
	err := tx.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE"}).
		First(&subVariant, "id = ?", id).Error

	if err != nil {
		return nil, err
	}
	return &subVariant, nil
}

func (r *subVariantRepository) UpdateStock(
	ctx context.Context,
	tx *gorm.DB,
	subVariant *models.SubVariant,
) error {
	return tx.WithContext(ctx).
		Model(&models.SubVariant{}).
		Where("id = ?", subVariant.ID).
		Update("stock", subVariant.Stock).Error
}

// CRITICAL INTERVIEW POINT

// You used:

// Clauses(gorm.Locking{Strength: "UPDATE"})

// 👉 This proves you understand concurrency & race conditions
// 👉 Big + point for backend roles
