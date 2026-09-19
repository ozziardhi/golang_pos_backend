package model

import "time"

type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`          // Nama Produk *
	SKU         string    `json:"sku"`                           // SKU / Kode Produk
	Category    string    `json:"category"`                      // Kategori / Tipe
	Supplier    string    `json:"supplier"`                      // Vendor / Supplier
	Description string    `json:"description"`                   // Deskripsi Produk
	Price       float64   `json:"price" gorm:"not null"`         // Harga Jual (RP) *
	CostPrice   float64   `json:"cost_price"`                    // Harga Modal (RP)
	Stock       int       `json:"stock"`                         // Stok Awal
	TaxEnabled  bool      `json:"tax_enabled"`                   // Kenakan pajak pada produk ini
	ImageURL    string    `json:"image_url"`                     // Foto Produk
	IsActive    bool      `json:"is_active" gorm:"default:true"` // Status Produk (Aktif / Nonaktif)
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
}
