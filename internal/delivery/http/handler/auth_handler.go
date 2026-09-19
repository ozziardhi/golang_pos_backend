package handler

import (
	"net/http"
	"pos-backend/internal/usecase"

	"github.com/labstack/echo/v5"
)

// AuthHandler menangani komunikasi HTTP terkait autentikasi (menerima request dan mengirim respons).
type AuthHandler struct {
	authUsecase usecase.AuthUsecase
}

// NewAuthHandler adalah constructor untuk mendaftarkan endpoint rute autentikasi ke router Echo.
func NewAuthHandler(e *echo.Echo, authUsecase usecase.AuthUsecase) {
	handler := &AuthHandler{authUsecase: authUsecase}

	// Mendaftarkan endpoint POST /register dan POST /login.
	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)
}

// RegisterRequest merepresentasikan struktur data JSON yang dikirim oleh klien saat melakukan registrasi.
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register menangani HTTP request untuk pendaftaran akun baru.
func (h *AuthHandler) Register(c *echo.Context) error {
	var req RegisterRequest

	// Membaca dan memetakan payload JSON dari body request ke struct RegisterRequest.
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
	}

	// Memvalidasi apakah semua kolom wajib sudah terisi dengan benar.
	if req.Name == "" || req.Email == "" || req.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "all fields are required"})
	}

	// Menjalankan logika pendaftaran melalui usecase.
	err := h.authUsecase.Register(req.Name, req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Mengembalikan respons sukses dengan status kode 201 Created.
	return c.JSON(http.StatusCreated, map[string]string{"message": "user registered successfully"})
}

// LoginRequest merepresentasikan struktur data JSON yang dikirim klien saat proses login.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Login menangani HTTP request untuk masuk (autentikasi) dan pengembalian token JWT.
func (h *AuthHandler) Login(c *echo.Context) error {
	var req LoginRequest

	// Membaca dan memetakan payload JSON dari body request ke struct LoginRequest.
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request payload"})
	}

	// Menjalankan proses autentikasi via usecase untuk mendapatkan token dan data user.
	token, user, err := h.authUsecase.Login(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	// Mengembalikan respons sukses berupa token JWT dan informasi ringkas pengguna.
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login successful",
		"token":   token,
		"user": map[string]interface{}{
			"id":    user.ID,
			"name":  user.Name,
			"email": user.Email,
			"role":  user.Role,
		},
	})
}
