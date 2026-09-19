package model

import (
	"time"
)

type OrderItem struct {
	ID        uint    `json:"id" gorm:"primaryKey"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id" json:"product_id"`
	Quantity  int     `json:"quantity" json:"quantity"`
	Price     float64 `json:"price" json:"price"`
}

type Order struct {
	ID                uint        `json:"id" gorm:"primaryKey"`
	OrderReference    string      `json:"order_reference" gorm:"unique;not null"`
	Channel           string      `json:"channel"`
	Date              time.Time   `json:"date"`
	CustomerName      string      `json:"customer_name"`
	PaymentStatus     string      `json:"payment_status"`
	FulfillmentStatus string      `json:"fulfillment_status"`
	Total             float64     `json:"total"`
	Items             []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
}
