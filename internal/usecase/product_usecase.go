package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/reinhardjs/dot-backend-test/internal/domain/entity"
	"github.com/reinhardjs/dot-backend-test/internal/domain/repository"
	"github.com/reinhardjs/dot-backend-test/internal/infrastructure/cache"
	"gorm.io/gorm"
)

type ProductUsecase interface {
	CreateProduct(product *entity.Product) error
	GetProductByID(id uint) (*entity.Product, error)
	UpdateProduct(product *entity.Product) error
	DeleteProduct(id uint) error
	GetAllProducts() ([]entity.Product, error)
}

type productUsecase struct {
	repo  repository.ProductRepository
	cache *cache.RedisClient
	db    *gorm.DB
}

func NewProductUsecase(db *gorm.DB, cache *cache.RedisClient) ProductUsecase {
	return &productUsecase{
		repo:  repository.NewProductRepository(db),
		cache: cache,
		db:    db,
	}
}

func (u *productUsecase) CreateProduct(product *entity.Product) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := u.repo.Create(product); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	u.invalidateCache()
	return nil
}

func (u *productUsecase) GetProductByID(id uint) (*entity.Product, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("product:%d", id)

	cachedProduct, err := u.cache.Client.Get(ctx, cacheKey).Result()
	if err == nil {
		var product entity.Product
		if err := json.Unmarshal([]byte(cachedProduct), &product); err == nil {
			return &product, nil
		}
	}

	product, err := u.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	productJSON, _ := json.Marshal(product)
	u.cache.Set(ctx, cacheKey, productJSON, time.Minute*5)

	return product, nil
}

func (u *productUsecase) UpdateProduct(product *entity.Product) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := u.repo.Update(product); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	u.invalidateCache()
	return nil
}

func (u *productUsecase) DeleteProduct(id uint) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := u.repo.Delete(id); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	u.invalidateCache()
	return nil
}

func (u *productUsecase) GetAllProducts() ([]entity.Product, error) {
	return u.repo.FindAll()
}

func (u *productUsecase) invalidateCache() {
	ctx := context.Background()
	u.cache.FlushDB(ctx)
}
