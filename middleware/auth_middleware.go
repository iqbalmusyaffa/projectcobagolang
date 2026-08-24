package middleware

import (
	"net/http"
	"strings"

	"projectgolangnyoba/utils"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware memproteksi endpoint API agar hanya dapat diakses dengan JWT Token yang valid.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. Ambil header Authorization (format: "Bearer <token>")
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Header Authorization tidak ditemukan",
			})
			c.Abort()
			return
		}

		// 2. Pisahkan kata 'Bearer' dari tokennya
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "Format header Authorization harus: Bearer <token>",
			})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// 3. Validasi token menggunakan helper JWT
		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		// 4. Simpan userID dan userRole ke context agar bisa dibaca oleh Handler & RoleMiddleware
		c.Set("userID", claims.UserID)

		role := claims.Role
		if role == "" {
			role = "admin"
		}
		c.Set("userRole", role)

		c.Next()
	}
}
