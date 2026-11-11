package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.GetString("role")
		if role == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message":    "You dont have required roles to do this action.",
				"is_success": false,
			})

			return
		}

		// for _, allowedRole := range allowedRoles {
		// 	if role == allowedRole {
		// 		c.Next()
		// 		return
		// 	}
		// }

		if slices.Contains(allowedRoles, role) {
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message":    "You dont have required roles to do this action.",
			"is_success": false,
		})

	}
}
