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
	fmt.Println("=== Memulai Aplikasi Auth Clean Architecture ===")

	// 1. Inisialisasi Database PostgreSQL
	db := config.InitDB()

	// 2. Dependency Injection Wiring (Database -> Repository -> Usecase -> Handler)
	userRepo := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// 3. Inisialisasi Router Gin & Load HTML Templates
	r := gin.Default()
	r.LoadHTMLGlob("views/*")

	// Web Routes (Frontend Pages)
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/login")
	})
	r.GET("/login", userHandler.ShowLoginPage)
	r.GET("/register", userHandler.ShowRegisterPage)
	r.GET("/dashboard", userHandler.ShowDashboardPage)

	// API Routes Publik (Tanpa Autentikasi)
	api := r.Group("/api")
	{
		api.POST("/register", userHandler.Register)
		api.POST("/login", userHandler.Login)
	}

	// API Routes Privat (Wajib Menggunakan Header 'Authorization: Bearer <token>')
	protected := r.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", userHandler.Profile)
		protected.PUT("/profile", userHandler.UpdateProfile)
		protected.PUT("/change-password", userHandler.ChangePassword)
	}

	// 4. Jalankan Server HTTP pada port 8080
	port := ":8080"
	fmt.Printf("Server berjalan di http://localhost%s\n", port)
	if err := r.Run(port); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}
