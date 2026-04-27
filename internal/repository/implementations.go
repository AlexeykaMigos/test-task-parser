package repository

import (
	"github.com/user/product-parser/internal/domain/entities"
	"github.com/user/product-parser/pkg/parser"
)

type ProductRepo struct {
	client *parser.Client
}

func NewProductRepo(client *parser.Client) *ProductRepo {
	return &ProductRepo{client: client}
}

func (r *ProductRepo) GetProducts(categorySlug string) ([]entities.Product, error) {
	products, err := r.client.GetProducts(categorySlug)
	if err != nil {
		return nil, err
	}

	result := make([]entities.Product, len(products))
	for i, p := range products {
		result[i] = entities.Product{
			Name:  p.Name,
			Price: p.Price,
			Unit:  p.Unit,
			URL:   p.URL,
			Brand: p.Brand,
		}
	}

	return result, nil
}

type CategoryRepo struct {
	client *parser.Client
}

func NewCategoryRepo(client *parser.Client) *CategoryRepo {
	return &CategoryRepo{client: client}
}

func (r *CategoryRepo) GetCategories() ([]entities.Category, error) {
	categories, err := r.client.GetCategories()
	if err != nil {
		return nil, err
	}

	result := make([]entities.Category, len(categories))
	for i, c := range categories {
		result[i] = entities.Category{
			ID:   c.ID,
			Name: c.Name,
			Slug: c.Slug,
		}
	}

	return result, nil
}

type StoreRepo struct {
	client *parser.Client
}

func NewStoreRepo(client *parser.Client) *StoreRepo {
	return &StoreRepo{client: client}
}

func (r *StoreRepo) GetStores() ([]entities.Store, error) {
	stores, err := r.client.GetStores()
	if err != nil {
		return nil, err
	}

	result := make([]entities.Store, len(stores))
	for i, s := range stores {
		result[i] = entities.Store{
			ID:      s.ID,
			Name:    s.Name,
			Address: s.Address,
		}
	}

	return result, nil
}