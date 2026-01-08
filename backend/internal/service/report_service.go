package service

import (
	"context"
	"time"

	"github.com/fathimasithara01/product-inventory-system/internal/models"
	"github.com/fathimasithara01/product-inventory-system/internal/repository"
)

type ReportService interface {
	StockReport(ctx context.Context, from, to time.Time) ([]models.StockTransaction, error)
}

type reportService struct {
	stockRepo repository.StockRepository
}

func NewReportService(stockRepo repository.StockRepository) ReportService {
	return &reportService{stockRepo: stockRepo}
}

func (s *reportService) StockReport(
	ctx context.Context,
	from, to time.Time,
) ([]models.StockTransaction, error) {
	return s.stockRepo.GetReport(ctx, from, to)
}
