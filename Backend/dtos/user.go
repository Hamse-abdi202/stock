package dtos

import "Github.com/Hams/abdi/InventoryManagmentSystem/models"

type CreateUserDto struct {
	EmailAddress string      `json:"emailAddress" binding:"required,email"`
	Password     string      `json:"password" binding:"required,min=8,max=128"`
	FullName     string      `json:"fullName" binding:"required"`
	Role         models.Role `json:"role" binding:"required,oneof=ADMIN STUDENT_AFFAIRS CASHIER STUDENT"`
}

type LoginUserRequest struct {
	EmailAddress string `json:"emailAddress" binding:"required,email"`
	Password     string `json:"password" binding:"required,min=8,max=128"`
}

type LoginUserResponse struct {
	User         models.User `json:"user"`
	AccessToken  string      `json:"access_token"`
	RefreshToken string      `json:"refresh_token"`
}
