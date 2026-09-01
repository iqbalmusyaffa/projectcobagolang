package entity

import "time"

// UserSession merepresentasikan tabel 'user_sessions' di PostgreSQL untuk mencatat sesi aktif pengguna.
type UserSession struct {
	ID               uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID           uint      `gorm:"not null;index" json:"user_id"`
	RefreshTokenHash string    `gorm:"type:varchar(255);index;not null" json:"-"`
	IPAddress        string    `gorm:"type:varchar(100)" json:"ip_address"`
	UserAgent        string    `gorm:"type:text" json:"user_agent"`
	DeviceName       string    `gorm:"type:varchar(100)" json:"device_name"`
	DeviceType       string    `gorm:"type:varchar(50)" json:"device_type"` // desktop, mobile, tablet
	Browser          string    `gorm:"type:varchar(100)" json:"browser"`
	OS               string    `gorm:"type:varchar(100)" json:"os"`
	LastActiveAt     time.Time `json:"last_active_at"`
	ExpiresAt        time.Time `json:"expires_at"`
	CreatedAt        time.Time `gorm:"autoCreateTime" json:"created_at"`
}

// UserSessionResponse merepresentasikan format JSON informasi sesi yang dikembalikan ke client.
type UserSessionResponse struct {
	ID               uint   `json:"id"`
	UserID           uint   `json:"user_id"`
	IPAddress        string `json:"ip_address"`
	UserAgent        string `json:"user_agent"`
	DeviceName       string `json:"device_name"`
	DeviceType       string `json:"device_type"`
	Browser          string `json:"browser"`
	OS               string `json:"os"`
	IsCurrentSession bool   `json:"is_current_session"`
	LastActiveAt     string `json:"last_active_at"`
	CreatedAt        string `json:"created_at"`
}

// FormatUserSession mengonversi entitas UserSession menjadi UserSessionResponse.
func FormatUserSession(session UserSession, isCurrent bool) UserSessionResponse {
	return UserSessionResponse{
		ID:               session.ID,
		UserID:           session.UserID,
		IPAddress:        session.IPAddress,
		UserAgent:        session.UserAgent,
		DeviceName:       session.DeviceName,
		DeviceType:       session.DeviceType,
		Browser:          session.Browser,
		OS:               session.OS,
		IsCurrentSession: isCurrent,
		LastActiveAt:     session.LastActiveAt.Format(time.RFC3339),
		CreatedAt:        session.CreatedAt.Format(time.RFC3339),
	}
}
