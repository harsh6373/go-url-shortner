package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/harsh6373/go-url-shortner/config"
	"github.com/harsh6373/go-url-shortner/internal/handler"
	"github.com/harsh6373/go-url-shortner/internal/model"
	"github.com/harsh6373/go-url-shortner/internal/repository"
	"github.com/harsh6373/go-url-shortner/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	cfg := config.LoadConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&model.URL{}, &model.Click{})

	urlRepo := repository.NewURLRepository(db)
	urlService := service.NewURLService(urlRepo)
	urlHandler := handler.NewURLHandler(urlService)

	app := fiber.New()

	app.Post("/api/shorten", urlHandler.Shorten)
	app.Get("/:slug", urlHandler.Redirect)

	log.Println("Server running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
