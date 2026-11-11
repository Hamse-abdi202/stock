
package helpers

import (
	"time"

	"Github.com/Hams/abdi/InventoryManagmentSystem/infra"
	"Github.com/Hams/abdi/InventoryManagmentSystem/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(role models.Role, sub string, expiresIn int64, isRefreshToken bool) (string, error) {
	config := infra.Configurations

	var jwtSecret []byte

	if isRefreshToken {
		jwtSecret = []byte(config.RefreshJwtTokenSecret)
	} else {
		jwtSecret = []byte(config.AccessJwtTokenSecret)
	}

	claims := jwt.MapClaims{
		"sub":            sub,
		"nbf":            time.Now().Unix(),
		"exp":            expiresIn,
		"isRefreshToken": isRefreshToken,
		"role":           role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}
