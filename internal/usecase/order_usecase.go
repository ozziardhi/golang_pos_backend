package usecase

import (
	"pos-backend/internal/model"
	"pos-backend/internal/repository"
)

type OrderUsecase interface {
	GetAllOrders() ([]model.Order, error)
	CreateOrder(order *model.Order) error
}

type orderUsecase struct {
	orderRepo repository.OrderRepository
}

func NewOrderUsecase(orderRepo repository.OrderRepository) OrderUsecase {
	return &orderUsecase{orderRepo: orderRepo}
}

func (uc *orderUsecase) GetAllOrders() ([]model.Order, error) {
	return uc.orderRepo.FindAll()
}

func (uc *orderUsecase) CreateOrder(order *model.Order) error {
	return uc.orderRepo.Create(order)
}
