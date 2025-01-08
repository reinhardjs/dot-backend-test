package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/reinhardjs/dot-backend-test/config"
	delivery_http "github.com/reinhardjs/dot-backend-test/internal/delivery/http"
	"github.com/reinhardjs/dot-backend-test/internal/domain/entity"
	"github.com/reinhardjs/dot-backend-test/internal/infrastructure/cache"
	"github.com/reinhardjs/dot-backend-test/internal/infrastructure/database"
	"github.com/reinhardjs/dot-backend-test/internal/usecase"
	"github.com/stretchr/testify/assert"
)

func setupTestEnvironment(t *testing.T) *gin.Engine {
    cfg := config.Load()

    // Connect to test database
    db, err := database.NewPostgresDB(cfg.DatabaseURL)
    assert.NoError(t, err)

    // Connect to test Redis
    cache, err := cache.NewRedisClient(cfg.RedisURL)
    assert.NoError(t, err)

    // Run migrations
    err = db.AutoMigrate(&entity.Category{}, &entity.Product{})
    assert.NoError(t, err)

    // Clean up database
    db.Exec("DELETE FROM products")
    db.Exec("DELETE FROM categories")

    // Initialize usecases with real implementations
    productUsecase := usecase.NewProductUsecase(db, cache)
    categoryUsecase := usecase.NewCategoryUsecase(db, cache.Client)

    // Set gin mode to release for production
    gin.SetMode(gin.ReleaseMode)

    return delivery_http.NewRouter(productUsecase, categoryUsecase)
}

func TestProductE2E(t *testing.T) {
    router := setupTestEnvironment(t)

    t.Run("Product CRUD Operations", func(t *testing.T) {
        // Create Category first to avoid foreign key constraint violation
        category := entity.Category{Name: "Test Category"}
        body, _ := json.Marshal(category)
        w := httptest.NewRecorder()
        req, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewBuffer(body))
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusCreated, w.Code)

        var createdCategory entity.Category
        json.Unmarshal(w.Body.Bytes(), &createdCategory)
        assert.NotZero(t, createdCategory.ID)

        // Create Product - Success
        product := entity.Product{Name: "Test Product", Price: 9.99, CategoryID: createdCategory.ID}
        body, _ = json.Marshal(product)
        w = httptest.NewRecorder()
        req, _ = http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(body))
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusCreated, w.Code)

        var createdProduct entity.Product
        json.Unmarshal(w.Body.Bytes(), &createdProduct)
        assert.NotZero(t, createdProduct.ID)

        // Create Product - Invalid Data
        invalidProduct := map[string]interface{}{"name": "", "price": -1}
        body, _ = json.Marshal(invalidProduct)
        w = httptest.NewRecorder()
        req, _ = http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(body))
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusBadRequest, w.Code)

        // Get Product - Success
        w = httptest.NewRecorder()
        req, _ = http.NewRequest("GET", fmt.Sprintf("/api/v1/products/%d", createdProduct.ID), nil)
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusOK, w.Code)

        // Get Product - Not Found
        w = httptest.NewRecorder()
        req, _ = http.NewRequest("GET", "/api/v1/products/999", nil)
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusNotFound, w.Code)

        // Update Product - Success
        updatedProduct := entity.Product{Name: "Updated Product", Price: 19.99, CategoryID: createdCategory.ID}
        body, _ = json.Marshal(updatedProduct)
        w = httptest.NewRecorder()
        req, _ = http.NewRequest("PUT", fmt.Sprintf("/api/v1/products/%d", createdProduct.ID), bytes.NewBuffer(body))
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusOK, w.Code)

        // Update Product - Not Found
        w = httptest.NewRecorder()
        req, _ = http.NewRequest("PUT", "/api/v1/products/999", bytes.NewBuffer(body))
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusNotFound, w.Code)

        // Delete Product - Success
        w = httptest.NewRecorder()
        req, _ = http.NewRequest("DELETE", fmt.Sprintf("/api/v1/products/%d", createdProduct.ID), nil)
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusOK, w.Code)

        // Delete Product - Not Found
        w = httptest.NewRecorder()
        req, _ = http.NewRequest("DELETE", "/api/v1/products/999", nil)
        router.ServeHTTP(w, req)
        assert.Equal(t, http.StatusNotFound, w.Code)
    })
}

