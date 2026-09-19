package usecase

import (
	"errors"
	"os"
	"time"

	"pos-backend/internal/model"
	"pos-backend/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// AuthUsecase mendefinisikan kontrak interface untuk seluruh logika bisnis autentikasi (registrasi dan login).
type AuthUsecase interface {
	Register(name, email, password string) error
	Login(email, password string) (string, *model.User, error)
}

// authUsecase adalah struct konkret yang mengimplementasikan AuthUsecase dan membutuhkan akses ke UserRepository.
type authUsecase struct {
	userRepo repository.UserRepository
}

// NewAuthUsecase adalah fungsi constructor untuk membuat instance baru dari authUsecase.
func NewAuthUsecase(userRepo repository.UserRepository) AuthUsecase {
	return &authUsecase{userRepo: userRepo}
}

// Register menangani logika bisnis pendaftaran pengguna baru.
func (uc *authUsecase) Register(name, email, password string) error {
	// Memeriksa apakah email sudah terdaftar di database sebelumnya.
	existingUser, _ := uc.userRepo.FindByEmail(email)
	if existingUser != nil {
		return errors.New("email already registered")
	}

	// Mengamankan password dengan melakukan hashing menggunakan bcrypt.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Membentuk objek User baru dengan role default sebagai "cashier".
	user := &model.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPassword),
		Role:     "cashier",
	}

	// Menyimpan data user baru ke database melalui repository.
	return uc.userRepo.Create(user)
}

// Login menangani logika autentikasi masuk pengguna dan pembuatan token JWT.
func (uc *authUsecase) Login(email, password string) (string, *model.User, error) {
	// Mencari data pengguna berdasarkan email.
	user, err := uc.userRepo.FindByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	// Memverifikasi kecocokan password mentah dengan hash password yang ada di database.
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", nil, errors.New("invalid email or password")
	}

	// Membuat klaim (claims) untuk token JWT yang berisi data user dan masa aktif 24 jam.
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	// Membuat objek token baru menggunakan metode enkripsi HS256 dan klaim yang telah dibuat.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Mengambil kunci rahasia (secret key) dari variabel lingkungan (.env).
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		secretKey = "default_secret"
	}

	// Menandatangani token menggunakan secret key untuk menghasilkan string token yang aman.
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", nil, err
	}

	// Mengembalikan token JWT, data user, dan nil (tanda tanpa error).
	return tokenString, user, nil
}
