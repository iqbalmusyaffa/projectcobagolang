package repository

import (
	"projectgolangnyoba/entity"

	"gorm.io/gorm"
)

// AuditLogRepository mendefinisikan kontrak interface operasi database untuk AuditLog.
type AuditLogRepository interface {
	Create(log *entity.AuditLog) error
	FindByUserID(userID uint) ([]entity.AuditLog, error)
	FindAll() ([]entity.AuditLog, error)
}

type auditLogRepository struct {
	db *gorm.DB
}

// NewAuditLogRepository adalah konstruktor untuk menginisialisasi repository audit log.
func NewAuditLogRepository(db *gorm.DB) AuditLogRepository {
	return &auditLogRepository{db: db}
}

// Create menyimpan entitas AuditLog baru ke PostgreSQL.
func (r *auditLogRepository) Create(log *entity.AuditLog) error {
	return r.db.Create(log).Error
}

// FindByUserID mengambil riwayat aktivitas milik satu pengguna tertentu (diurutkan terbaru).
func (r *auditLogRepository) FindByUserID(userID uint) ([]entity.AuditLog, error) {
	var logs []entity.AuditLog
	err := r.db.Where("user_id = ?", userID).Order("id desc").Limit(100).Find(&logs).Error
	return logs, err
}

// FindAll mengambil seluruh catatan aktivitas sistem (diurutkan terbaru).
func (r *auditLogRepository) FindAll() ([]entity.AuditLog, error) {
	var logs []entity.AuditLog
	err := r.db.Order("id desc").Limit(200).Find(&logs).Error
	return logs, err
}
