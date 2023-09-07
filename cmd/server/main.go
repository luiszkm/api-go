package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/luiszkm/api/configs"
	"github.com/luiszkm/api/internal/Domain/entity"
	"github.com/luiszkm/api/internal/infra/database"
	"github.com/luiszkm/api/internal/infra/http/handlers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/luiszkm/api/docs"
)

// @title           Go Expert API Example
// @version         1.0
// @description     Product API with auhtentication
// @termsOfService  http://swagger.io/terms/

// @contact.name   Luis Soares


// @host      localhost:8080
// @BasePath  /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	rootDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	envPath := filepath.Join(rootDir, ".env")
	configs, err := configs.LoadConfig(envPath)
	if err != nil {
		panic(err)
	}
	dsn := configs.DBUser + ":" + configs.DBPassword + "@tcp(" + configs.DBHost + ":" + configs.DBPort + ")/" + configs.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open((mysql.Open(dsn)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})
	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDb := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDb)

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.WithValue("jwt", configs.TokenAuth))
	router.Use(middleware.WithValue("JwtExpiriesIn", configs.JwtExperesIn))


	router.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)

		r.Post("/", productHandler.CreateProduct)
		r.Get("/", productHandler.GetAllProducts)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
	})

	router.Post("/users", userHandler.CreateUser)
	router.Post("/users/login", userHandler.GetJWT)
	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))
	http.ListenAndServe(":8080", router)
}
