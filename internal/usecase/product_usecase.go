package usecase

import (
	"errors"
	"pos-backend/internal/model"
	"pos-backend/internal/repository"
	"time"
)

type ProductUsecase interface {
	GetAllProducts() ([]model.Product, error)
	GetProductByID(id uint) (*model.Product, error)
	CreateProduct(product *model.Product) error
	UpdateProduct(id uint, product *model.Product) error
	DeleteProduct(id uint) error
	DeleteBatchProducts(ids []uint) error
	UpdateStatusBatchProducts(ids []uint, isActive bool) error
}

type productUsecase struct {
	productRepo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{productRepo: repo}
}

func (u *productUsecase) GetAllProducts() ([]model.Product, error) {
	return u.productRepo.FindAll()
}

func (u *productUsecase) GetProductByID(id uint) (*model.Product, error) {
	product, err := u.productRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("produk tidak ditemukan")
	}
	return product, nil
}

func (u *productUsecase) CreateProduct(product *model.Product) error {
	if product.Name == "" || product.Price <= 0 {
		return errors.New("nama produk dan harga jual wajib diisi dengan benar")
	}

	// Set waktu pembuatan dan pembaruan secara eksplisit saat produk baru dibuat
	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now
	product.IsActive = true // Pastikan default aktif

	return u.productRepo.Create(product)
}

func (u *productUsecase) UpdateProduct(id uint, productData *model.Product) error {
	existing, err := u.productRepo.FindByID(id)
	if err != nil {
		return errors.New("produk tidak ditemukan")
	}

	// Memperbarui nilai field sesuai form UI termasuk IsActive
	existing.Name = productData.Name
	existing.SKU = productData.SKU
	existing.Category = productData.Category
	existing.Supplier = productData.Supplier
	existing.Description = productData.Description
	existing.Price = productData.Price
	existing.CostPrice = productData.CostPrice
	existing.Stock = productData.Stock
	existing.TaxEnabled = productData.TaxEnabled
	existing.ImageURL = productData.ImageURL
	existing.IsActive = productData.IsActive // <-- Tambahkan ini agar status toggle tersimpan!

	// Pastikan waktu updated_at diperbarui ke waktu sekarang
	existing.UpdatedAt = time.Now()

	return u.productRepo.Update(existing)
}

func (u *productUsecase) DeleteProduct(id uint) error {
	_, err := u.productRepo.FindByID(id)
	if err != nil {
		return errors.New("produk tidak ditemukan")
	}
	return u.productRepo.Delete(id)
}

func (u *productUsecase) DeleteBatchProducts(ids []uint) error {
	if len(ids) == 0 {
		return errors.New("tidak ada produk yang dipilih")
	}
	return u.productRepo.DeleteBatch(ids)
}

func (u *productUsecase) UpdateStatusBatchProducts(ids []uint, isActive bool) error {
	if len(ids) == 0 {
		return errors.New("tidak ada produk yang dipilih")
	}
	return u.productRepo.UpdateStatusBatch(ids, isActive)
}
