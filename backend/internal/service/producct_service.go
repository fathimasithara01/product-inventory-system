package service

import (
	"context"
	"errors"

	"github.com/fathimasithara01/product-inventory-system/internal/dto"
	"github.com/fathimasithara01/product-inventory-system/internal/model"
	"github.com/fathimasithara01/product-inventory-system/internal/repository"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProductService interface {
	CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*model.Product, error)
	ListProducts(ctx context.Context, page, limit int) ([]model.Product, int64, error)
}

type productService struct {
	db          *gorm.DB
	productRepo repository.ProductRepository
	variantRepo repository.VariantRepository
	subRepo     repository.SubVariantRepository
}

func NewProductService(
	db *gorm.DB,
	productRepo repository.ProductRepository,
	variantRepo repository.VariantRepository,
	subRepo repository.SubVariantRepository,
) ProductService {
	return &productService{
		db:          db,
		productRepo: productRepo,
		variantRepo: variantRepo,
		subRepo:     subRepo,
	}
}

func (s *productService) CreateProduct(ctx context.Context, req *dto.CreateProductRequest) (*model.Product, error) {

	if req.ProductName == "" || req.ProductCode == "" {
		return nil, errors.New("product name and code are required")
	}

	product := &model.Product{
		ProductID:    req.ProductID,
		ProductCode:  req.ProductCode,
		ProductName:  req.ProductName,
		ProductImage: req.ProductImage,
		CreatedUser:  uuid.MustParse(req.CreatedUser),
		HSNCode:      req.HSNCode,
	}

	err := s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// === 1. Create Product ===
		product.ID = uuid.New()
		if err := s.productRepo.Create(ctx, tx, product); err != nil {
			return err
		}

		// === 2. Create Variants ===
		variantUUIDs := make([]uuid.UUID, len(req.Variants))
		for i, v := range req.Variants {
			variant := model.Variant{
				ID:        uuid.New(),
				ProductID: product.ID,
				Name:      v.Name,
			}
			if err := s.variantRepo.Create(ctx, tx, &variant); err != nil {
				return err
			}
			variantUUIDs[i] = variant.ID
		}

		// === 3. Create Options ===
		optionUUIDs := make([]uuid.UUID, len(req.Options))
		for i, o := range req.Options {
			if o.VariantIndex >= len(variantUUIDs) {
				return errors.New("invalid variant index for option")
			}
			option := model.VariantOption{
				ID:        uuid.New(),
				VariantID: variantUUIDs[o.VariantIndex],
				Value:     o.Value,
			}
			if err := s.variantRepo.CreateOptions(ctx, tx, []model.VariantOption{option}); err != nil {
				return err
			}
			optionUUIDs[i] = option.ID
		}

		// === 4. Create SubVariants ===
		for _, sv := range req.SubVariants {
			var optionIDs []string
			for _, idx := range sv.OptionIndices {
				if idx >= len(optionUUIDs) {
					return errors.New("invalid option index for subvariant")
				}
				optionIDs = append(optionIDs, optionUUIDs[idx].String())
			}
			subVariant := model.SubVariant{
				ID:        uuid.New(),
				ProductID: product.ID,
				OptionIDs: optionIDs,
				SKU:       sv.SKU,
				Stock:     decimal.RequireFromString(sv.Stock),
			}
			if err := s.subRepo.Create(ctx, tx, []model.SubVariant{subVariant}); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *productService) ListProducts(ctx context.Context, page, limit int) ([]model.Product, int64, error) {
	offset := (page - 1) * limit
	return s.productRepo.List(ctx, offset, limit)
}
