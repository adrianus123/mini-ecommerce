package entity

import "time"

type Order struct {
	ID            uint    `gorm:"primaryKey"`
	OrderCode     string  `gorm:"uniqueIndex"`
	UserID        string  `gorm:"not null"`
	TotalPrice    float64 `gorm:"type:decimal(10,2);not null"`
	Status        string  `gorm:"not null"`
	PaymentMethod string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
