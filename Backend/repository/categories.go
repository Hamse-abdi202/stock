
package repository

import (
	"errors"

	"Github.com/Hams/abdi/InventoryManagmentSystem/models"
	"gorm.io/gorm"
)

type CategoryRepo struct {
	DB *gorm.DB
}

func RegisterCategoryRepo(db *gorm.DB) *CategoryRepo {
	return &CategoryRepo{
		DB: db,
	}
}

func (repo *CategoryRepo) CreateCategory(data models.Category) error {
	return repo.DB.Create(&data).Error
}

func (repo *CategoryRepo) GetCategoryByID(categoryID uint) (models.Category, error) {
	var category models.Category

	err := repo.DB.First(&category, categoryID).Error
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (repo *CategoryRepo) GetCategoryByName(name string) (models.Category, error) {
	var category models.Category

	err := repo.DB.Where("name = ?", name).First(&category).Error
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (repo *CategoryRepo) UpdateCategory(categoryID uint, data models.Category) error {
	return repo.DB.Model(&models.Category{}).
		Where("id = ?", categoryID).
		Updates(map[string]interface{}{
			"name": data.Name,
		}).Error
}

func (repo *CategoryRepo) DeleteCategory(categoryID uint) error {
	count, err := repo.GetProductCount(categoryID)
	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("cannot delete category because products exist")
	}

	result := repo.DB.Delete(&models.Category{}, categoryID)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (repo *CategoryRepo) List() ([]models.Category, error) {
	var categories []models.Category

	err := repo.DB.Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (repo *CategoryRepo) GetProductCount(categoryID uint) (int64, error) {
	var count int64

	err := repo.DB.Model(&models.Product{}).
		Where("category_id = ?", categoryID).
		Count(&count).Error

	return count, err
}

func (repo *CategoryRepo) CategoryExists(categoryID uint) (bool, error) {
	var count int64

	err := repo.DB.Model(&models.Category{}).
		Where("id = ?", categoryID).
		Count(&count).Error

	return count > 0, err
}
