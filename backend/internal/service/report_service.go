package service

import (
	"context"
	"time"

	"github.com/fathimasithara01/product-inventory-system/internal/model"
	"github.com/fathimasithara01/product-inventory-system/internal/repository"
)

type ReportService interface {
	StockReport(ctx context.Context, from, to time.Time) ([]model.StockTransaction, error)
}

type reportService struct {
	stockRepo repository.StockRepository
}

func NewReportService(stockRepo repository.StockRepository) ReportService {
	return &reportService{stockRepo: stockRepo}
}

// StockReport fetches all stock transactions between two dates
func (s *reportService) StockReport(
	ctx context.Context,
	from, to time.Time,
) ([]model.StockTransaction, error) {
	if from.IsZero() || to.IsZero() {
		return nil, nil // Or return error if required
	}

	return s.stockRepo.GetReport(ctx, from, to)
}
