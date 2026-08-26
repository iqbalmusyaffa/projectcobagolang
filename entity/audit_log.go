package entity

import "time"

// AuditLog merepresentasikan entitas catatan aktivitas pengguna di database.
type AuditLog struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	UserName  string    `gorm:"type:varchar(100)" json:"user_name"`
	UserRole  string    `gorm:"type:varchar(50)" json:"user_role"`
	Action    string    `gorm:"type:varchar(255);not null" json:"action"`
	IPAddress string    `gorm:"type:varchar(100)" json:"ip_address"`
	UserAgent string    `gorm:"type:text" json:"user_agent"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// AuditLogResponse merepresentasikan format JSON response catatan aktivitas.
type AuditLogResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	UserName  string `json:"user_name"`
	UserRole  string `json:"user_role"`
	Action    string `json:"action"`
	IPAddress string `json:"ip_address"`
	UserAgent string `json:"user_agent"`
	CreatedAt string `json:"created_at"`
}

// FormatAuditLog mengonversi entity AuditLog ke DTO AuditLogResponse.
func FormatAuditLog(log AuditLog) AuditLogResponse {
	userName := log.UserName
	if userName == "" {
		userName = "Pengguna"
	}
	userRole := log.UserRole
	if userRole == "" {
		userRole = "user"
	}

	return AuditLogResponse{
		ID:        log.ID,
		UserID:    log.UserID,
		UserName:  userName,
		UserRole:  userRole,
		Action:    log.Action,
		IPAddress: log.IPAddress,
		UserAgent: log.UserAgent,
		CreatedAt: log.CreatedAt.Format(time.RFC3339),
	}
}
