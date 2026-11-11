
package repository

import (
	"errors"

	"Github.com/Hams/abdi/InventoryManagmentSystem/models"
	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func RegisterProductRepo(db *gorm.DB) *ProductRepo {
	return &ProductRepo{
		DB: db,
	}
}

func (repo *ProductRepo) CreateProduct(data models.Product) error {
	return repo.DB.Create(&data).Error
}

func (repo *ProductRepo) GetProductByID(productID uint) (models.Product, error) {
	var product models.Product

	err := repo.DB.Preload("Category").
		First(&product).Error

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (repo *ProductRepo) GetProductBySKU(sku string) (models.Product, error) {
	var product models.Product

	err := repo.DB.Where("sku = ?", sku).
		First(&product).Error

	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (repo *ProductRepo) UpdateProduct(productID uint, data models.Product) error {
	result := repo.DB.Model(&models.Product{}).
		Where("id = ?", productID).
		Updates(map[string]interface{}{
			"name":                data.Name,
			"sku":                 data.SKU,
			"category_id":         data.CategoryID,
			"purchase_price":      data.PurchasePrice,
			"selling_price":       data.SellingPrice,
			"low_stock_threshold": data.LowStockThreshold,
			"status":              data.Status,
			"description":         data.Description,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo *ProductRepo) DeleteProduct(productID uint) error {
	result := repo.DB.Delete(&models.Product{}, productID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo *ProductRepo) List() ([]models.Product, error) {
	var products []models.Product

	err := repo.DB.Preload("Category").Find(&products).Error

	if err != nil {
		return nil, err
	}

	return products, nil
}

func (repo *ProductRepo) GetProductCount() (int64, error) {
	var count int64

	err := repo.DB.Model(&models.Product{}).
		Count(&count).Error

	return count, err
}

func (repo *ProductRepo) CategoryExists(categoryID uint) (bool, error) {
	var count int64

	err := repo.DB.Model(&models.Category{}).
		Where("id = ?", categoryID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *ProductRepo) SKUExists(sku string) (bool, error) {
	var count int64

	err := repo.DB.Model(&models.Product{}).
		Where("sku = ?", sku).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *ProductRepo) SKUExistsForOtherProduct(sku string, productID uint) (bool, error) {
	var count int64

	err := repo.DB.Model(&models.Product{}).
		Where("sku = ? AND id <> ?", sku, productID).
		Count(&count).Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (repo *ProductRepo) ValidateProductCategory(categoryID uint) error {
	exists, err := repo.CategoryExists(categoryID)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New("category not found")
	}

	return nil
}
