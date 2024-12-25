package routes

import (
	"github.com/gin-gonic/gin"
	"unairsatu_v2/middlewares" // Import middleware
	"unairsatu_v2/controllers" // Import controllers untuk fungsi CRUD
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Group Route untuk Admin
	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middlewares.MockAuthMiddleware(), middlewares.CheckRole("admin"))
	{
		// Dashboard Route (route lama)
		adminRoutes.GET("/dashboard", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome Admin"})
		})

		// CRUD Users
		adminRoutes.GET("/users", controllers.GetAllUsers)
		adminRoutes.GET("/users/:id", controllers.GetUserById)
		adminRoutes.POST("/users", controllers.CreateUser)
		adminRoutes.PUT("/users/:id", controllers.UpdateUser)
		adminRoutes.DELETE("/users/:id", controllers.DeleteUser)

		// CRUD Modul
		adminRoutes.GET("/modul", controllers.GetAllModuls)
		adminRoutes.POST("/modul", controllers.CreateModul)
		adminRoutes.PUT("/modul/:id", controllers.UpdateModul)
		adminRoutes.DELETE("/modul/:id", controllers.DeleteModul)

		// CRUD Template_Modul
		adminRoutes.POST("/template-modul", controllers.AddTemplateModul)
		adminRoutes.DELETE("/template-modul/:id", controllers.RemoveTemplateModul) // Tambahkan route ini

		// Modul User
		adminRoutes.GET("/users/:id/modul", controllers.GetUserModuls)
		adminRoutes.PUT("/users/:id/change-jenis-user", controllers.ChangeUserType)

		// Modul Tambahan User
		adminRoutes.POST("/users/:id/add-modul", controllers.AddModulToUser)
		adminRoutes.DELETE("/users/:id/remove-modul", controllers.RemoveModulFromUser)
	}

	// Group Route untuk Mahasiswa
	mahasiswaRoutes := r.Group("/mahasiswa")
	mahasiswaRoutes.Use(middlewares.MockAuthMiddleware(), middlewares.CheckJenisUser("Mahasiswa"))
	{
		// Profile Route (route lama)
		mahasiswaRoutes.GET("/profile", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome Mahasiswa"})
		})
	}

	return r
}
