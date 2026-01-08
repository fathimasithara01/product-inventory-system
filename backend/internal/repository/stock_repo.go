package repository

import (
	"context"
	"time"

	"github.com/fathimasithara01/product-inventory-system/internal/models"
	"gorm.io/gorm"
)

type StockRepository interface {
	CreateTransaction(ctx context.Context, tx *gorm.DB, txn *models.StockTransaction) error
	GetReport(ctx context.Context, from, to time.Time) ([]models.StockTransaction, error)
}

type stockRepository struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &stockRepository{db: db}
}

func (r *stockRepository) CreateTransaction(
	ctx context.Context,
	tx *gorm.DB,
	txn *models.StockTransaction,
) error {
	return tx.WithContext(ctx).Create(txn).Error
}

func (r *stockRepository) GetReport(
	ctx context.Context,
	from, to time.Time,
) ([]models.StockTransaction, error) {

	var txns []models.StockTransaction

	err := r.db.WithContext(ctx).
		Where("transaction_date BETWEEN ? AND ?", from, to).
		Order("transaction_date DESC").
		Find(&txns).Error

	return txns, err
}
