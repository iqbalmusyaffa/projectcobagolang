package repository

import (
	"projectgolangnyoba/entity"

	"gorm.io/gorm"
)

// UserRepository mendefinisikan kontrak method akses data User ke database.
type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id uint) (*entity.User, error)
	FindAll() ([]entity.User, error)
	Update(user *entity.User) error
	Delete(id uint) error
	GetDashboardStats() (entity.DashboardStats, error)
}

// userRepository implementasi konkret dari UserRepository yang menggunakan GORM.
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository fungsi konstruktor untuk membuat instance UserRepository baru.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// Create menyimpan User baru ke database.
func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

// FindByEmail mencari User berdasarkan alamat email.
func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByID mencari User berdasarkan Primary Key (ID).
func (r *userRepository) FindByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindAll mengambil seluruh data pengguna dari database (diurutkan ID ascending).
func (r *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Order("id asc").Find(&users).Error
	return users, err
}

// Update memperbarui data User di database.
func (r *userRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

// Delete melakukan Soft Delete data User berdasarkan Primary Key (ID).
func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}

// GetDashboardStats menghitung statistik ringkasan total user & distribusi role dari PostgreSQL.
func (r *userRepository) GetDashboardStats() (entity.DashboardStats, error) {
	var stats entity.DashboardStats
	r.db.Model(&entity.User{}).Count(&stats.TotalUsers)
	r.db.Model(&entity.User{}).Where("role = ?", "superadmin").Count(&stats.TotalSuperadmin)
	r.db.Model(&entity.User{}).Where("role = ?", "owner").Count(&stats.TotalOwner)
	r.db.Model(&entity.User{}).Where("role = ?", "admin").Count(&stats.TotalAdmin)
	r.db.Model(&entity.User{}).Where("role = ?", "user").Count(&stats.TotalUserRole)
	return stats, nil
}
