package handlers

import (
	"log/slog"
	"net/http"

	"Github.com/Hams/abdi/InventoryManagmentSystem/dtos"
	"Github.com/Hams/abdi/InventoryManagmentSystem/infra"
	"Github.com/Hams/abdi/InventoryManagmentSystem/repository"
	services "Github.com/Hams/abdi/InventoryManagmentSystem/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func RegisterUserHandler() *UserHandler {

	userRepo := repository.RegisterUserRepo(infra.DB)
	userSvc := services.RegisterUserService(userRepo)

	return &UserHandler{
		userService: userSvc,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var requestBody dtos.CreateUserDto

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Failed bind body request",
			"is_success": false,
			"error":      err.Error(),
		})
		return
	}

	statusCode, err := h.userService.CreateUser(&requestBody)

	if err != nil {
		c.JSON(statusCode, gin.H{
			"is_success": false,
			"message":    err.Error(),
		})

		return
	}

	c.JSON(statusCode, gin.H{
		"message":    "User created successfully!",
		"is_success": true,
	})

}

func (h *UserHandler) Getallusers(c *gin.Context){
	users,err:=h.userService.Getallusers()
	if err!=nil{
	c.JSON(http.StatusUnauthorized,gin.H{
		"message":"failed too list users",
	})
	return
	}
	c.JSON(http.StatusOK,gin.H{
		"message":"successfully listed this users",
	"data":users,
	})
}

func (h *UserHandler) LoginUser(c *gin.Context) {
	var requestBody dtos.LoginUserRequest

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Failed bind body request",
			"is_success": false,
		})
		return
	}

	resp, statusCode, err := h.userService.LoginUser(requestBody)

	if err != nil {
		c.JSON(statusCode, gin.H{
			"message":    err.Error(),
			"is_success": false,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message":    "User logged in successfully",
		"is_success": true,
		"data":       resp,
	})

}

func (h *UserHandler) WhoAmI(c *gin.Context) {

	email := c.GetString("email")

	user, err := h.userService.WhoAmI(email)

	if err != nil {

		c.JSON(http.StatusOK, gin.H{
			"message":    err.Error(),
			"is_success": false,
			"data":       nil,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "success",
		"is_success": true,
		"data":       user,
	})
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	email := c.GetString("email")

	data, statusCode, err := h.userService.RefreshToken(email)

	if err != nil {

		slog.Info("Failed to refresh token", "error", err.Error())

		c.JSON(http.StatusOK, gin.H{
			"message":    "Unauthorized",
			"is_success": false,
			"data":       nil,
		})

		return
	}

	c.JSON(statusCode, gin.H{
		"message":    "User refreshed token successfully",
		"is_success": true,
		"data":       data,
	})

}
