package service

import (
	"context"
	"errors"

	"github.com/fathimasithara01/product-inventory-system/internal/model"
	"github.com/fathimasithara01/product-inventory-system/internal/repository"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type StockService interface {
	AddStock(ctx context.Context, productID, subVariantID string, qty decimal.Decimal) error
	RemoveStock(ctx context.Context, productID, subVariantID string, qty decimal.Decimal) error
}

type stockService struct {
	db          *gorm.DB
	subRepo     repository.SubVariantRepository
	stockRepo   repository.StockRepository
	productRepo repository.ProductRepository
}

func NewStockService(
	db *gorm.DB,
	subRepo repository.SubVariantRepository,
	stockRepo repository.StockRepository,
	productRepo repository.ProductRepository,
) StockService {
	return &stockService{
		db:          db,
		subRepo:     subRepo,
		stockRepo:   stockRepo,
		productRepo: productRepo,
	}
}

func (s *stockService) AddStock(
	ctx context.Context,
	productID, subVariantID string,
	qty decimal.Decimal,
) error {

	if qty.LessThanOrEqual(decimal.Zero) {
		return errors.New("quantity must be greater than zero")
	}

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		sub, err := s.subRepo.FindByIDForUpdate(ctx, tx, subVariantID)
		if err != nil {
			return err
		}

		sub.Stock = sub.Stock.Add(qty)

		if err := s.subRepo.UpdateStock(ctx, tx, sub); err != nil {
			return err
		}

		txn := &model.StockTransaction{
			ID:              uuid.New(),
			ProductID:       sub.ProductID,
			SubVariantID:    sub.ID,
			Quantity:        qty,
			TransactionType: "IN",
		}
		if err := s.stockRepo.CreateTransaction(ctx, tx, txn); err != nil {
			return err
		}

		return s.productRepo.UpdateTotalStock(ctx, tx, productID)
	})
}

func (s *stockService) RemoveStock(
	ctx context.Context,
	productID, subVariantID string,
	qty decimal.Decimal,
) error {

	if qty.LessThanOrEqual(decimal.Zero) {
		return errors.New("quantity must be greater than zero")
	}

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		sub, err := s.subRepo.FindByIDForUpdate(ctx, tx, subVariantID)
		if err != nil {
			return err
		}

		if sub.Stock.LessThan(qty) {
			return errors.New("insufficient stock")
		}

		sub.Stock = sub.Stock.Sub(qty)

		if err := s.subRepo.UpdateStock(ctx, tx, sub); err != nil {
			return err
		}

		txn := &model.StockTransaction{
			ID:              uuid.New(),
			ProductID:       sub.ProductID,
			SubVariantID:    sub.ID,
			Quantity:        qty.Neg(),
			TransactionType: "OUT",
		}

		if err := s.stockRepo.CreateTransaction(ctx, tx, txn); err != nil {
			return err
		}

		return s.productRepo.UpdateTotalStock(ctx, tx, productID)
	})
}
