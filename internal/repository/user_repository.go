package repository

import (
	"pos-backend/internal/model"

	"gorm.io/gorm"
)

// UserRepository mendefinisikan kontrak interface untuk operasi database yang berkaitan dengan entitas User.
type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}

// userRepository adalah struct konkret yang mengimplementasikan UserRepository dengan menyimpan instance koneksi GORM DB.
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository adalah fungsi constructor untuk menginisialisasi repository user baru menggunakan koneksi database.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create bertugas untuk menyimpan data user baru ke dalam database PostgreSQL menggunakan GORM.
func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

// FindByEmail bertugas untuk mencari data user di database berdasarkan alamat email yang spesifik.
func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	// Melakukan query pencarian data pertama yang cocok dengan kondisi kolom email.
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
