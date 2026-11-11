package models

import "time"

type Role string

const (
	RoleAdmin          Role = "ADMIN"
	RoleStudentAffairs Role = "Employee"
)

type User struct {
	ID           uint      `json:"id"`
	FullName     string    `json:"fullName"`
	EmailAddress string    `json:"emailAddress"`
	Password     string    `json:"-"`
	LastLoginAt  time.Time `json:"lastLoginAt"`
	Role         Role      `json:"role"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	DeletedAt    time.Time `json:"deletedAt"`
}
