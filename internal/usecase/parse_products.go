package usecase

import (
	"fmt"
	"time"

	"github.com/user/product-parser/internal/domain/entities"
	"github.com/user/product-parser/internal/repository"
)

type ParseProductsUseCase struct {
	productRepo  repository.ProductRepository
	categoryRepo repository.CategoryRepository
}

func NewParseProductsUseCase(
	productRepo repository.ProductRepository,
	categoryRepo repository.CategoryRepository,
) *ParseProductsUseCase {
	return &ParseProductsUseCase{
		productRepo:  productRepo,
		categoryRepo: categoryRepo,
	}
}

func (uc *ParseProductsUseCase) Execute(categorySlug string) (*entities.ParseResult, error) {
	products, err := uc.productRepo.GetProducts(categorySlug)
	if err != nil {
		return nil, fmt.Errorf("failed to get products: %w", err)
	}

	return &entities.ParseResult{
		Category:   categorySlug,
		ParsedAt:   time.Now().Format(time.RFC3339),
		TotalCount: len(products),
		Products:   products,
	}, nil
}

func (uc *ParseProductsUseCase) GetCategories() ([]entities.Category, error) {
	return uc.categoryRepo.GetCategories()
}