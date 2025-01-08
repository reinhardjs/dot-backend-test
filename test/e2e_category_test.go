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

func TestE2ECategories(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockCategoryUsecase := &mockCategoryUsecase{}
	mockProductUsecase := &mockProductUsecase{}

	router := delivery_http.NewRouter(mockProductUsecase, mockCategoryUsecase)

	t.Run("Create Category", func(t *testing.T) {
		category := entity.Category{Name: "Test Category"}
		body, _ := json.Marshal(category)

		mockCategoryUsecase.On("CreateCategory", mock.AnythingOfType("*entity.Category")).Return(nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)

		var response entity.Category
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, category.Name, response.Name)
	})

	t.Run("Get Category", func(t *testing.T) {
		mockCategoryUsecase.On("GetCategoryByID", uint(1)).Return(&entity.Category{ID: 1, Name: "Test Category"}, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/categories/1", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response entity.Category
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "Test Category", response.Name)
	})

	t.Run("Update Category", func(t *testing.T) {
		category := entity.Category{ID: 1, Name: "Updated Category"}
		body, _ := json.Marshal(category)

		mockCategoryUsecase.On("UpdateCategory", mock.AnythingOfType("*entity.Category")).Return(nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/api/v1/categories/1", bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response entity.Category
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, category.Name, response.Name)
	})

	t.Run("Delete Category", func(t *testing.T) {
		mockCategoryUsecase.On("DeleteCategory", uint(1)).Return(nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/api/v1/categories/1", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"message": "Category deleted successfully"}`, w.Body.String())
	})

	t.Run("Get All Categories", func(t *testing.T) {
		mockCategoryUsecase.On("GetAllCategories").Return([]entity.Category{
			{ID: 1, Name: "Test Category 1"},
			{ID: 2, Name: "Test Category 2"},
		}, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/categories", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var response []entity.Category
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, 2, len(response))
	})

	t.Run("Create Category With Invalid Data", func(t *testing.T) {
		invalidCategory := map[string]interface{}{
			"name": "",
		}
		body, _ := json.Marshal(invalidCategory)

		// Set up mock expectation for invalid data
		mockCategoryUsecase.On("CreateCategory", mock.AnythingOfType("*entity.Category")).Return(fmt.Errorf("invalid category data"))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/categories", bytes.NewBuffer(body))
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Get Non-Existent Category", func(t *testing.T) {
		// Set up mock to return nil and error for non-existent category
		mockCategoryUsecase.On("GetCategoryByID", uint(999)).Return((*entity.Category)(nil), fmt.Errorf("category not found"))

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/categories/999", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}

type mockCategoryUsecase struct {
	mock.Mock
}

func (m *mockCategoryUsecase) CreateCategory(category *entity.Category) error {
	m.Called(category)
	category.ID = 1
	return nil
}

func (m *mockCategoryUsecase) GetCategoryByID(id uint) (*entity.Category, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Category), args.Error(1)
}

func (m *mockCategoryUsecase) UpdateCategory(category *entity.Category) error {
	m.Called(category)
	return nil
}

func (m *mockCategoryUsecase) DeleteCategory(id uint) error {
	m.Called(id)
	return nil
}

func (m *mockCategoryUsecase) GetAllCategories() ([]entity.Category, error) {
	args := m.Called()
	return args.Get(0).([]entity.Category), args.Error(1)
}
