package main

import (
	"log"

	"github.com/reinhardjs/dot-backend-test/config"
	"github.com/reinhardjs/dot-backend-test/internal/delivery/http"
	"github.com/reinhardjs/dot-backend-test/internal/infrastructure/cache"
	"github.com/reinhardjs/dot-backend-test/internal/infrastructure/database"
	"github.com/reinhardjs/dot-backend-test/internal/usecase"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgresDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	cache, err := cache.NewRedisClient(cfg.RedisURL)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %v", err)
	}
	redisClient := cache.Client

	productUsecase := usecase.NewProductUsecase(db, cache)
	categoryUsecase := usecase.NewCategoryUsecase(db, redisClient)

	router := http.NewRouter(productUsecase, categoryUsecase)

	log.Printf("Server starting on %s", cfg.ServerAddress)
	if err := router.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
