package main

import (
	"net/http"

	//"github.com/luiszkm/api/configs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/luiszkm/api/internal/Domain/entity"
	"github.com/luiszkm/api/internal/infra/database"
	"github.com/luiszkm/api/internal/infra/http/handlers"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	// _, err := configs.LoadConfig(".")
	// if err != nil {
	// 	panic(err)
	// }
	db, err := gorm.Open((sqlite.Open("test.db")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/products", productHandler.CreateProduct)
	router.Get("/products/{id}", productHandler.GetProduct)
	router.Get("/products", productHandler.GetAllProducts)
	router.Put("/products/{id}", productHandler.UpdateProduct)
	router.Delete("/products/{id}", productHandler.DeleteProduct)
	http.ListenAndServe(":8080", router)
}
