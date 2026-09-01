package repository

import (
	"time"

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
	SaveResetToken(email, token string, expiresAt time.Time) error
	ClearResetToken(userID uint) error
	SaveVerificationToken(email, token string, expiresAt time.Time) error
	MarkEmailVerified(userID uint) error
	FindTrashedUsers() ([]entity.User, error)
	RestoreUser(userID uint) error
	PermanentDeleteUser(userID uint) error
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

// SaveResetToken menyimpan token reset password & waktu kadaluarsanya ke database.
func (r *userRepository) SaveResetToken(email, token string, expiresAt time.Time) error {
	return r.db.Model(&entity.User{}).Where("email = ?", email).Updates(map[string]interface{}{
		"reset_token":            token,
		"reset_token_expires_at": expiresAt,
	}).Error
}

// ClearResetToken menghapus token reset password setelah berhasil digunakan.
func (r *userRepository) ClearResetToken(userID uint) error {
	return r.db.Model(&entity.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"reset_token":            "",
		"reset_token_expires_at": nil,
	}).Error
}

// SaveVerificationToken menyimpan token OTP verifikasi email & masa berlakunya ke database.
func (r *userRepository) SaveVerificationToken(email, token string, expiresAt time.Time) error {
	return r.db.Model(&entity.User{}).Where("email = ?", email).Updates(map[string]interface{}{
		"verification_token":            token,
		"verification_token_expires_at": expiresAt,
	}).Error
}

// MarkEmailVerified memperbarui status verifikasi email menjadi aktif (true) dan mencatat waktu verifikasi.
func (r *userRepository) MarkEmailVerified(userID uint) error {
	now := time.Now()
	return r.db.Model(&entity.User{}).Where("id = ?", userID).Updates(map[string]interface{}{
		"is_email_verified":             true,
		"email_verified_at":             &now,
		"verification_token":            "",
		"verification_token_expires_at": nil,
	}).Error
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

// FindTrashedUsers mengambil seluruh data pengguna yang telah di-soft delete (deleted_at IS NOT NULL).
func (r *userRepository) FindTrashedUsers() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Unscoped().Where("deleted_at IS NOT NULL").Order("deleted_at desc").Find(&users).Error
	return users, err
}

// RestoreUser memulihkan pengguna terhapus dengan mengeset deleted_at menjadi NULL.
func (r *userRepository) RestoreUser(userID uint) error {
	return r.db.Unscoped().Model(&entity.User{}).Where("id = ?", userID).Update("deleted_at", nil).Error
}

// PermanentDeleteUser menghapus pengguna secara permanen (Hard Delete) dari PostgreSQL.
func (r *userRepository) PermanentDeleteUser(userID uint) error {
	return r.db.Unscoped().Delete(&entity.User{}, userID).Error
}
