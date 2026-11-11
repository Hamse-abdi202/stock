
package models

import "time"

type StockMovement struct {
	ID        uint      `json:"id"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Type      string    `json:"type"`
	Quantity  int       `json:"quantity"`
	Note      string    `json:"note"`
	CreatedBy uint      `json:"created_by" `
	CreatedAt time.Time `json:"created_at"`
	Product Product `json:"product,omitempty" gorm:"foreignKey:ProductID"`
	User    User    `json:"user,omitempty" gorm:"foreignKey:CreatedBy"`
}
