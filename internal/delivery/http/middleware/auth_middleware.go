package middleware

import (
	"infared-backend/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthGuard() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Akses ditolak", "Token tidak ditemukan atau format salah")
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			utils.ErrorResponse(c, http.StatusUnauthorized, "Sesi tidak valid atau telah kedaluwarsa", err.Error())
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Set("role", claims["role"])

		c.Next()
	}
}

func RoleGuard(allowedRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")

		if !exists || userRole.(string) != allowedRole {
			utils.ErrorResponse(c, http.StatusForbidden, "Akses ditolak", "Anda tidak memiliki izin (bukan admin) untuk mengakses fitur ini")
			c.Abort()
			return
		}

		c.Next()
	}
}
