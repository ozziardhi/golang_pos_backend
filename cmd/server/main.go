package main

import (
	"log"
	"net/http"
	"os"

	"pos-backend/internal/config"
	"pos-backend/internal/delivery/http/handler"
	"pos-backend/internal/model"
	"pos-backend/internal/repository"
	"pos-backend/internal/usecase"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	// Load environment variables dari file .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables")
	}

	// Inisialisasi Database & Auto Migration (User, Product, & Order)
	db := config.ConnectDB()
	if err := db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{}); err != nil {
		log.Fatalf("Failed to run database migration: %v", err)
	}

	// Inisialisasi Router Echo v5
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://127.0.0.1:5173"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	// Dependency Injection & Pendaftaran Rute untuk Auth
	userRepo := repository.NewUserRepository(db)
	authUC := usecase.NewAuthUsecase(userRepo)
	handler.NewAuthHandler(e, authUC)

	// Dependency Injection & Pendaftaran Rute untuk Product (CRUD)
	productRepo := repository.NewProductRepository(db)
	productUC := usecase.NewProductUsecase(productRepo)
	handler.NewProductHandler(e, productUC)

	// Dependency Injection & Pendaftaran Rute untuk Orders (Pesanan)
	orderRepo := repository.NewOrderRepository(db)
	orderUC := usecase.NewOrderUsecase(orderRepo)
	handler.NewOrderHandler(e, orderUC)

	// Ambil port dari environment atau default ke 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Jalankan server
	log.Printf("Server is running on port %s...", port)
	if err := e.Start(":" + port); err != nil {
		log.Fatalf("Server shutdown abruptly: %v", err)
	}
}
