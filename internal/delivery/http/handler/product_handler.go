package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/reinhardjs/dot-backend-test/internal/domain/entity"
	"github.com/reinhardjs/dot-backend-test/internal/usecase"
)

type ProductHandler struct {
	usecase usecase.ProductUsecase
}

func NewProductHandler(usecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: usecase}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if product.Name == "" || product.Price < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product data: name cannot be empty and price must be non-negative"})
		return
	}

	if err := h.usecase.CreateProduct(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := h.usecase.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var product entity.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product.ID = uint(id)
	if err := h.usecase.UpdateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.usecase.DeleteProduct(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.usecase.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
