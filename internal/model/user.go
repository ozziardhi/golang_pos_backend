package model

import (
	"time"

	gorm "gorm.io/gorm"
)

// User merepresentasikan struktur entitas data (tabel) "users" di dalam database.
type User struct {
	// ID bertindak sebagai kolom Primary Key yang bertambah secara otomatis (auto-increment).
	ID uint `gorm:"primaryKey" json:"id"`

	// CreatedAt mencatat waktu secara otomatis kapan record data user pertama kali dibuat.
	CreatedAt time.Time `json:"created_at"`

	// UpdatedAt mencatat waktu secara otomatis setiap kali ada pembaruan pada data user.
	UpdatedAt time.Time `json:"updated_at"`

	// DeletedAt mendukung fitur Soft Delete (GORM), di mana data tidak dihapus permanen melainkan dicatat waktu penghapusannya.
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`

	// Name menyimpan nama lengkap pengguna dengan tipe data varchar(255) dan tidak boleh kosong (not null).
	Name string `gorm:"type:varchar(255);not null" json:"name"`

	// Email menyimpan alamat email pengguna yang bersifat unik (tidak boleh ada yang sama) dan tidak boleh kosong.
	Email string `gorm:"type:varchar(255);unique;not null" json:"email"`

	// Password menyimpan hasil hash dari kata sandi. Tag json:"-" memastikan password tidak akan pernah ikut terekspos saat data dikonversi ke JSON.
	Password string `gorm:"type:varchar(255);not null" json:"-"`

	// Role menyimpan hak akses pengguna (misalnya 'admin' atau 'cashier') dengan nilai bawaan (default) 'admin'.
	Role string `gorm:"type:varchar(50);not null;default:'admin'" json:"role"`
}
