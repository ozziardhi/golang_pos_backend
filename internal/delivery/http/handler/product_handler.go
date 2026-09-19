package handler

import (
	"net/http"
	"strconv"

	"pos-backend/internal/model"
	"pos-backend/internal/usecase"

	"github.com/labstack/echo/v5"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(e *echo.Echo, usecase usecase.ProductUsecase) {
	handler := &ProductHandler{productUsecase: usecase}

	// Mendaftarkan rute grup produk
	prodGroup := e.Group("/products")
	prodGroup.GET("", handler.GetProducts)
	prodGroup.GET("/:id", handler.GetProductByID)
	prodGroup.POST("", handler.CreateProduct)
	prodGroup.PUT("/:id", handler.UpdateProduct)
	prodGroup.DELETE("/:id", handler.DeleteProduct)
	prodGroup.POST("/bulk-delete", handler.BulkDeleteProducts)
	prodGroup.PUT("/bulk-status", handler.BulkUpdateStatus)
}

func (h *ProductHandler) GetProductByID(c *echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID produk tidak valid"})
	}

	// Pastikan di Usecase Anda ada method untuk get by ID, atau Anda bisa gunakan repository langsung jika tersedia
	product, err := h.productUsecase.GetProductByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Produk tidak ditemukan"})
	}

	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) GetProducts(c *echo.Context) error {
	products, err := h.productUsecase.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) CreateProduct(c *echo.Context) error {
	var product model.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "format data tidak valid"})
	}

	if err := h.productUsecase.CreateProduct(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) UpdateProduct(c *echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID produk tidak valid"})
	}

	var product model.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "format data tidak valid"})
	}

	if err := h.productUsecase.UpdateProduct(uint(id), &product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) DeleteProduct(c *echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "ID produk tidak valid"})
	}

	if err := h.productUsecase.DeleteProduct(uint(id)); err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "produk berhasil dihapus"})
}

func (h *ProductHandler) BulkDeleteProducts(c *echo.Context) error {
	var req struct {
		IDs []uint `json:"ids"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "format data tidak valid"})
	}

	if err := h.productUsecase.DeleteBatchProducts(req.IDs); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "produk terpilih berhasil dihapus"})
}

// Fungsi Handler untuk Nonaktifkan Massal
func (h *ProductHandler) BulkUpdateStatus(c *echo.Context) error {
	var req struct {
		IDs      []uint `json:"ids"`
		IsActive bool   `json:"is_active"`
	}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "format data tidak valid"})
	}

	if err := h.productUsecase.UpdateStatusBatchProducts(req.IDs, req.IsActive); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "status produk berhasil diperbarui"})
}
