package models

import (
	"time"

	"gorm.io/gorm"
)

type TransactionType string

const (
	StockIn  TransactionType = "IN"
	StockOut TransactionType = "OUT"
)

type Transaction struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	ItemID    uint            `json:"item_id"`
	Item      Item            `gorm:"foreignKey:ItemID" json:"item,omitempty"`
	Type      TransactionType `gorm:"type:varchar(10);not null" json:"type"`
	Quantity  int             `gorm:"not null" json:"quantity"`
	Note      string          `gorm:"type:text" json:"note,omitempty"`
	CreatedAt time.Time       `json:"created_at"`
	DeletedAt gorm.DeletedAt  `gorm:"index" json:"-"`
}
