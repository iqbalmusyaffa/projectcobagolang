package config

import (
	"fmt"
	"log"

	"projectgolangnyoba/entity"
	"projectgolangnyoba/utils"

	"gorm.io/gorm"
)

// SeedSuperAdmin secara otomatis menyisipkan akun Super Admin 'faridwimansyah8@gmail.com' ke PostgreSQL jika belum ada.
func SeedSuperAdmin(db *gorm.DB) {
	superAdminEmail := "faridwimansyah8@gmail.com"

	var existingUser entity.User
	err := db.Where("email = ?", superAdminEmail).First(&existingUser).Error

	if err != nil {
		// jika user belum ada di database, buat akun Super Admin baru
		hashedPassword, err := utils.HashPassword("admin123456")
		if err != nil {
			log.Printf("❌ Gagal meng-hash password seeder: %v\n", err)
			return
		}

		superAdmin := entity.User{
			Name:     "Farid Wiansyah (Super Admin)",
			Email:    superAdminEmail,
			Password: hashedPassword,
			Role:     "superadmin",
		}

		if err := db.Create(&superAdmin).Error; err != nil {
			log.Printf("❌ Gagal menyisipkan seeder Super Admin: %v\n", err)
			return
		}

		fmt.Printf("⚡ Database Seeder: Akun Super Admin (%s) berhasil dibuat! (Default Password: admin123456)\n", superAdminEmail)
	} else {
		// Jika user sudah ada, pastikan rolenya tetap superadmin
		if existingUser.Role != "superadmin" {
			existingUser.Role = "superadmin"
			db.Save(&existingUser)
			fmt.Printf("⚡ Database Seeder: Role untuk %s berhasil diperbarui menjadi 'superadmin'!\n", superAdminEmail)
		}
	}
}
