package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	delivery_http "github.com/reinhardjs/dot-backend-test/internal/delivery/http"
	"github.com/reinhardjs/dot-backend-test/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockProductUsecase := &mockProductUsecase{}
	mockCategoryUsecase := &mockCategoryUsecase{}

	router := delivery_http.NewRouter(mockProductUsecase, mockCategoryUsecase)

	product := entity.Product{Name: "Test Product", Price: 9.99}
	body, _ := json.Marshal(product)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response entity.Product
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, product.Name, response.Name)
	assert.Equal(t, product.Price, response.Price)
}

type mockProductUsecase struct{}

func (m *mockProductUsecase) CreateProduct(product *entity.Product) error {
	product.ID = 1
	return nil
}

func (m *mockProductUsecase) GetProductByID(id uint) (*entity.Product, error) {
	return nil, nil
}

func (m *mockProductUsecase) UpdateProduct(product *entity.Product) error {
	return nil
}

func (m *mockProductUsecase) DeleteProduct(id uint) error {
	return nil
}

func (m *mockProductUsecase) GetAllProducts() ([]entity.Product, error) {
	return nil, nil
}

type mockCategoryUsecase struct{}

func (m *mockCategoryUsecase) CreateCategory(category *entity.Category) error {
	return nil
}

func (m *mockCategoryUsecase) GetCategoryByID(id uint) (*entity.Category, error) {
	return nil, nil
}

func (m *mockCategoryUsecase) UpdateCategory(category *entity.Category) error {
	return nil
}

func (m *mockCategoryUsecase) DeleteCategory(id uint) error {
	return nil
}

func (m *mockCategoryUsecase) GetAllCategories() ([]entity.Category, error) {
	return nil, nil
}
