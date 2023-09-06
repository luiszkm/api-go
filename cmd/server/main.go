package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/luiszkm/api/configs"
	"github.com/luiszkm/api/internal/Domain/entity"
	"github.com/luiszkm/api/internal/infra/database"
	"github.com/luiszkm/api/internal/infra/http/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
			panic(err)
	}
	envPath := filepath.Join(rootDir, ".env")
	log.Println(envPath)
	configs, err := configs.LoadConfig(envPath)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open((sqlite.Open("test.db")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDb := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDb, configs.TokenAuth, configs.JwtExperesIn )

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/products", productHandler.CreateProduct)
	router.Get("/products/{id}", productHandler.GetProduct)
	router.Get("/products", productHandler.GetAllProducts)
	router.Put("/products/{id}", productHandler.UpdateProduct)
	router.Delete("/products/{id}", productHandler.DeleteProduct)

	router.Post("/users", userHandler.CreateUser)
	router.Post("/users/login", userHandler.GetJWT)
	
	http.ListenAndServe(":8080", router)
}
