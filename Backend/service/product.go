
package services

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"Github.com/Hams/abdi/InventoryManagmentSystem/dtos"
	"Github.com/Hams/abdi/InventoryManagmentSystem/models"
	"Github.com/Hams/abdi/InventoryManagmentSystem/repository"
	"gorm.io/gorm"
)

type ProductService struct {
	repo *repository.ProductRepo
}

func RegisterProductService(repo *repository.ProductRepo) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (svc *ProductService) CreateProduct(data *dtos.CreateProductDto) (int, error) {
	name := strings.TrimSpace(data.Name)
	sku := strings.TrimSpace(data.SKU)

	if name == "" {
		return http.StatusBadRequest, errors.New("product name is required")
	}
	if sku == "" {
		return http.StatusBadRequest, errors.New("product SKU is required")
	}

	if data.PurchasePrice < 0 || data.SellingPrice < 0 {
		return http.StatusBadRequest, errors.New("prices cannot be negative")
	}

	if data.SellingPrice < data.PurchasePrice {
		return http.StatusBadRequest, errors.New("selling price cannot be less than purchase price")
	}

	if data.Quantity < 0 || data.LowStockThreshold < 0 {
		return http.StatusBadRequest, errors.New("quantity and low stock threshold cannot be negative")
	}

	exists, err := svc.repo.SKUExists(sku)
	if err != nil {
		slog.Error("Failed to check product SKU", "error", err)
		return http.StatusInternalServerError, errors.New("failed to create product")
	}

	if exists {
		return http.StatusConflict, errors.New("product SKU already exists")
	}

	err = svc.repo.ValidateProductCategory(data.CategoryID)
	if err != nil {
		slog.Error("Failed to validate category", "error", err)
		return http.StatusInternalServerError, errors.New("failed to create product")
	}

	status := strings.TrimSpace(data.Status)
	if status == "" {
		status = "active"
	}

	product := models.Product{
		Name:              name,
		SKU:               sku,
		CategoryID:        data.CategoryID,
		PurchasePrice:     data.PurchasePrice,
		SellingPrice:      data.SellingPrice,
		Quantity:          data.Quantity,
		LowStockThreshold: data.LowStockThreshold,
		Status:            status,
		Description:       strings.TrimSpace(data.Description),
	}

	err = svc.repo.CreateProduct(product)
	if err != nil {
		slog.Error("Failed to create product", "error", err)
		return http.StatusInternalServerError, errors.New("failed to create product")
	}

	slog.Info("Product successfully created")
	return http.StatusCreated, nil
}

func (svc *ProductService) GetProductByID(productID uint) (*models.Product, int, error) {
	
	product, err := svc.repo.GetProductByID(productID)
	if err != nil {
		slog.Error("Failed to get product", "error", err)
		return nil, http.StatusInternalServerError, errors.New("failed to get product")
	}

	return &product, http.StatusOK, nil
}

func (svc *ProductService) GetAllProducts() ([]models.Product, error) {
	products, err := svc.repo.List()
	if err != nil {
		slog.Error("Failed to get products", "error", err)
		return nil, errors.New("failed to get products")
	}

	return products, nil
}

func (svc *ProductService) UpdateProduct(productID uint, data *dtos.UpdateProductDto) (int, error) {
	if data == nil {
		return http.StatusBadRequest, errors.New("invalid request body")
	}

	name := strings.TrimSpace(data.Name)
	sku := strings.TrimSpace(data.SKU)

	if name == "" || sku == "" {
		return http.StatusBadRequest, errors.New("product name and SKU are required")
	}

	if data.PurchasePrice < 0 || data.SellingPrice < 0 {
		return http.StatusBadRequest, errors.New("prices cannot be negative")
	}

	if data.SellingPrice < data.PurchasePrice {
		return http.StatusBadRequest, errors.New("selling price cannot be less than purchase price")
	}

	if data.LowStockThreshold < 0 {
		return http.StatusBadRequest, errors.New("low stock threshold cannot be negative")
	}

	_, err := svc.repo.GetProductByID(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, errors.New("product not found")
		}

		slog.Error("Failed to get product", "error", err)
		return http.StatusInternalServerError, errors.New("failed to update product")
	}

	exists, err := svc.repo.SKUExistsForOtherProduct(sku, productID)
	if err != nil {
		slog.Error("Failed to check product SKU", "error", err)
		return http.StatusInternalServerError, errors.New("failed to update product")
	}

	if exists {
		return http.StatusConflict, errors.New("product SKU already exists")
	}

	err = svc.repo.ValidateProductCategory(data.CategoryID)
	if err != nil {
		if err.Error() == "category not found" {
			return http.StatusBadRequest, err
		}

		slog.Error("Failed to validate category", "error", err)
		return http.StatusInternalServerError, errors.New("failed to update product")
	}

	status := strings.TrimSpace(data.Status)
	if status == "" {
		status = "active"
	}

	product := models.Product{
		Name:              name,
		SKU:               sku,
		CategoryID:        data.CategoryID,
		PurchasePrice:     data.PurchasePrice,
		SellingPrice:      data.SellingPrice,
		LowStockThreshold: data.LowStockThreshold,
		Status:            status,
		Description:       strings.TrimSpace(data.Description),
	}

	err = svc.repo.UpdateProduct(productID, product)
	if err != nil {
		slog.Error("Failed to update product", "error", err)
		return http.StatusInternalServerError, errors.New("failed to update product")
	}

	slog.Info("Product successfully updated")
	return http.StatusOK, nil
}

func (svc *ProductService) DeleteProduct(productID uint) (int, error) {
	if productID == 0 {
		return http.StatusBadRequest, errors.New("invalid product ID")
	}

	_, err := svc.repo.GetProductByID(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, errors.New("product not found")
		}

		slog.Error("Failed to get product", "error", err)
		return http.StatusInternalServerError, errors.New("failed to delete product")
	}

	err = svc.repo.DeleteProduct(productID)
	if err != nil {
		slog.Error("Failed to delete product", "error", err)
		return http.StatusInternalServerError, errors.New("failed to delete product")
	}

	slog.Info("Product successfully deleted")
	return http.StatusOK, nil
}

func (svc *ProductService) GetProductCount() (int64, error) {
	count, err := svc.repo.GetProductCount()
	if err != nil {
		slog.Error("Failed to get product count", "error", err)
		return 0, errors.New("failed to get product count")
	}

	return count, nil
}
