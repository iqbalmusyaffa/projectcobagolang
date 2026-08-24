package config

import (
	"fmt"
	"log"
	"os"

	"projectgolangnyoba/entity"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB membuat dan mengembalikan instance koneksi database PostgreSQL via GORM.
func InitDB() *gorm.DB {
	// Memuat variabel dari file .env (jika file .env ada)
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Info: File .env tidak ditemukan, menggunakan environment variables sistem.")
	}

	// Ambil variabel dari .env / sistem
	host := getEnv("DB_HOST", "localhost")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "golang_db")
	port := getEnv("DB_PORT", "5432")
	sslmode := getEnv("DB_SSLMODE", "disable")

	// DSN (Data Source Name) untuk PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Jakarta",
		host, user, password, dbname, port, sslmode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database PostgreSQL: %v", err)
	}

	fmt.Println("Berhasil terhubung ke database PostgreSQL!")

	// Auto-Migrate tabel 'users' otomatis berdasarkan struct entity.User
	err = db.AutoMigrate(&entity.User{})
	if err != nil {
		log.Fatalf("Gagal melakukan migrasi database: %v", err)
	}

	fmt.Println("Migrasi tabel database berhasil!")

	return db
}

// getEnv adalah helper untuk mengambil nilai ENV atau nilai defaultnya
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
