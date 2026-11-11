
package dtos

type CreateProductDto struct {
	Name              string  `json:"name" binding:"required"`
	SKU               string  `json:"sku" binding:"required"`
	CategoryID        uint    `json:"category_id" binding:"required"`
	PurchasePrice     float64 `json:"purchase_price" binding:"required,gte=0"`
	SellingPrice      float64 `json:"selling_price" binding:"required,gte=0"`
	Quantity          int     `json:"quantity" binding:"gte=0"`
	LowStockThreshold int     `json:"low_stock_threshold" binding:"gte=0"`
	Status            string  `json:"status"`
	Description       string  `json:"description"`
}

type UpdateProductDto struct {
	Name              string  `json:"name" binding:"required"`
	SKU               string  `json:"sku" binding:"required"`
	CategoryID        uint    `json:"category_id" binding:"required"`
	PurchasePrice     float64 `json:"purchase_price" binding:"required,gte=0"`
	SellingPrice      float64 `json:"selling_price" binding:"required,gte=0"`
	LowStockThreshold int     `json:"low_stock_threshold" binding:"gte=0"`
	Status            string  `json:"status" binding:"required"`
	Description       string  `json:"description"`
}
