package main

import (
	"fmt"
	"log"

	"projectgolangnyoba/config"
	"projectgolangnyoba/handler"
	"projectgolangnyoba/middleware"
	"projectgolangnyoba/repository"
	"projectgolangnyoba/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("=== Memulai Server REST API Golang (Clean Architecture) ===")

	// 1. Inisialisasi Database PostgreSQL & Redis
	db := config.InitDB()
	rdb := config.InitRedis()

	// 2. Dependency Injection Wiring (Database/Redis -> Repository -> Usecase -> Handler)
	userRepo := repository.NewUserRepository(db)
	redisRepo := repository.NewRedisRepository(rdb)

	userUsecase := usecase.NewUserUsecase(userRepo, redisRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// 3. Inisialisasi Router Gin & Global Middleware
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.LoadHTMLGlob("views/*")

	// Web Routes (API Status & Auto-Redirect ke Vue 3 Frontend)
	r.GET("/", userHandler.ShowAPIStatus)
	r.GET("/login", func(c *gin.Context) {
		c.Redirect(302, "http://localhost:5173/login")
	})
	r.GET("/register", func(c *gin.Context) {
		c.Redirect(302, "http://localhost:5173/register")
	})
	r.GET("/dashboard", func(c *gin.Context) {
		c.Redirect(302, "http://localhost:5173/dashboard")
	})

	// API Routes Publik (Tanpa Autentikasi)
	api := r.Group("/api")
	{
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
		api.POST("/refresh-token", userHandler.RefreshToken)
	}

	// API Routes Privat (Wajib Menggunakan Header 'Authorization: Bearer <token>')
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware(redisRepo))
	{
		protected.POST("/logout", userHandler.Logout)
		protected.GET("/profile", userHandler.Profile)
		protected.PUT("/profile", userHandler.UpdateProfile)
		protected.PUT("/change-password", userHandler.ChangePassword)

		// Rute Khusus Peran Superadmin & Owner
		adminGroup := protected.Group("/admin")
		adminGroup.Use(middleware.RoleMiddleware("superadmin", "owner"))
		{
			adminGroup.GET("/users", userHandler.GetAllUsers)
		}

		// Rute Khusus Peran Superadmin Sahaja
		superadminGroup := protected.Group("/superadmin")
		superadminGroup.Use(middleware.RoleMiddleware("superadmin"))
		{
			superadminGroup.PUT("/change-role", userHandler.ChangeRole)
		}
	}

	// 4. Jalankan Server HTTP pada port 8080
	port := ":8080"
	fmt.Printf("REST API Golang berjalan di http://localhost%s\n", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
