package usecase

import (
	"errors"
	"time"

	"projectgolangnyoba/entity"
	"projectgolangnyoba/repository"
	"projectgolangnyoba/utils"
)

// UserUsecase mendefinisikan kontrak logika bisnis terkait User.
type UserUsecase interface {
	Register(input entity.RegisterInput) (entity.UserResponse, error)
	Login(input entity.LoginInput) (entity.TokenPairResponse, error)
	RefreshToken(refreshTokenStr string) (string, error)
	GetProfile(userID uint) (entity.UserResponse, error)
	UpdateProfile(userID uint, input entity.UpdateProfileInput) (entity.UserResponse, error)
	ChangePassword(userID uint, input entity.ChangePasswordInput) error
	GetAllUsers() ([]entity.UserResponse, error)
	ChangeUserRole(input entity.ChangeRoleInput) (entity.UserResponse, error)
	LogoutToken(token string) error
}

// userUsecase implementasi dari UserUsecase yang bergantung pada UserRepository & RedisRepository.
type userUsecase struct {
	userRepo  repository.UserRepository
	redisRepo repository.RedisRepository
}

// NewUserUsecase adalah konstruktor untuk membuat instance UserUsecase.
func NewUserUsecase(userRepo repository.UserRepository, redisRepo repository.RedisRepository) UserUsecase {
	return &userUsecase{
		userRepo:  userRepo,
		redisRepo: redisRepo,
	}
}

// Register memproses pendaftaran akun pengguna baru.
func (u *userUsecase) Register(input entity.RegisterInput) (entity.UserResponse, error) {
	// 1. Cek apakah email sudah terdaftar
	existingUser, _ := u.userRepo.FindByEmail(input.Email)
	if existingUser != nil {
		return entity.UserResponse{}, errors.New("email sudah terdaftar, silakan gunakan email lain")
	}

	// 2. Hash password menggunakan bcrypt
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return entity.UserResponse{}, errors.New("gagal mengamankan password")
	}

	// 3. Registrasi publik SELALU dipaksa menjadi role 'admin' demi keamanan
	user := entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Role:     "admin",
	}

	// 5. Simpan ke database melalui repository
	err = u.userRepo.Create(&user)
	if err != nil {
		return entity.UserResponse{}, err
	}

	// 6. Kembalikan data user dalam format UserResponse (tanpa password)
	return entity.FormatUser(user), nil
}

// Login memverifikasi kredensial email & password dan menghasilkan Access Token (15m) & Refresh Token (7 Hari).
func (u *userUsecase) Login(input entity.LoginInput) (entity.TokenPairResponse, error) {
	// 1. Cari user berdasarkan email
	user, err := u.userRepo.FindByEmail(input.Email)
	if err != nil {
		return entity.TokenPairResponse{}, errors.New("email atau password salah")
	}

	// 2. Verifikasi kesesuaian password
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return entity.TokenPairResponse{}, errors.New("email atau password salah")
	}

	role := user.Role
	if role == "" {
		role = "admin"
	}

	// 3. Generate Access Token (15 Menit) & Refresh Token (7 Hari)
	accessToken, err := utils.GenerateAccessToken(user.ID, user.Email, role)
	if err != nil {
		return entity.TokenPairResponse{}, errors.New("gagal membuat access token")
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Email)
	if err != nil {
		return entity.TokenPairResponse{}, errors.New("gagal membuat refresh token")
	}

	// 4. Simpan Refresh Token ke Redis dengan TTL 7 Hari (168 Jam)
	if u.redisRepo != nil {
		_ = u.redisRepo.StoreRefreshToken(user.ID, refreshToken, 7*24*time.Hour)
	}

	return entity.TokenPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// RefreshToken memvalidasi Refresh Token dan mengembalikan Access Token 15 menit baru.
