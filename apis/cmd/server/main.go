package main

import (
	"log"
	"net/http"

	"github.com/Guilherme-Joviniano/go.expert/apis/configs"
	"github.com/Guilherme-Joviniano/go.expert/apis/internal/entity"
	database "github.com/Guilherme-Joviniano/go.expert/apis/internal/infra/database/sqlite"
	"github.com/Guilherme-Joviniano/go.expert/apis/internal/infra/webserver/handlers"
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
	
	http.HandleFunc("/products", productHandler.CreateProduct)
	http.ListenAndServe(":8000", nil)
}
