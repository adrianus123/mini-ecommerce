package entity

import "time"

type Payment struct {
	ID              uint    `gorm:"primaryKey"`
	PaymentCode     string  `gorm:"not null"`
	OrderID         string  `gorm:"not null"`
	UserID          string  `gorm:"not null"`
	Amount          float64 `gorm:"not null"`
	PaymentMethod   string  `gorm:"not null"`
	PaymentProvider string  `gorm:"not null"`
	Status          string  `gorm:"not null"`
	ExternalID      string
	PaidAt          time.Time
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