func (u *userUsecase) RefreshToken(refreshTokenStr string) (string, error) {
	// 1. Validasi signature & struktur Refresh Token
	claims, err := utils.ValidateToken(refreshTokenStr)
	if err != nil {
		return "", errors.New("refresh token tidak valid atau sudah kadaluwarsa")
	}

	if claims.TokenType != "refresh" {
		return "", errors.New("token bukan berjenis refresh token")
	}

	// 2. Cek apakah Refresh Token masih terdaftar dan cocok di Redis
	if u.redisRepo != nil {
		storedToken, err := u.redisRepo.GetRefreshToken(claims.UserID)
		if err != nil || storedToken != refreshTokenStr {
			return "", errors.New("refresh token telah dicabut atau di-logout")
		}
	}

	// 3. Ambil data user dari database untuk mendapatkan Role terbaru
	user, err := u.userRepo.FindByID(claims.UserID)
	if err != nil {
		return "", errors.New("pengguna tidak ditemukan")
	}

	role := user.Role
	if role == "" {
		role = "admin"
	}

	// 4. Terbitkan Access Token 15 Menit yang baru
	newAccessToken, err := utils.GenerateAccessToken(user.ID, user.Email, role)
	if err != nil {
		return "", errors.New("gagal membuat access token baru")
	}

	return newAccessToken, nil
}

// GetProfile mengambil profil pengguna berdasarkan ID dari JWT Token.
func (u *userUsecase) GetProfile(userID uint) (entity.UserResponse, error) {
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return entity.UserResponse{}, errors.New("user tidak ditemukan")
	}
	return entity.FormatUser(*user), nil
}

// UpdateProfile memperbarui Nama & Email pengguna.
func (u *userUsecase) UpdateProfile(userID uint, input entity.UpdateProfileInput) (entity.UserResponse, error) {
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return entity.UserResponse{}, errors.New("user tidak ditemukan")
	}

	// Jika email diubah, pastikan email baru belum dipakai user lain
	if input.Email != user.Email {
		existingUser, _ := u.userRepo.FindByEmail(input.Email)
		if existingUser != nil && existingUser.ID != userID {
			return entity.UserResponse{}, errors.New("email sudah digunakan oleh pengguna lain")
		}
	}

	user.Name = input.Name
	user.Email = input.Email

	err = u.userRepo.Update(user)
	if err != nil {
		return entity.UserResponse{}, errors.New("gagal memperbarui profil")
	}

	return entity.FormatUser(*user), nil
}

// ChangePassword mengganti password pengguna setelah memverifikasi password lama.
func (u *userUsecase) ChangePassword(userID uint, input entity.ChangePasswordInput) error {
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}

	// 1. Verifikasi password lama
	if !utils.CheckPasswordHash(input.OldPassword, user.Password) {
		return errors.New("password lama tidak sesuai")
	}

	// 2. Hash password baru
	newHashedPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		return errors.New("gagal mengamankan password baru")
	}

	user.Password = newHashedPassword

	// 3. Simpan perubahan password
	return u.userRepo.Update(user)
}

// GetAllUsers mengambil seluruh daftar pengguna (Khusus Superadmin & Owner).
func (u *userUsecase) GetAllUsers() ([]entity.UserResponse, error) {
	users, err := u.userRepo.FindAll()
	if err != nil {
		return nil, errors.New("gagal mengambil daftar pengguna")
	}

	var formattedUsers []entity.UserResponse
	for _, user := range users {
		formattedUsers = append(formattedUsers, entity.FormatUser(user))
	}

	return formattedUsers, nil
}

// ChangeUserRole mengubah peran (role) pengguna lain (Khusus Superadmin).
func (u *userUsecase) ChangeUserRole(input entity.ChangeRoleInput) (entity.UserResponse, error) {
	// Validasi role target
	if input.Role != "superadmin" && input.Role != "owner" && input.Role != "admin" {
		return entity.UserResponse{}, errors.New("role tidak valid (pilih: superadmin, owner, atau admin)")
	}

	user, err := u.userRepo.FindByID(input.UserID)
	if err != nil {
		return entity.UserResponse{}, errors.New("pengguna target tidak ditemukan")
	}

	user.Role = input.Role

	err = u.userRepo.Update(user)
	if err != nil {
		return entity.UserResponse{}, errors.New("gagal memperbarui peran pengguna")
	}

	return entity.FormatUser(*user), nil
}

// LogoutToken memasukkan token JWT ke Redis Blacklist & menghapus Refresh Token (Instant Logout).
func (u *userUsecase) LogoutToken(token string) error {
	if u.redisRepo != nil && token != "" {
		claims, err := utils.ValidateToken(token)
		if err == nil && claims.UserID != 0 {
			// Hapus Refresh Token dari Redis
			_ = u.redisRepo.DeleteRefreshToken(claims.UserID)
		}
		// Simpan Access Token ke Redis Blacklist dengan TTL 24 jam
		return u.redisRepo.BlacklistToken(token, 24*time.Hour)
	}
	return nil
}
