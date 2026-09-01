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

	// 5. Seed 25 Akun Random Baru (Kombinasi User, Admin, Owner)
	SeedRandomUsers(db)
}

// SeedRandomUsers menyisipkan 25 akun acak dengan variasi role (admin, owner, user) ke database PostgreSQL.
func SeedRandomUsers(db *gorm.DB) {
	randomUsers := []entity.User{
		{Name: "Ahmad Fauzi", Email: "ahmad.fauzi@demo.com", Role: "user", Gender: "Laki-laki", Address: "Jl. Merdeka No. 10, Jakarta", Phone: "081234500001", Bio: "Frontend Developer & UI Designer"},
		{Name: "Dewi Lestari", Email: "dewi.lestari@demo.com", Role: "admin", Gender: "Perempuan", Address: "Jl. Diponegoro No. 25, Surabaya", Phone: "081234500002", Bio: "Operational Supervisor"},
		{Name: "Andi Wijaya", Email: "andi.wijaya@demo.com", Role: "owner", Gender: "Laki-laki", Address: "Jl. Dago No. 100, Bandung", Phone: "081234500003", Bio: "Regional Business Director"},
		{Name: "Rina Kartika", Email: "rina.kartika@demo.com", Role: "user", Gender: "Perempuan", Address: "Jl. Malioboro No. 15, Yogyakarta", Phone: "081234500004", Bio: "Content Creator & Marketer"},
		{Name: "Doni Kurniawan", Email: "doni.kurniawan@demo.com", Role: "admin", Gender: "Laki-laki", Address: "Jl. Pemuda No. 42, Semarang", Phone: "081234500005", Bio: "System & Infrastructure Admin"},
		{Name: "Eka Putri", Email: "eka.putri@demo.com", Role: "user", Gender: "Perempuan", Address: "Jl. Gajah Mada No. 8, Medan", Phone: "081234500006", Bio: "Customer Support Specialist"},
		{Name: "Fajar Hidayat", Email: "fajar.hidayat@demo.com", Role: "user", Gender: "Laki-laki", Address: "Jl. Sudirman No. 55, Palembang", Phone: "081234500007", Bio: "Backend Golang Engineer"},
		{Name: "Gita Gutawa", Email: "gita.gutawa@demo.com", Role: "owner", Gender: "Perempuan", Address: "Jl. Sunset Road No. 88, Bali", Phone: "081234500008", Bio: "Investor & Managing Partner"},
		{Name: "Hendra Setiawan", Email: "hendra.setiawan@demo.com", Role: "user", Gender: "Laki-laki", Address: "Jl. Urip Sumoharjo No. 12, Makassar", Phone: "081234500009", Bio: "Data Analyst & Quality Control"},
		{Name: "Indah Permata", Email: "indah.permata@demo.com", Role: "admin", Gender: "Perempuan", Address: "Jl. Ahmad Yani No. 30, Malang", Phone: "081234500010", Bio: "HR & Talent Acquisition Lead"},
		{Name: "Bagus Saputra", Email: "bagus.saputra@demo.com", Role: "user", Gender: "Laki-laki", Address: "Jl. Slamet Riyadi No. 9, Solo", Phone: "081234500011", Bio: "DevOps Engineer"},
		{Name: "Karin Novilda", Email: "karin.novilda@demo.com", Role: "user", Gender: "Perempuan", Address: "Jl. Kemang Raya No. 17, Jakarta Selatan", Phone: "081234500012", Bio: "Social Media Specialist"},
		{Name: "Lukman Hakim", Email: "lukman.hakim@demo.com", Role: "admin", Gender: "Laki-laki", Address: "Jl. Raden Intan No. 20, Lampung", Phone: "081234500013", Bio: "Security Auditor & Compliance"},
		{Name: "Maya Septha", Email: "maya.septha@demo.com", Role: "user", Gender: "Perempuan", Address: "Jl. Pajajaran No. 4, Bogor", Phone: "081234500014", Bio: "Accountant & Finance Lead"},
		{Name: "Naufal Samudra", Email: "naufal.samudra@demo.com", Role: "user", Gender: "Laki-laki", Address: "Jl. Pahlawan No. 77, Samarinda", Phone: "081234500015", Bio: "Fullstack Web Developer"},
		{Name: "Olga Syahputra", Email: "olga.syahputra@demo.com", Role: "owner", Gender: "Laki-laki", Address: "Jl. Imam Bonjol No. 3, Pontianak", Phone: "081234500016", Bio: "Co-Founder & Product Strategist"},
		{Name: "Putri Marino", Email: "putri.marino@demo.com", Role: "user", Gender: "Perempuan", Address: "Jl. Teuku Umar No. 22, Denpasar", Phone: "081234500017", Bio: "UX Researcher"},
		{Name: "Qori Sandioriva", Email: "qori.sandioriva@demo.com", Role: "admin", Gender: "Perempuan", Address: "Jl. Banceuy No. 11, Bandung", Phone: "081234500018", Bio: "Database Administrator"},
		{Name: "Reza Rahadian", Email: "reza.rahadian@demo.com", Role: "user", Gender: "Laki-laki", Address: "Jl. Veteran No. 5, Surakarta", Phone: "081234500019", Bio: "Product Manager"},
		{Name: "Siska Kohl", Email: "siska.kohl@demo.com", Role: "user", Gender: "Perempuan", Address: "Jl. Pantai Indah Kapuk No. 88, Jakarta Utara", Phone: "081234500020", Bio: "Creative Director"},
		{Name: "Taufik Hidayat", Email: "taufik.hidayat@demo.com", Role: "admin", Gender: "Laki-laki", Address: "Jl. Riau No. 40, Pekanbaru", Phone: "081234500021", Bio: "Operations Manager"},
		{Name: "Utari Dewi", Email: "utari.dewi@demo.com", Role: "user", Gender: "Perempuan", Address: "Jl. Somba Opu No. 14, Makassar", Phone: "081234500022", Bio: "Quality Assurance Specialist"},
		{Name: "Vidi Aldiano", Email: "vidi.aldiano@demo.com", Role: "user", Gender: "Laki-laki", Address: "Jl. Pemuda No. 101, Semarang", Phone: "081234500023", Bio: "Event & PR Executive"},
		{Name: "Wulan Guritno", Email: "wulan.guritno@demo.com", Role: "owner", Gender: "Perempuan", Address: "Jl. Radio Dalam No. 33, Jakarta Selatan", Phone: "081234500024", Bio: "Chief Executive Officer (CEO)"},
		{Name: "Yusuf Mansur", Email: "yusuf.mansur@demo.com", Role: "user", Gender: "Laki-laki", Address: "Jl. Basuki Rahmat No. 60, Surabaya", Phone: "081234500025", Bio: "Community Manager"},
	}

	for _, u := range randomUsers {
		u.Password = "password123"
		seedUser(db, u)
	}
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
		now := time.Now()

		newUser := entity.User{
			Name:            u.Name,
			Email:           u.Email,
			Password:        hashedPassword,
			Role:            u.Role,
			Phone:           u.Phone,
			Gender:          u.Gender,
			BirthDate:       &birthDate,
			Address:         u.Address,
			Bio:             u.Bio,
			IsEmailVerified: true,
			EmailVerifiedAt: &now,
		}

		if err := db.Create(&newUser).Error; err != nil {
			log.Printf("❌ Gagal menyisipkan seeder (%s): %v\n", u.Email, err)
			return
		}

		fmt.Printf("⚡ Database Seeder: Akun [%s] - %s berhasil dibuat! (Default Pass: password123)\n", u.Role, u.Email)
	} else {
		// Jika user sudah ada, pastikan role, data dasar, dan verifikasi email aktif
		updated := false
		if existingUser.Role != u.Role && u.Role != "" {
			existingUser.Role = u.Role
			updated = true
		}
		if !existingUser.IsEmailVerified {
			now := time.Now()
			existingUser.IsEmailVerified = true
			existingUser.EmailVerifiedAt = &now
			updated = true
		}
		if updated {
			db.Save(&existingUser)
			fmt.Printf("⚡ Database Seeder: Data & Role akun %s diperbarui\n", u.Email)
		}
	}
}
