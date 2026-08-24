package config

import (
	"fmt"
	"log"
	"time"

	"projectgolangnyoba/entity"
	"projectgolangnyoba/utils"

	"gorm.io/gorm"
)

// SeedSuperAdmin & SeedAllData secara otomatis menyisipkan data awal (Seeder) ke PostgreSQL jika belum ada.
func SeedSuperAdmin(db *gorm.DB) {
	// 1. Seed Akun Super Admin Utama
	seedUser(db, entity.User{
		Name:      "Farid Wimansyah (Super Admin)",
		Email:     "faridwimansyah8@gmail.com",
		Password:  "password123",
		Role:      "superadmin",
		Phone:     "081234567890",
		Gender:    "Laki-laki",
		Address:   "Jl. Sudirman No. 1, Jakarta Pusat",
		Bio:       "Head of System Administrator & Lead Developer",
	})

	// 2. Seed Akun Owner (Pemilik Laporan)
	seedUser(db, entity.User{
		Name:      "Budi Santoso (Owner)",
		Email:     "owner@demo.com",
		Password:  "password123",
		Role:      "owner",
		Phone:     "085678901234",
		Gender:    "Laki-laki",
		Address:   "Jl. Gatot Subroto No. 45, Jakarta Selatan",
		Bio:       "Owner & Executive Stakeholder",
	})

	// 3. Seed Akun Admin Standar
	seedUser(db, entity.User{
		Name:      "Siti Rahma (Admin)",
		Email:     "admin@demo.com",
		Password:  "password123",
		Role:      "admin",
		Phone:     "087890123456",
		Gender:    "Perempuan",
		Address:   "Jl. M.H. Thamrin No. 12, Jakarta Pusat",
		Bio:       "Operational Staff & General Admin",
	})

	// 4. Seed Akun User Regular (Pengguna Biasa)
	seedUser(db, entity.User{
		Name:      "Rizky Pratama (User)",
		Email:     "user@demo.com",
		Password:  "password123",
		Role:      "user",
		Phone:     "081987654321",
		Gender:    "Laki-laki",
		Address:   "Jl. Asia Afrika No. 88, Bandung",
		Bio:       "Pengguna Biasa / Regular User",
	})
}

// Helper internal untuk mengecek & memasukkan seeder user
func seedUser(db *gorm.DB, u entity.User) {
	var existingUser entity.User
	err := db.Where("email = ?", u.Email).First(&existingUser).Error

	if err != nil {
		// Jika email belum ada di database, buat akun baru
		hashedPassword, err := utils.HashPassword(u.Password)
		if err != nil {
			log.Printf("❌ Gagal meng-hash password seeder untuk %s: %v\n", u.Email, err)
			return
		}

		birthDate, _ := time.Parse("2006-01-02", "1995-08-17")

		newUser := entity.User{
			Name:      u.Name,
			Email:     u.Email,
			Password:  hashedPassword,
			Role:      u.Role,
			Phone:     u.Phone,
			Gender:    u.Gender,
			BirthDate: &birthDate,
			Address:   u.Address,
			Bio:       u.Bio,
		}

		if err := db.Create(&newUser).Error; err != nil {
			log.Printf("❌ Gagal menyisipkan seeder (%s): %v\n", u.Email, err)
			return
		}

		fmt.Printf("⚡ Database Seeder: Akun [%s] - %s berhasil dibuat! (Default Pass: password123)\n", u.Role, u.Email)
	} else {
		// Jika user sudah ada, pastikan role & data dasar tidak kosong
		if existingUser.Role != u.Role && u.Role != "" {
			existingUser.Role = u.Role
			db.Save(&existingUser)
			fmt.Printf("⚡ Database Seeder: Role akun %s diperbarui menjadi '%s'\n", u.Email, u.Role)
		}
	}
}
