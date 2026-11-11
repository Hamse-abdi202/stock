
package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"Github.com/Hams/abdi/InventoryManagmentSystem/dtos"
	"Github.com/Hams/abdi/InventoryManagmentSystem/infra"
	"Github.com/Hams/abdi/InventoryManagmentSystem/repository"
	services "Github.com/Hams/abdi/InventoryManagmentSystem/service"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	categoryService *services.CategoryService
}

func RegisterCategoryHandler() *CategoryHandler {
	categoryRepo := repository.RegisterCategoryRepo(infra.DB)
	categorySvc := services.RegisterCategoryService(categoryRepo)

	return &CategoryHandler{
		categoryService: categorySvc,
	}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var requestBody dtos.CreateCategoryDto

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Failed to bind request body",
			"is_success": false,
			"error":      err.Error(),
		})
		return
	}

	statusCode, err := h.categoryService.CreateCategory(&requestBody)

	if err != nil {
		c.JSON(statusCode, gin.H{
			"message":    err.Error(),
			"is_success": false,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message":    "Category created successfully",
		"is_success": true,
	})
}

func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.categoryService.GetAllCategories()

	if err != nil {
		slog.Error("Failed to list categories", "error", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"message":    "Failed to list categories",
			"is_success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Categories listed successfully",
		"is_success": true,
		"data":       categories,
	})
}

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Failed to get category ID",
			"is_success": false,
			"error":      "Invalid category ID",
		})
		return
	}

	category, statusCode, err := h.categoryService.GetCategoryByID(uint(id))

	if err != nil {
		c.JSON(statusCode, gin.H{
			"message":    err.Error(),
			"is_success": false,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message":    "Category retrieved successfully",
		"is_success": true,
		"data":       category,
	})
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Failed to get category ID",
			"is_success": false,
			"error":      "Invalid category ID",
		})
		return
	}

	var requestBody dtos.UpdateCategoryDto

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Failed to bind request body",
			"is_success": false,
			"error":      err.Error(),
		})
		return
	}

	statusCode, err := h.categoryService.UpdateCategory(uint(id), &requestBody)

	if err != nil {
		c.JSON(statusCode, gin.H{
			"message":    err.Error(),
			"is_success": false,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message":    "Category updated successfully",
		"is_success": true,
	})
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Failed to get category ID",
			"is_success": false,
			"error":      "Invalid category ID",
		})
		return
	}

	statusCode, err := h.categoryService.DeleteCategory(uint(id))

	if err != nil {
		c.JSON(statusCode, gin.H{
			"message":    err.Error(),
			"is_success": false,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message":    "Category deleted successfully",
		"is_success": true,
	})
}
