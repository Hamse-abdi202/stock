
package repository

import (
	"errors"

	"Github.com/Hams/abdi/InventoryManagmentSystem/models"
	"gorm.io/gorm"
)

type StockMovementRepo struct {
	DB *gorm.DB
}

func RegisterStockMovementRepo(db *gorm.DB) *StockMovementRepo {
	return &StockMovementRepo{
		DB: db,
	}
}

func (repo *StockMovementRepo) CreateStockMovement(data models.StockMovement) error {
	return repo.DB.Transaction(func(tx *gorm.DB) error {
		var product models.Product

		err := tx.Clauses().First(&product, data.ProductID).Error
		if err != nil {
			return err
		}

		if data.Quantity <= 0 {
			return errors.New("quantity must be greater than zero")
		}

		if data.Type == "IN" {
			product.Quantity += data.Quantity
		} else if data.Type == "OUT" {
			if product.Quantity < data.Quantity {
				return errors.New("insufficient stock quantity")
			}

			product.Quantity -= data.Quantity
		} else {
			return errors.New("invalid stock movement type")
		}

		err = tx.Model(&product).Update("quantity", product.Quantity).Error
		if err != nil {
			return err
		}

		err = tx.Create(&data).Error
		if err != nil {
			return err
		}

		return nil
	})
}

func (repo *StockMovementRepo) GetStockMovementByID(movementID uint) (models.StockMovement, error) {
	var movement models.StockMovement

	err := repo.DB.Preload("Product").
		Preload("User").
		First(&movement, movementID).Error

	if err != nil {
		return models.StockMovement{}, err
	}

	return movement, nil
}

func (repo *StockMovementRepo) GetAllStockMovements() ([]models.StockMovement, error) {
	var movements []models.StockMovement

	err := repo.DB.Preload("Product").
		Preload("User").
		Order("created_at DESC").
		Find(&movements).Error

	if err != nil {
		return nil, err
	}

	return movements, nil
}

func (repo *StockMovementRepo) GetProductStockMovements(productID uint) ([]models.StockMovement, error) {
	var movements []models.StockMovement

	err := repo.DB.Preload("Product").
		Preload("User").
		Where("product_id = ?", productID).
		Order("created_at DESC").
		Find(&movements).Error

	if err != nil {
		return nil, err
	}

	return movements, nil
}

func (repo *StockMovementRepo) GetProductQuantity(productID uint) (int, error) {
	var product models.Product

	err := repo.DB.Select("quantity").
		First(&product, productID).Error

	if err != nil {
		return 0, err
	}

	return product.Quantity, nil
}
