package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Middleware untuk memeriksa role pengguna
func CheckRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil role pengguna dari context
		userRole := c.GetString("role")

		// Log untuk debugging (opsional)
		// fmt.Println("Role pengguna:", userRole)

		// Jika role pengguna tidak sesuai, kembalikan error 403 (Forbidden)
		if userRole != role {
			c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
			c.Abort()
			return
		}

		// Jika role sesuai, lanjutkan ke handler berikutnya
		c.Next()
	}
}

// Middleware untuk memeriksa jenis user
func CheckJenisUser(ju string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Ambil jenis_user pengguna dari context
		userJenis := c.GetString("jenis_user")

		// Log untuk debugging (opsional)
		// fmt.Println("Jenis user:", userJenis)

		// Jika jenis_user pengguna tidak sesuai, kembalikan error 403 (Forbidden)
		if userJenis != ju {
			c.JSON(http.StatusForbidden, gin.H{"message": "Access denied"})
			c.Abort()
			return
		}

		// Jika jenis_user sesuai, lanjutkan ke handler berikutnya
		c.Next()
	}
}

// Middleware simulasi untuk testing (opsional)
func MockAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Simulasi nilai role dan jenis_user untuk testing
		c.Set("role", "admin")          // Ubah sesuai kebutuhan, misalnya: "admin", "civitas", dll.
		c.Set("jenis_user", "Mahasiswa") // Ubah sesuai kebutuhan, misalnya: "Mahasiswa", "Dosen", dll.

		// Lanjutkan ke middleware atau handler berikutnya
		c.Next()
	}
}
