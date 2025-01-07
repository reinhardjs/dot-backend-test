package http

import (
	"github.com/gin-gonic/gin"
	"github.com/reinhardjs/dot-backend-test/internal/delivery/http/handler"
	"github.com/reinhardjs/dot-backend-test/internal/usecase"
	"github.com/reinhardjs/dot-backend-test/pkg/errors"
)

func NewRouter(productUsecase usecase.ProductUsecase, categoryUsecase usecase.CategoryUsecase) *gin.Engine {
	router := gin.Default()

	router.Use(errors.ErrorHandler())

	productHandler := handler.NewProductHandler(productUsecase)
	categoryHandler := handler.NewCategoryHandler(categoryUsecase)

	v1 := router.Group("/api/v1")
	{
		products := v1.Group("/products")
		{
			products.POST("", productHandler.CreateProduct)
			products.GET("", productHandler.GetAllProducts)
			products.GET("/:id", productHandler.GetProduct)
			products.PUT("/:id", productHandler.UpdateProduct)
			products.PATCH("/:id", productHandler.UpdateProduct)
			products.DELETE("/:id", productHandler.DeleteProduct)
		}

		categories := v1.Group("/categories")
		{
			categories.POST("", categoryHandler.CreateCategory)
			categories.GET("", categoryHandler.GetAllCategories)
			categories.GET("/:id", categoryHandler.GetCategory)
			categories.PUT("/:id", categoryHandler.UpdateCategory)
			categories.PATCH("/:id", categoryHandler.UpdateCategory)
			categories.DELETE("/:id", categoryHandler.DeleteCategory)
		}
	}

	return router
}
