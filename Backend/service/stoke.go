
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

type StockMovementService struct {
	repo *repository.StockMovementRepo
}

func RegisterStockMovementService(repo *repository.StockMovementRepo) *StockMovementService {
	return &StockMovementService{
		repo: repo,
	}
}

func (svc *StockMovementService) CreateStockMovement(data *dtos.CreateStockMovementDto, userID uint) (int, error) {


	if data.Quantity <= 0 {
		return http.StatusBadRequest, errors.New("quantity must be greater than zero")
	}

	movementType := strings.ToUpper(data.Type)

	if movementType != "IN" && movementType != "OUT" {
		return http.StatusBadRequest, errors.New("movement type must be IN or OUT")
	}

	if userID == 0 {
		return http.StatusUnauthorized, errors.New("unauthorized user")
	}

	_, err := svc.repo.GetProductQuantity(data.ProductID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, errors.New("product not found")
		}

		slog.Error("Failed to find product", "error", err)
		return http.StatusInternalServerError, errors.New("failed to create stock movement")
	}

	movement := models.StockMovement{
		ProductID: data.ProductID,
		Type:      movementType,
		Quantity:  data.Quantity,
		Note:      strings.TrimSpace(data.Note),
		CreatedBy: userID,
	}

	err = svc.repo.CreateStockMovement(movement)
	if err != nil {
		if err.Error() == "insufficient stock quantity" {
			return http.StatusConflict, err
		}

		if err.Error() == "invalid stock movement type" {
			return http.StatusBadRequest, err
		}

		slog.Error("Failed to create stock movement", "error", err)
		return http.StatusInternalServerError, errors.New("failed to create stock movement")
	}

	slog.Info("Stock movement successfully created",
		"product_id", data.ProductID,
		"type", movementType,
		"quantity", data.Quantity,
		"created_by", userID,
	)

	return http.StatusCreated, nil
}

func (svc *StockMovementService) GetStockMovementByID(movementID uint) (*models.StockMovement, int, error) {
	if movementID == 0 {
		return nil, http.StatusBadRequest, errors.New("invalid stock movement ID")
	}

	movement, err := svc.repo.GetStockMovementByID(movementID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, http.StatusNotFound, errors.New("stock movement not found")
		}

		slog.Error("Failed to get stock movement", "error", err)
		return nil, http.StatusInternalServerError, errors.New("failed to get stock movement")
	}

	return &movement, http.StatusOK, nil
}

func (svc *StockMovementService) GetAllStockMovements() ([]models.StockMovement, error) {
	movements, err := svc.repo.GetAllStockMovements()
	if err != nil {
		slog.Error("Failed to get stock movements", "error", err)
		return nil, errors.New("failed to get stock movements")
	}

	return movements, nil
}

func (svc *StockMovementService) GetProductStockMovements(productID uint) ([]models.StockMovement, error) {
	if productID == 0 {
		return nil, errors.New("invalid product ID")
	}

	_, err := svc.repo.GetProductQuantity(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}

		slog.Error("Failed to find product", "error", err)
		return nil, errors.New("failed to get product stock movements")
	}

	movements, err := svc.repo.GetProductStockMovements(productID)
	if err != nil {
		slog.Error("Failed to get product stock movements", "error", err)
		return nil, errors.New("failed to get product stock movements")
	}

	return movements, nil
}
