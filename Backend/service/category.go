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

type CategoryService struct {
	repo *repository.CategoryRepo
}

func RegisterCategoryService(repo *repository.CategoryRepo) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (svc *CategoryService) CreateCategory(data *dtos.CreateCategoryDto) (int, error) {
	name := strings.TrimSpace(data.Name)
	_, err := svc.repo.GetCategoryByName(name)
	if err == nil {
		slog.Error("Category already exists")
		return http.StatusConflict, errors.New("category already exists")
	}


	slog.Info("Creating new category")

	err = svc.repo.CreateCategory(models.Category{
		Name: name,
	})

	if err != nil {
		slog.Error("Failed to create category", "error", err)
		return http.StatusInternalServerError, errors.New("failed to create category")
	}

	slog.Info("Category successfully created")

	return http.StatusCreated, nil
}

func (svc *CategoryService) GetCategoryByID(categoryID uint) (*models.Category, int, error) {
	if categoryID == 0 {
		return nil, http.StatusBadRequest, errors.New("invalid category id")
	}

	category, err := svc.repo.GetCategoryByID(categoryID)
	if err != nil {
		slog.Error("Failed to get category", "error", err)
		return nil, http.StatusInternalServerError, errors.New("failed to get category")
	}

	return &category, http.StatusOK, nil
}

func (svc *CategoryService) GetAllCategories() ([]models.Category, error) {
	categories, err := svc.repo.List()
	if err != nil {
		slog.Error("Failed to get categories", "error", err)
		return nil, errors.New("failed to get categories")
	}

	return categories, nil
}

func (svc *CategoryService) UpdateCategory(categoryID uint, data *dtos.UpdateCategoryDto) (int, error) {

	name := strings.TrimSpace(data.Name)
	_, err := svc.repo.GetCategoryByID(categoryID)
	if err != nil {
		slog.Error("Failed to get category", "error", err)
		return http.StatusInternalServerError, errors.New("failed to update category")
	}

	existingCategory, err := svc.repo.GetCategoryByName(name)
	if err == nil && existingCategory.ID != categoryID {
		return http.StatusConflict, errors.New("category name already exists")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		slog.Error("Failed to check category name", "error", err)
		return http.StatusInternalServerError, errors.New("failed to update category")
	}

	err = svc.repo.UpdateCategory(categoryID, models.Category{
		Name: name,
	})

	if err != nil {
		slog.Error("Failed to update category", "error", err)
		return http.StatusInternalServerError, errors.New("failed to update category")
	}

	slog.Info("Category successfully updated")

	return http.StatusOK, nil
}

func (svc *CategoryService) DeleteCategory(categoryID uint) (int, error) {
	if categoryID == 0 {
		return http.StatusBadRequest, errors.New("invalid category id")
	}

	_, err := svc.repo.GetCategoryByID(categoryID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return http.StatusNotFound, errors.New("category not found")
		}

		slog.Error("Failed to get category", "error", err)
		return http.StatusInternalServerError, errors.New("failed to delete category")
	}

	count, err := svc.repo.GetProductCount(categoryID)
	if err != nil {
		slog.Error("Failed to count category products", "error", err)
		return http.StatusInternalServerError, errors.New("failed to delete category")
	}

	if count > 0 {
		return http.StatusConflict, errors.New("cannot delete category because products exist")
	}

	err = svc.repo.DeleteCategory(categoryID)
	if err != nil {
		slog.Error("Failed to delete category", "error", err)
		return http.StatusInternalServerError, errors.New("failed to delete category")
	}

	slog.Info("Category successfully deleted")

	return http.StatusNoContent, nil
}

func (svc *CategoryService) GetProductCount(categoryID uint) (int64, error) {
	count, err := svc.repo.GetProductCount(categoryID)
	if err != nil {
		slog.Error("Failed to get product count", "error", err)
		return 0, errors.New("failed to get product count")
	}

	return count, nil
}
