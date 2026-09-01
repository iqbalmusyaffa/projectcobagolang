package repository

import (
	"time"

	"projectgolangnyoba/entity"

	"gorm.io/gorm"
)

// SessionRepository mendefinisikan operasi basis data untuk mengelola sesi pengguna dan perangkat aktif.
type SessionRepository interface {
	Create(session *entity.UserSession) error
	FindByUserID(userID uint) ([]entity.UserSession, error)
	FindByID(id uint) (*entity.UserSession, error)
	FindByRefreshTokenHash(hash string) (*entity.UserSession, error)
	UpdateLastActive(sessionID uint) error
	Delete(id uint) error
	DeleteByUserIDExcept(userID uint, exceptSessionID uint) error
	DeleteByUserID(userID uint) error
	DeleteByRefreshTokenHash(hash string) error
}

type sessionRepository struct {
	db *gorm.DB
}

// NewSessionRepository membuat instance baru dari SessionRepository.
func NewSessionRepository(db *gorm.DB) SessionRepository {
	return &sessionRepository{db: db}
}

// Create menyimpan data sesi pengguna baru ke PostgreSQL.
func (r *sessionRepository) Create(session *entity.UserSession) error {
	return r.db.Create(session).Error
}

// FindByUserID mengambil seluruh sesi aktif milik seorang pengguna diurutkan dari yang paling baru aktif.
func (r *sessionRepository) FindByUserID(userID uint) ([]entity.UserSession, error) {
	var sessions []entity.UserSession
	err := r.db.Where("user_id = ?", userID).Order("last_active_at desc").Find(&sessions).Error
	return sessions, err
}

// FindByID mencari satu data sesi berdasarkan ID Primary Key.
func (r *sessionRepository) FindByID(id uint) (*entity.UserSession, error) {
	var session entity.UserSession
	err := r.db.First(&session, id).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

// FindByRefreshTokenHash mencari data sesi berdasarkan hash refresh token.
func (r *sessionRepository) FindByRefreshTokenHash(hash string) (*entity.UserSession, error) {
	var session entity.UserSession
	err := r.db.Where("refresh_token_hash = ?", hash).First(&session).Error
	if err != nil {
		return nil, err
	}
	return &session, nil
}

// UpdateLastActive memperbarui waktu aktivitas terakhir (last_active_at) pada sebuah sesi.
func (r *sessionRepository) UpdateLastActive(sessionID uint) error {
	return r.db.Model(&entity.UserSession{}).Where("id = ?", sessionID).Update("last_active_at", time.Now()).Error
}

// Delete menghapus satu sesi (Logout dari perangkat tertentu).
func (r *sessionRepository) Delete(id uint) error {
	return r.db.Delete(&entity.UserSession{}, id).Error
}

// DeleteByUserIDExcept menghapus semua sesi aktif milik user KECUALI sesi yang sedang digunakan saat ini.
func (r *sessionRepository) DeleteByUserIDExcept(userID uint, exceptSessionID uint) error {
	return r.db.Where("user_id = ? AND id != ?", userID, exceptSessionID).Delete(&entity.UserSession{}).Error
}

// DeleteByUserID menghapus semua sesi aktif milik user tertentu.
func (r *sessionRepository) DeleteByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&entity.UserSession{}).Error
}

// DeleteByRefreshTokenHash menghapus sesi berdasarkan hash refresh token saat logout.
func (r *sessionRepository) DeleteByRefreshTokenHash(hash string) error {
	return r.db.Where("refresh_token_hash = ?", hash).Delete(&entity.UserSession{}).Error
}
