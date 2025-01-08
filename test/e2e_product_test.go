package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	delivery_http "github.com/reinhardjs/dot-backend-test/internal/delivery/http"
	"github.com/reinhardjs/dot-backend-test/internal/domain/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockProductUsecase := &mockProductUsecase{}
	mockCategoryUsecase := &mockCategoryUsecase{}

	router := delivery_http.NewRouter(mockProductUsecase, mockCategoryUsecase)

	product := entity.Product{Name: "Test Product", Price: 9.99}
	body, _ := json.Marshal(product)

	mockProductUsecase.On("CreateProduct", mock.AnythingOfType("*entity.Product")).Return(nil)

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

func TestGetProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockProductUsecase := &mockProductUsecase{}
	mockCategoryUsecase := &mockCategoryUsecase{}

	router := delivery_http.NewRouter(mockProductUsecase, mockCategoryUsecase)

	mockProductUsecase.On("GetProductByID", uint(1)).Return(&entity.Product{ID: 1, Name: "Test Product", Price: 9.99}, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response entity.Product
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Test Product", response.Name)
	assert.Equal(t, 9.99, response.Price)
}

func TestUpdateProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockProductUsecase := &mockProductUsecase{}
	mockCategoryUsecase := &mockCategoryUsecase{}

	router := delivery_http.NewRouter(mockProductUsecase, mockCategoryUsecase)

	product := entity.Product{ID: 1, Name: "Updated Product", Price: 19.99}
	body, _ := json.Marshal(product)

	mockProductUsecase.On("UpdateProduct", mock.AnythingOfType("*entity.Product")).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/products/1", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response entity.Product
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, product.Name, response.Name)
	assert.Equal(t, product.Price, response.Price)
}

func TestDeleteProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockProductUsecase := &mockProductUsecase{}
	mockCategoryUsecase := &mockCategoryUsecase{}

	router := delivery_http.NewRouter(mockProductUsecase, mockCategoryUsecase)

	mockProductUsecase.On("DeleteProduct", uint(1)).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/products/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "Product deleted successfully"}`, w.Body.String())
}

func TestGetAllProducts(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockProductUsecase := &mockProductUsecase{}
	mockCategoryUsecase := &mockCategoryUsecase{}

	router := delivery_http.NewRouter(mockProductUsecase, mockCategoryUsecase)

	mockProductUsecase.On("GetAllProducts").Return([]entity.Product{
		{ID: 1, Name: "Test Product 1", Price: 9.99},
		{ID: 2, Name: "Test Product 2", Price: 19.99},
	}, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response []entity.Product
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, 2, len(response))
}

func TestCreateProductWithInvalidData(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockProductUsecase := &mockProductUsecase{}
	mockCategoryUsecase := &mockCategoryUsecase{}

	router := delivery_http.NewRouter(mockProductUsecase, mockCategoryUsecase)

	invalidProduct := map[string]interface{}{
			"name":  "",
			"price": -1,
	}
	body, _ := json.Marshal(invalidProduct)

	// Set up mock expectation for invalid data
	mockProductUsecase.On("CreateProduct", mock.AnythingOfType("*entity.Product")).Return(fmt.Errorf("invalid product data"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/products", bytes.NewBuffer(body))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetNonExistentProduct(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockProductUsecase := &mockProductUsecase{}
	mockCategoryUsecase := &mockCategoryUsecase{}

	router := delivery_http.NewRouter(mockProductUsecase, mockCategoryUsecase)

	// Return nil, error instead of trying to cast nil to *entity.Product
	mockProductUsecase.On("GetProductByID", uint(999)).Return((*entity.Product)(nil), fmt.Errorf("product not found"))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/products/999", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Product not found")
}

type mockProductUsecase struct {
	mock.Mock
}

func (m *mockProductUsecase) CreateProduct(product *entity.Product) error {
	args := m.Called(product)
	product.ID = 1
	return args.Error(0)
}

func (m *mockProductUsecase) GetProductByID(id uint) (*entity.Product, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Product), args.Error(1)
}

func (m *mockProductUsecase) UpdateProduct(product *entity.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *mockProductUsecase) DeleteProduct(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *mockProductUsecase) GetAllProducts() ([]entity.Product, error) {
	args := m.Called()
	return args.Get(0).([]entity.Product), args.Error(1)
}
