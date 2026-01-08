package service

import (
	"context"

	"github.com/fathimasithara01/product-inventory-system/internal/models"
	"github.com/fathimasithara01/product-inventory-system/internal/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductService interface {
	CreateProduct(
		ctx context.Context,
		product *models.Product,
		variants []models.Variant,
		options []models.VariantOption,
		subVariants []models.SubVariant,
	) error

	ListProducts(ctx context.Context, page, limit int) ([]models.Product, int64, error)
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

func (s *productService) CreateProduct(
	ctx context.Context,
	product *models.Product,
	variants []models.Variant,
	options []models.VariantOption,
	subVariants []models.SubVariant,
) error {

	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		product.ID = uuid.New()
		if err := s.productRepo.Create(ctx, tx, product); err != nil {
			return err
		}

		if len(variants) > 0 {
			for i := range variants {
				variants[i].ID = uuid.New()
				variants[i].ProductID = product.ID

				if err := s.variantRepo.Create(ctx, tx, &variants[i]); err != nil {
					return err
				}
			}
		}

		// ===== Create Variant Options =====
		if len(options) > 0 {
			for i := range options {
				options[i].ID = uuid.New()
				options[i].VariantID = options[i].VariantID // keep VariantID from request
			}

			if err := s.variantRepo.CreateOptions(ctx, tx, options); err != nil {
				return err
			}
		}

		// ===== Create SubVariants =====
		if len(subVariants) > 0 {
			for i := range subVariants {
				subVariants[i].ID = uuid.New()
				subVariants[i].ProductID = product.ID
			}

			if err := s.subRepo.Create(ctx, tx, subVariants); err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *productService) ListProducts(
	ctx context.Context,
	page, limit int,
) ([]models.Product, int64, error) {

	offset := (page - 1) * limit
	return s.productRepo.List(ctx, offset, limit)
}
