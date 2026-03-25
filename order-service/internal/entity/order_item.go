package entity

import "time"

type OrderItem struct {
	ID           uint    `gorm:"primaryKey"`
	OrderID      uint    `gorm:"not null"`
	ProductID    uint    `gorm:"not null"`
	ProductName  string  `gorm:"not null"`
	ProductPrice float64 `gorm:"type:decimal(10,2);not null"`
	Qty          uint    `gorm:"not null"`
	Subtotal     float64 `gorm:"type:decimal(10,2);not null"`
	CreatedAt    time.Time
}
