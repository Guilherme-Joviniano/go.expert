package main

import (
	"log"
	"net/http"

	"github.com/Guilherme-Joviniano/go.expert/apis/configs"
	"github.com/Guilherme-Joviniano/go.expert/apis/internal/entity"
	database "github.com/Guilherme-Joviniano/go.expert/apis/internal/infra/database/sqlite"
	"github.com/Guilherme-Joviniano/go.expert/apis/internal/infra/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("[FATAL DB ERROR]: ", err)
		panic(err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productService := database.NewProductRepository(db)
	productHandler := handlers.NewProductHandler(productService)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/products", productHandler.CreateProduct)
	router.Get("/products/{id}", productHandler.GetProduct)
	router.Put("/products/{id}", productHandler.UpdateProduct)
	router.Delete("/products/{id}", productHandler.DeleteProduct)
	

	http.ListenAndServe(":8000", router)
}
