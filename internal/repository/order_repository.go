package repository

import (
	"pos-backend/internal/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAll() ([]model.Order, error)
	Create(order *model.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) FindAll() ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Find(&orders).Error
	return orders, err
}

func (r *orderRepository) Create(order *model.Order) error {
	return r.db.Create(order).Error
}
