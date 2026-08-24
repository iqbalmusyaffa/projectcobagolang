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
