package dtos

import "time"

type CreateStockMovementDto struct {
	ProductID uint   `json:"product_id" binding:"required"`
	Type      string `json:"type" binding:"required,oneof=IN OUT"`
	Quantity  int    `json:"quantity" binding:"required,gt=0"`
	Note      string `json:"note"`
}

type StockMovementResponseDto struct {
	ID        uint      `json:"id"`
	ProductID uint      `json:"product_id"`
	Type      string    `json:"type"`
	Quantity  int       `json:"quantity"`
	Note      string    `json:"note"`
	CreatedBy uint      `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
}