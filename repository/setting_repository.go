package repository

import (
	"projectgolangnyoba/entity"

	"gorm.io/gorm"
)

// SettingRepository mendefinisikan kontrak interface akses data untuk pengaturan sistem.
type SettingRepository interface {
	GetSettings() (*entity.SystemSetting, error)
	UpdateSettings(setting *entity.SystemSetting) error
	UpdateLogo(logoPath string) error
}

type settingRepository struct {
	db *gorm.DB
}

// NewSettingRepository adalah konstruktor untuk membuat instance SettingRepository.
func NewSettingRepository(db *gorm.DB) SettingRepository {
	return &settingRepository{db: db}
}

// GetSettings mengambil baris pertama pengaturan sistem (atau membuat nilai default jika belum ada).
func (r *settingRepository) GetSettings() (*entity.SystemSetting, error) {
	var setting entity.SystemSetting
	err := r.db.First(&setting).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Jika belum ada, buat baris pengaturan default pertama
			defaultSetting := entity.SystemSetting{
				AppName:                 "TailAdmin Vue 3",
				AppDescription:         "Dashboard REST API Golang Clean Architecture",
				DefaultLanguage:        "id",
				Timezone:               "Asia/Jakarta",
				SMTPPort:               "587",
				EnableEmailNotification: true,
			}
			if createErr := r.db.Create(&defaultSetting).Error; createErr != nil {
				return nil, createErr
			}
			return &defaultSetting, nil
		}
		return nil, err
	}
	return &setting, nil
}

// UpdateSettings memperbarui data pengaturan sistem di PostgreSQL.
func (r *settingRepository) UpdateSettings(setting *entity.SystemSetting) error {
	return r.db.Save(setting).Error
}

// UpdateLogo memperbarui path logo aplikasi di PostgreSQL.
func (r *settingRepository) UpdateLogo(logoPath string) error {
	setting, err := r.GetSettings()
	if err != nil {
		return err
	}
	return r.db.Model(setting).Update("app_logo", logoPath).Error
}
