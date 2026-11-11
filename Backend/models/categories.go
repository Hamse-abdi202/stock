package models

import "time"

type Category struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name" gorm:"uniqueIndex;not null"`
	Products  []*Product `json:"products,omitempty" gorm:"foreignKey:CategoryID"`
	CreatedAt time.Time  `json:"created_at"`
}