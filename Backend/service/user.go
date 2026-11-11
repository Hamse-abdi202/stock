package services

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"Github.com/Hams/abdi/InventoryManagmentSystem/constants"
	"Github.com/Hams/abdi/InventoryManagmentSystem/dtos"
	"Github.com/Hams/abdi/InventoryManagmentSystem/helpers"
	"Github.com/Hams/abdi/InventoryManagmentSystem/models"
	"Github.com/Hams/abdi/InventoryManagmentSystem/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UserRepo
}

func RegisterUserService(repo *repository.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) CreateUser(data *dtos.CreateUserDto) (int, error) {
	email := strings.ToLower(data.EmailAddress)

	_, err := svc.repo.GetUserByEmail(email)
	if err == nil {
		errStr := "User with that email already exists"
		slog.Error(errStr)
		return http.StatusConflict, errors.New(errStr)
	}

	slog.Info("Hashing user password")

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to hash password")
		return http.StatusInternalServerError, errors.New(constants.DefaultErrorMsg)
	}

	data.Password = string(hashBytes)

	slog.Info("Creating new user")
	err = svc.repo.CreateUser(models.User{
		FullName:     data.FullName,
		EmailAddress: email,
		Password:     data.Password,
		Role:         data.Role,
	})

	if err != nil {
		slog.Error("failed to create user", "error", err)
		return http.StatusInternalServerError, errors.New(constants.FailedToCreateUser)
	}

	slog.Info("Successfully user was created")

	return http.StatusCreated, nil
}
func (svc *UserService) Getallusers() ([]models.User, error) {
	return svc.repo.List()
}
func (svc *UserService) LoginUser(data dtos.LoginUserRequest) (response *dtos.LoginUserResponse, statusCode int, err error) {
	slog.Info(	"Get user by email")
	email := strings.ToLower(data.EmailAddress)

	user, err := svc.repo.GetUserByEmail(email)
	if err != nil {
		slog.Error("Invalid email")
		statusCode = http.StatusUnauthorized
		err = errors.New(constants.InvalidCredentials)

		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		slog.Error("Invalid password")
		statusCode = http.StatusUnauthorized
		err = errors.New(constants.InvalidCredentials)

		return
	}

	accessToken, err := helpers.GenerateJWT(user.Role, user.EmailAddress, time.Now().Add(15*time.Minute).Unix(), false)
	if err != nil {
		slog.Error("Failed to generate access jwt token")
		statusCode = http.StatusInternalServerError
		err = errors.New(constants.DefaultErrorMsg)

		return
	}

	refreshToken, err := helpers.GenerateJWT(user.Role, user.EmailAddress, time.Now().Add(72*time.Hour).Unix(), true)
	if err != nil {
		slog.Error("Failed to generate refresh jwt token")
		statusCode = http.StatusInternalServerError
		err = errors.New(constants.DefaultErrorMsg)

		return
	}

	response = &dtos.LoginUserResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	statusCode = http.StatusOK
	err = nil

	err = svc.repo.UpdateLastLoginAt(user.ID)
	if err != nil {
		slog.Error("Failed to update last login at")
		return nil, http.StatusUnauthorized, errors.New(constants.DefaultErrorMsg)
	}

	return
}

func (svc *UserService) WhoAmI(email string) (*models.User, error) {
	user, err := svc.repo.GetUserByEmail(email)
	if err != nil {
		return nil, errors.New(constants.NotFoundError)
	}

	return &user, nil
}

func (svc *UserService) RefreshToken(email string) (*dtos.LoginUserResponse, int, error) {
	user, err := svc.repo.GetUserByEmail(email)
	if err != nil {
		return nil, http.StatusUnauthorized, errors.New(constants.DefaultErrorMsg)
	}

	accessToken, err := helpers.GenerateJWT(user.Role, user.EmailAddress, time.Now().Add(15*time.Minute).Unix(), false)
	if err != nil {
		slog.Error("Failed to generate access jwt token")
		return nil, http.StatusUnauthorized, errors.New(constants.DefaultErrorMsg)
	}

	refreshToken, err := helpers.GenerateJWT(user.Role, user.EmailAddress, time.Now().Add(72*time.Hour).Unix(), true)
	if err != nil {
		slog.Error("Failed to generate refresh jwt token")
		return nil, http.StatusUnauthorized, errors.New(constants.DefaultErrorMsg)
	}

	err = svc.repo.UpdateLastLoginAt(user.ID)
	if err != nil {
		slog.Error("Failed to update last login at")
		return nil, http.StatusUnauthorized, errors.New(constants.DefaultErrorMsg)
	}

	return &dtos.LoginUserResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, http.StatusOK, nil

}
