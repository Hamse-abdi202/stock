package middleware

import (
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"Github.com/Hams/abdi/InventoryManagmentSystem/infra"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing or invalid authentication header",
			})
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		secret := []byte(infra.Configurations.AccessJwtTokenSecret)

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("unexpected signing method")
			}

			return secret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid or expired token",
			})
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token claims",
			})
			return
		}

		email, _ := claims["sub"].(string)
		role, _ := claims["role"].(string)
		userID, ok := claims["user_id"].(float64)

		if !ok || userID <= 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "user ID missing from token",
			})
			return
		}

		c.Set("userID", uint(userID))
		c.Set("email", email)
		c.Set("role", role)

		slog.Info("Authenticated user", "email", email, "userID", uint(userID))

		c.Next()
	}
}