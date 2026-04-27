package repository

import "github.com/user/product-parser/internal/domain/entities"

type ProductRepository interface {
	GetProducts(categorySlug string) ([]entities.Product, error)
}

type CategoryRepository interface {
	GetCategories() ([]entities.Category, error)
}

type StoreRepository interface {
	GetStores() ([]entities.Store, error)
}