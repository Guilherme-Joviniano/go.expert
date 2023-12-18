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
	"github.com/go-chi/jwtauth"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configs, err := configs.LoadConfig(".")

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

	userService := database.NewUserRepository(db)

	userHandler := handlers.NewUserHandler(userService, configs.TokenAuth, configs.JWTExpiresIn)

	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/products",
		func(router chi.Router) {
			router.Use(jwtauth.Verifier(configs.TokenAuth)) 
			router.Use(jwtauth.Authenticator)
			 
			router.Post("/", productHandler.CreateProduct)
			router.Get("/", productHandler.ListProducts)
			router.Get("/{id}", productHandler.GetProduct)
			router.Put("/{id}", productHandler.UpdateProduct)
			router.Delete("/{id}", productHandler.DeleteProduct)
		})

	router.Route("/users",
		func(router chi.Router) {
			router.Post("/users", userHandler.CreateUser)
			router.Post("/users/token", userHandler.GetToken)
		})

	http.ListenAndServe(":8000", router)
}
