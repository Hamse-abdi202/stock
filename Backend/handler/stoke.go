package handlers

import (
	"net/http"
	"strconv"

	"Github.com/Hams/abdi/InventoryManagmentSystem/dtos"
	"Github.com/Hams/abdi/InventoryManagmentSystem/infra"
	"Github.com/Hams/abdi/InventoryManagmentSystem/repository"
	services "Github.com/Hams/abdi/InventoryManagmentSystem/service"
	"github.com/gin-gonic/gin"
)

type StockMovementHandler struct {
	services *services.StockMovementService
}

func RegisterStockMovementHandler() *StockMovementHandler {
	repo := repository.RegisterStockMovementRepo(infra.DB)
	svc := services.RegisterStockMovementService(repo)

	return &StockMovementHandler{
		services: svc,
	}
}

func (h *StockMovementHandler) CreateStockMovement(c *gin.Context) {
	var data dtos.CreateStockMovementDto

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized user",
		})
		return
	}

	var userID uint

	switch id := userIDValue.(type) {
	case uint:
		userID = id
	case int:
		if id > 0 {
			userID = uint(id)
		}
	case uint64:
		userID = uint(id)
	case int64:
		if id > 0 {
			userID = uint(id)
		}
	case string:
		parsedID, err := strconv.Atoi(id)
		if err == nil && parsedID > 0 {
			userID = uint(parsedID)
		}
	}

	if userID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user ID",
		})
		return
	}

	statusCode, err := h.services.CreateStockMovement(&data, userID)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message": "stock movement created successfully",
	})
}

func (h *StockMovementHandler) GetAllStockMovements(c *gin.Context) {
	movements, err := h.services.GetAllStockMovements()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve stock movements",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": movements,
	})
}

func (h *StockMovementHandler) GetStockMovementByID(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid stock movement ID",
		})
		return
	}

	movement, statusCode, err := h.services.GetStockMovementByID(uint(id))
	if err != nil {
		c.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": movement,
	})
}

func (h *StockMovementHandler) GetProductStockMovements(c *gin.Context) {
	productIDStr := c.Param("productID")

	productID, err := strconv.Atoi(productIDStr)
	if err != nil || productID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid product ID",
		})
		return
	}

	movements, err := h.services.GetProductStockMovements(uint(productID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve product stock movements",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": movements,
	})
}