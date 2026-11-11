package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID                uint            `json:"id"`
	Name              string          `json:"name"`
	SKU               string          `json:"sku" gorm:"uniqueIndex;"`
	CategoryID        uint            `json:"category_id" `
	Category          *Category       `json:"category,omitempty" gorm:"foreignKey:CategoryID"`
	PurchasePrice     float64         `json:"purchase_price" gorm:"not null"`
	SellingPrice      float64         `json:"selling_price" gorm:"not null"`
	Quantity          int             `json:"quantity" gorm:"not null;default:0"`
	LowStockThreshold int             `json:"low_stock_threshold" gorm:"not null;default:5"`
	Status            string          `json:"status" gorm:"not null;default:active"`
	Description       string          `json:"description" gorm:"type:text"`
	StockMovements    []*StockMovement `json:"stock_movements,omitempty" gorm:"foreignKey:ProductID"`
	CreatedAt         time.Time       `json:"created_at"`
	UpdatedAt         time.Time       `json:"updated_at"`
	DeletedAt         gorm.DeletedAt  `json:"-" gorm:"index"`
}