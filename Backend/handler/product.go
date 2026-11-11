
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

type ProductHandler struct {
	productService *services.ProductService
}

func RegisterProductHandler() *ProductHandler {
	productRepo := repository.RegisterProductRepo(infra.DB)
	productSvc := services.RegisterProductService(productRepo)

	return &ProductHandler{
		productService: productSvc,
	}
}

func (h *ProductHandler) CreateProduct(c *gin.Context) {
	var requestBody dtos.CreateProductDto

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Failed to bind request body",
			"is_success": false,
			"error":      err.Error(),
		})
		return
	}

	statusCode, err := h.productService.CreateProduct(&requestBody)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message":    err.Error(),
			"is_success": false,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message":    "Product created successfully",
		"is_success": true,
	})
}

func (h *ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.productService.GetAllProducts()
	if err != nil {
		slog.Error("Failed to list products", "error", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":    "Failed to list products",
			"is_success": false,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "Products listed successfully",
		"is_success": true,
		"data":       products,
	})
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Invalid product ID",
			"is_success": false,
		})
		return
	}

	product, statusCode, err := h.productService.GetProductByID(uint(id))
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message":    err.Error(),
			"is_success": false,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message":    "Product retrieved successfully",
		"is_success": true,
		"data":       product,
	})
}

func (h *ProductHandler) UpdateProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Invalid product ID",
			"is_success": false,
		})
		return
	}

	var requestBody dtos.UpdateProductDto

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Failed to bind request body",
			"is_success": false,
			"error":      err.Error(),
		})
		return
	}

	statusCode, err := h.productService.UpdateProduct(uint(id), &requestBody)
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message":    err.Error(),
			"is_success": false,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message":    "Product updated successfully",
		"is_success": true,
	})
}

func (h *ProductHandler) DeleteProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Invalid product ID",
			"is_success": false,
		})
		return
	}

	statusCode, err := h.productService.DeleteProduct(uint(id))
	if err != nil {
		c.JSON(statusCode, gin.H{
			"message":    err.Error(),
			"is_success": false,
		})
		return
	}

	c.JSON(statusCode, gin.H{
		"message":    "Product deleted successfully",
		"is_success": true,
	})
}
