package entity

import "time"

// User melambangkan struktur tabel 'users' di database PostgreSQL.
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100);not null" json:"name"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"-"` // "-" artinya password tidak dimasukkan ke respon JSON
	Role      string    `gorm:"type:varchar(20);not null;default:'admin'" json:"role"` // superadmin, owner, admin
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
	CreatedAt time.Time `json:"created_at"`
}

// UpdateProfileInput melambangkan data JSON saat mengedit profil.
type UpdateProfileInput struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
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
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      role,
		CreatedAt: user.CreatedAt,
	}
}
