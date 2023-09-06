package main

import (
	"net/http"

	//"github.com/luiszkm/api/configs"
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

	http.HandleFunc("/products", productHandler.CreateProduct )
	http.ListenAndServe(":8080", nil)
}

