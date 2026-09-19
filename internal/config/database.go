package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB bertugas untuk memuat konfigurasi lingkungan, membangun koneksi ke database PostgreSQL, dan mengembalikan instance GORM DB.
func ConnectDB() *gorm.DB {
	// Memuat variabel lingkungan dari file .env ke dalam sistem runtime Go.
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Menyusun string DSN (Data Source Name) untuk konfigurasi koneksi PostgreSQL berdasarkan variabel lingkungan.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// Membuka koneksi ke database PostgreSQL menggunakan driver GORM.
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Mengembalikan pointer instance *gorm.DB yang siap digunakan oleh lapisan repository.
	return db
}
