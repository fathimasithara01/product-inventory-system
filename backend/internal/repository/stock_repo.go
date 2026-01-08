package repository

import (
	"context"
	"time"

	"github.com/fathimasithara01/product-inventory-system/internal/model"
	"gorm.io/gorm"
)

type StockRepository interface {
	CreateTransaction(ctx context.Context, tx *gorm.DB, txn *model.StockTransaction) error
	GetReport(ctx context.Context, from, to time.Time) ([]model.StockTransaction, error)
}

type stockRepo struct {
	db *gorm.DB
}

func NewStockRepository(db *gorm.DB) StockRepository {
	return &stockRepo{db: db}
}

func (r *stockRepo) CreateTransaction(ctx context.Context, tx *gorm.DB, txn *model.StockTransaction) error {
	return tx.WithContext(ctx).Create(txn).Error
}

func (r *stockRepo) GetReport(ctx context.Context, from, to time.Time) ([]model.StockTransaction, error) {
	var txns []model.StockTransaction
	if err := r.db.WithContext(ctx).
		Where("transaction_date BETWEEN ? AND ?", from, to).
		Order("transaction_date DESC").
		Find(&txns).Error; err != nil {
		return nil, err
	}
	return txns, nil
}
