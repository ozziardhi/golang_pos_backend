package repository

import (
	"pos-backend/internal/model"
	"time"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll() ([]model.Product, error)
	Create(product *model.Product) error
	Update(product *model.Product) error
	Delete(id uint) error
	FindByID(id uint) (*model.Product, error)
	DeleteBatch(ids []uint) error
	UpdateStatusBatch(ids []uint, isActive bool) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// Mengambil seluruh data produk dari database
func (r *productRepository) FindAll() ([]model.Product, error) {
	var products []model.Product
	err := r.db.Find(&products).Error
	return products, err
}

// Mencari produk berdasarkan ID unik
func (r *productRepository) FindByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Menyimpan produk baru ke database
func (r *productRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

// Memperbarui data produk yang sudah ada
func (r *productRepository) Update(product *model.Product) error {
	product.UpdatedAt = time.Now()

	// Menggunakan Select untuk memastikan kolom-kolom ini disimpan dengan benar,
	// atau gunakan Updates agar GORM menangani auto-update time-nya.
	return r.db.Save(product).Error
}

// Menghapus beberapa produk berdasarkan ID
func (r *productRepository) DeleteBatch(ids []uint) error {
	return r.db.Where("id IN ?", ids).Delete(&model.Product{}).Error
}

// Menghapus produk berdasarkan ID
func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}

// Memperbarui status aktif/nonaktif beberapa produk berdasarkan ID
func (r *productRepository) UpdateStatusBatch(ids []uint, isActive bool) error {
	return r.db.Model(&model.Product{}).Where("id IN ?", ids).Update("is_active", isActive).Error
}
