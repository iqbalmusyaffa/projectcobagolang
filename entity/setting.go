package entity

import "time"

// SystemSetting merepresentasikan tabel pengaturan sistem di database PostgreSQL.
type SystemSetting struct {
	ID                      uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	AppName                 string    `gorm:"type:varchar(100);default:'TailAdmin Vue 3'" json:"app_name"`
	AppDescription         string    `gorm:"type:varchar(255);default:'Dashboard REST API Golang Clean Architecture'" json:"app_description"`
	AppLogo                string    `gorm:"type:varchar(255);default:''" json:"app_logo"`
	DefaultLanguage        string    `gorm:"type:varchar(20);default:'id'" json:"default_language"`
	Timezone               string    `gorm:"type:varchar(50);default:'Asia/Jakarta'" json:"timezone"`
	SMTPHost               string    `gorm:"type:varchar(100);default:''" json:"smtp_host"`
	SMTPPort               string    `gorm:"type:varchar(10);default:'587'" json:"smtp_port"`
	SMTPSenderEmail        string    `gorm:"type:varchar(100);default:''" json:"smtp_sender_email"`
	SMTPSenderPassword     string    `gorm:"type:varchar(255);default:''" json:"smtp_sender_password"`
	SMTPUser               string    `gorm:"type:varchar(100);default:''" json:"smtp_user"`
	EnableEmailNotification bool      `gorm:"default:true" json:"enable_email_notification"`
	UpdatedAt               time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// SystemSettingInput DTO untuk menerima input pembaruan pengaturan dari frontend.
type SystemSettingInput struct {
	AppName                 string `json:"app_name"`
	AppDescription         string `json:"app_description"`
	DefaultLanguage        string `json:"default_language"`
	Timezone               string `json:"timezone"`
	SMTPHost               string `json:"smtp_host"`
	SMTPPort               string `json:"smtp_port"`
	SMTPSenderEmail        string `json:"smtp_sender_email"`
	SMTPSenderPassword     string `json:"smtp_sender_password"`
	SMTPUser               string `json:"smtp_user"`
	EnableEmailNotification bool   `json:"enable_email_notification"`
}

// TestEmailInput DTO untuk menerima email tujuan uji coba pengiriman email.
type TestEmailInput struct {
	TargetEmail string `json:"target_email" binding:"required,email"`
}
