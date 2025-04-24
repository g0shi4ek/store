package main

import (
	"context"
	"log"

	"github.com/g0shi4ek/store/config"
	"github.com/g0shi4ek/store/internal/store/handlers"
	"github.com/g0shi4ek/store/internal/store/repository"
	"github.com/g0shi4ek/store/internal/store/services"
	"github.com/g0shi4ek/store/pkg/db"
)

func main() {
	cfg := config.LoadConfig()

	pool, err := db.NewPool(context.Background(), cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	repo := repository.NewRepository(pool)
	service := services.NewStoreService(repo, cfg)
	handler := handlers.NewStoreHandler(service)

	router := handler.InitRoutes()
	log.Println("Starting server on :8080")
	log.Fatal(router.Run(":8080"))
}
