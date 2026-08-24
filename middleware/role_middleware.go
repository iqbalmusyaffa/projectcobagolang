package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RoleMiddleware memvalidasi apakah peran pengguna (userRole) diizinkan mengakses endpoint.
func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleVal, exists := c.Get("userRole")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  "error",
				"message": "Akses ditolak: Peran pengguna tidak ditemukan",
			})
			c.Abort()
			return
		}

		userRole := roleVal.(string)

		// Cek apakah userRole termasuk dalam daftar allowedRoles
		isAllowed := false
		for _, allowed := range allowedRoles {
			if userRole == allowed {
				isAllowed = true
				break
			}
		}

		if !isAllowed {
			c.JSON(http.StatusForbidden, gin.H{
				"status":  "error",
				"message": fmt.Sprintf("Akses ditolak: Peran '%s' tidak memiliki izin untuk mengakses fitur ini", userRole),
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
