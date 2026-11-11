package repository

import (
	"time"

	"Github.com/Hams/abdi/InventoryManagmentSystem/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func RegisterUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

func (repo *UserRepo) CreateUser(data models.User) error {
	return repo.DB.Create(&data).Error
}

func (repo *UserRepo) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	err := repo.DB.Where("email_address = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repo *UserRepo) UpdateLastLoginAt(userID uint) error {
	// When you want to update one column
	// err := repo.DB.Model(&models.User{}).Update("last_login_at", time.Now()).Error

	// When you want to update multiple columns
	err := repo.DB.Model(&models.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"last_login_at": time.Now(),
	}).Error

	if err != nil {
		return err
	}

	return nil
}


func (repo *UserRepo) List() ([]models.User,error){
	var Users []models.User

	err:=repo.DB.Find(&Users).Error
	if err!=nil{
		return  nil,err
	}

	return  Users,nil
}