package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/reinhardjs/dot-backend-test/internal/domain/entity"
	"github.com/reinhardjs/dot-backend-test/internal/domain/repository"
	"gorm.io/gorm"
)

type CategoryUsecase interface {
	CreateCategory(category *entity.Category) error
	GetCategoryByID(id uint) (*entity.Category, error)
	UpdateCategory(category *entity.Category) error
	DeleteCategory(id uint) error
	GetAllCategories() ([]entity.Category, error)
}

type categoryUsecase struct {
	repo  repository.CategoryRepository
	cache *redis.Client
	db    *gorm.DB
}

func NewCategoryUsecase(db *gorm.DB, cache *redis.Client) CategoryUsecase {
	return &categoryUsecase{
		repo:  repository.NewCategoryRepository(db),
		cache: cache,
		db:    db,
	}
}

func (u *categoryUsecase) CreateCategory(category *entity.Category) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := u.repo.Create(category); err != nil {
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

func (u *categoryUsecase) GetCategoryByID(id uint) (*entity.Category, error) {
	ctx := context.Background()
	cacheKey := fmt.Sprintf("category:%d", id)

	cachedCategory, err := u.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var category entity.Category
		if err := json.Unmarshal([]byte(cachedCategory), &category); err == nil {
			return &category, nil
		}
	}

	category, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	categoryJSON, _ := json.Marshal(category)
	u.cache.Set(ctx, cacheKey, categoryJSON, time.Minute*5)

	return category, nil
}

func (u *categoryUsecase) UpdateCategory(category *entity.Category) error {
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if err := u.repo.Update(category); err != nil {
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

func (u *categoryUsecase) DeleteCategory(id uint) error {
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

func (u *categoryUsecase) GetAllCategories() ([]entity.Category, error) {
	return u.repo.GetAll()
}

func (u *categoryUsecase) invalidateCache() {
	ctx := context.Background()
	u.cache.FlushDB(ctx)
}
