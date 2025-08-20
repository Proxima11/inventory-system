package models

import (
	"time"

	"gorm.io/gorm"
)

// Item merepresentasikan barang dalam sistem
type Item struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"type:varchar(100);not null" json:"name"`
	SKU         string         `gorm:"type:varchar(50);unique;not null" json:"sku"`
	Description string         `gorm:"type:text" json:"description,omitempty"`
	Quantity    int            `gorm:"not null" json:"quantity"`
	MinStock    int            `gorm:"default:0" json:"min_stock"` // Batas minimum stok (untuk peringatan)
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}
