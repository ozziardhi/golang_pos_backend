package handler

import (
	"net/http"
	"pos-backend/internal/model"
	"pos-backend/internal/usecase"

	"github.com/labstack/echo/v5"
)

type OrderHandler struct {
	orderUC usecase.OrderUsecase
}

func NewOrderHandler(e *echo.Echo, orderUC usecase.OrderUsecase) {
	handler := &OrderHandler{orderUC: orderUC}

	e.GET("/orders", handler.GetOrders)
	e.POST("/orders", handler.CreateOrder)
}

func (h *OrderHandler) GetOrders(c *echo.Context) error {
	orders, err := h.orderUC.GetAllOrders()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Gagal mengambil data pesanan",
		})
	}
	return c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) CreateOrder(c *echo.Context) error {
	var order model.Order
	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request payload",
		})
	}

	if err := h.orderUC.CreateOrder(&order); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Gagal menyimpan transaksi",
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Pesanan berhasil disimpan",
		"data":    order,
	})
}
