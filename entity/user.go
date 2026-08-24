package entity

import (
	"time"

	"gorm.io/gorm"
)

// User melambangkan struktur tabel 'users' di database PostgreSQL.
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"type:varchar(100);not null" json:"name"`
	Email     string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password  string         `gorm:"type:varchar(255);not null" json:"-"` // "-" artinya password tidak dimasukkan ke respon JSON
	Role      string         `gorm:"type:varchar(20);not null;default:'admin'" json:"role"` // superadmin, owner, admin
	Avatar    string         `gorm:"type:varchar(255)" json:"avatar"`
	Phone     string         `gorm:"type:varchar(20)" json:"phone"`
	Gender    string         `gorm:"type:varchar(10)" json:"gender"` // Laki-laki / Perempuan
	BirthDate *time.Time     `gorm:"type:date" json:"birth_date"`
	Address   string         `gorm:"type:text" json:"address"`
	Bio       string         `gorm:"type:text" json:"bio"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// RegisterInput melambangkan data JSON yang dikirim client saat Registrasi.
type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role"` // Opsional, default 'admin'
}

// LoginInput melambangkan data JSON yang dikirim client saat Login.
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// UserResponse melambangkan data JSON kembalian ke client setelah berhasil Auth.
type UserResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Avatar    string    `json:"avatar"`
	Phone     string    `json:"phone"`
	Gender    string    `json:"gender"`
	BirthDate string    `json:"birth_date"`
	Address   string    `json:"address"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// UpdateProfileInput melambangkan data JSON saat mengedit profil.
type UpdateProfileInput struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
	Address   string `json:"address"`
	Bio       string `json:"bio"`
}

// ChangePasswordInput melambangkan data JSON saat mengganti password.
type ChangePasswordInput struct {
	OldPassword string `json:"old_password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

// ChangeRoleInput melambangkan data JSON saat Superadmin mengubah role pengguna lain.
type ChangeRoleInput struct {
	UserID uint   `json:"user_id" binding:"required"`
	Role   string `json:"role" binding:"required"` // superadmin, owner, admin
}

// AdminCreateUserInput melambangkan data JSON saat Super Admin menambahkan pengguna baru.
type AdminCreateUserInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required"` // superadmin, owner, admin
}

// AdminUpdateUserInput melambangkan data JSON saat Super Admin memperbarui data pengguna lain.
type AdminUpdateUserInput struct {
	UserID    uint   `json:"user_id" binding:"required"`
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Role      string `json:"role" binding:"required"` // superadmin, owner, admin
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
	BirthDate string `json:"birth_date"`
	Address   string `json:"address"`
	Bio       string `json:"bio"`
}

// RefreshTokenInput melambangkan data JSON saat meminta Access Token baru.
type RefreshTokenInput struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// TokenPairResponse melambangkan pasangan Access Token & Refresh Token.
type TokenPairResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}


// FormatUser mengubah struct User menjadi UserResponse (menghilangkan password)
func FormatUser(user User) UserResponse {
	role := user.Role
	if role == "" {
		role = "admin"
	}
	birthDateStr := ""
	if user.BirthDate != nil {
		birthDateStr = user.BirthDate.Format("2006-01-02")
	}
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      role,
		Avatar:    user.Avatar,
		Phone:     user.Phone,
		Gender:    user.Gender,
		BirthDate: birthDateStr,
		Address:   user.Address,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
