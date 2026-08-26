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
	AdminCreateUser(input entity.AdminCreateUserInput) (entity.UserResponse, error)
	AdminUpdateUser(input entity.AdminUpdateUserInput) (entity.UserResponse, error)
	AdminDeleteUser(targetUserID uint) error
	Login(input entity.LoginInput) (entity.TokenPairResponse, error)
	RefreshToken(refreshTokenStr string) (string, error)
	GetProfile(userID uint) (entity.UserResponse, error)
	UpdateProfile(userID uint, input entity.UpdateProfileInput) (entity.UserResponse, error)
	UploadAvatar(userID uint, avatarPath string) (entity.UserResponse, error)
	ChangePassword(userID uint, input entity.ChangePasswordInput) error
	GetAllUsers() ([]entity.UserResponse, error)
	ChangeUserRole(input entity.ChangeRoleInput) (entity.UserResponse, error)
	LogoutToken(token string) error
	DeleteAccount(userID uint, token string) error
	GetDashboardStats() (entity.DashboardStats, error)
	RequestPasswordReset(email string) error
	ResetPassword(email, otp, newPassword string) error
	AdminResetUserPassword(input entity.AdminResetPasswordInput) error
	LogActivity(userID uint, userName, userRole, action, ip, userAgent string)
	GetMyAuditLogs(userID uint) ([]entity.AuditLogResponse, error)
	GetAllAuditLogs() ([]entity.AuditLogResponse, error)
	GetTrashedUsers() ([]entity.UserResponse, error)
	RestoreUser(targetUserID uint) error
	PermanentDeleteUser(targetUserID uint) error
}

// userUsecase implementasi dari UserUsecase yang bergantung pada UserRepository, RedisRepository, & AuditLogRepository.
type userUsecase struct {
	userRepo  repository.UserRepository
	redisRepo repository.RedisRepository
	auditRepo repository.AuditLogRepository
}

// NewUserUsecase adalah konstruktor untuk membuat instance UserUsecase.
func NewUserUsecase(userRepo repository.UserRepository, redisRepo repository.RedisRepository, auditRepo repository.AuditLogRepository) UserUsecase {
	return &userUsecase{
		userRepo:  userRepo,
		redisRepo: redisRepo,
		auditRepo: auditRepo,
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

// AdminCreateUser memproses pendaftaran pengguna baru khusus oleh Super Admin dengan menentukan role.
func (u *userUsecase) AdminCreateUser(input entity.AdminCreateUserInput) (entity.UserResponse, error) {
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

	// 3. Validasi role (superadmin, owner, admin, user; default: user)
	role := input.Role
	if role != "superadmin" && role != "owner" && role != "admin" && role != "user" {
		role = "user"
	}

	// 4. Buat entity User baru
	user := entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Role:     role,
	}

	// 5. Simpan ke database melalui repository
	err = u.userRepo.Create(&user)
	if err != nil {
		return entity.UserResponse{}, err
	}

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

// UpdateProfile memperbarui profil pengguna (Nama, Email, HP, Gender, Tgl Lahir, Alamat, Bio).
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
	user.Phone = input.Phone
	user.Gender = input.Gender
	user.Address = input.Address
	user.Bio = input.Bio

	if input.BirthDate != "" {
		parsedDate, err := time.Parse("2006-01-02", input.BirthDate)
		if err == nil {
			user.BirthDate = &parsedDate
		}
	} else {
		user.BirthDate = nil
	}

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
	err = u.userRepo.Update(user)
	if err != nil {
		return err
	}

	// 4. Kirim email konfirmasi bahwa password berhasil diubah (asynchronous goroutine)
	go utils.SendPasswordChangedSuccessEmail(user.Email, user.Name)

	return nil
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

// UploadAvatar meng-update path foto profil user di database.
func (u *userUsecase) UploadAvatar(userID uint, avatarPath string) (entity.UserResponse, error) {
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return entity.UserResponse{}, errors.New("user tidak ditemukan")
	}

	user.Avatar = avatarPath
	err = u.userRepo.Update(user)
	if err != nil {
		return entity.UserResponse{}, errors.New("gagal menyimpan avatar di database")
	}

	return entity.FormatUser(*user), nil
}

// DeleteAccount menghapus akun secara Soft Delete dan meng-invalidasi token di Redis.
func (u *userUsecase) DeleteAccount(userID uint, token string) error {
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}

	err = u.userRepo.Delete(user.ID)
	if err != nil {
		return errors.New("gagal menghapus akun")
	}

	if u.redisRepo != nil {
		_ = u.redisRepo.DeleteRefreshToken(userID)
		if token != "" {
			_ = u.redisRepo.BlacklistToken(token, 24*time.Hour)
		}
	}

	return nil
}

// AdminUpdateUser memproses pembaruan data pengguna oleh Super Admin.
func (u *userUsecase) AdminUpdateUser(input entity.AdminUpdateUserInput) (entity.UserResponse, error) {
	user, err := u.userRepo.FindByID(input.UserID)
	if err != nil {
		return entity.UserResponse{}, errors.New("pengguna target tidak ditemukan")
	}

	// Cek unik email jika email diubah
	if input.Email != user.Email {
		existingUser, _ := u.userRepo.FindByEmail(input.Email)
		if existingUser != nil && existingUser.ID != input.UserID {
			return entity.UserResponse{}, errors.New("email sudah digunakan oleh pengguna lain")
		}
	}

	role := input.Role
	if role != "superadmin" && role != "owner" && role != "admin" && role != "user" {
		role = "user"
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Role = role
	user.Phone = input.Phone
	user.Gender = input.Gender
	user.Address = input.Address
	user.Bio = input.Bio

	if input.BirthDate != "" {
		parsedDate, err := time.Parse("2006-01-02", input.BirthDate)
		if err == nil {
			user.BirthDate = &parsedDate
		}
	} else {
		user.BirthDate = nil
	}

	err = u.userRepo.Update(user)
	if err != nil {
		return entity.UserResponse{}, errors.New("gagal memperbarui data pengguna")
	}

	return entity.FormatUser(*user), nil
}

// AdminDeleteUser menghapus pengguna lain oleh Super Admin (Soft Delete).
func (u *userUsecase) AdminDeleteUser(targetUserID uint) error {
	user, err := u.userRepo.FindByID(targetUserID)
	if err != nil {
		return errors.New("pengguna target tidak ditemukan")
	}

	err = u.userRepo.Delete(user.ID)
	if err != nil {
		return errors.New("gagal menghapus pengguna")
	}

	if u.redisRepo != nil {
		_ = u.redisRepo.DeleteRefreshToken(targetUserID)
	}

	return nil
}

// GetDashboardStats mengambil data statistik ringkasan total pengguna & role.
func (u *userUsecase) GetDashboardStats() (entity.DashboardStats, error) {
	return u.userRepo.GetDashboardStats()
}

// RequestPasswordReset memproses permintaan OTP untuk reset password via email.
func (u *userUsecase) RequestPasswordReset(email string) error {
	// 1. Cari user berdasarkan email
	user, err := u.userRepo.FindByEmail(email)
	if err != nil || user == nil {
		return errors.New("email tidak terdaftar di sistem")
	}

	// 2. Generate kode OTP 6 digit
	otpCode, err := utils.GenerateOTP()
	if err != nil {
		return errors.New("gagal menghasilkan kode OTP")
	}

	expiresAt := time.Now().Add(15 * time.Minute)

	// 3. Simpan OTP di Redis jika Redis tersedia (TTL 15 Menit)
	if u.redisRepo != nil {
		_ = u.redisRepo.SetCache("reset_code:"+email, otpCode, 15*time.Minute)
	}

	// 4. Simpan ke database sebagai fallback
	_ = u.userRepo.SaveResetToken(email, otpCode, expiresAt)

	// 5. Kirim email berisi OTP
	return utils.SendResetPasswordEmail(email, otpCode)
}

// ResetPassword memverifikasi OTP dan mengganti password pengguna dengan yang baru.
func (u *userUsecase) ResetPassword(email, otp, newPassword string) error {
	// 1. Cari user berdasarkan email
	user, err := u.userRepo.FindByEmail(email)
	if err != nil || user == nil {
		return errors.New("email tidak terdaftar di sistem")
	}

	// 2. Verifikasi OTP dari Redis terlebih dahulu jika ada
	var isValidOTP bool
	if u.redisRepo != nil {
		cachedOTP, err := u.redisRepo.GetCache("reset_code:" + email)
		if err == nil && cachedOTP == otp {
			isValidOTP = true
		}
	}

	// Fallback ke database jika Redis tidak merespon/offline
	if !isValidOTP {
		if user.ResetToken != "" && user.ResetToken == otp {
			if user.ResetTokenExpiresAt != nil && user.ResetTokenExpiresAt.After(time.Now()) {
				isValidOTP = true
			}
		}
	}

	if !isValidOTP {
		return errors.New("kode OTP tidak valid atau sudah kadaluwarsa")
	}

	// 3. Hash password baru
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New("gagal mengamankan password baru")
	}

	user.Password = hashedPassword

	// 4. Simpan perubahan ke database
	err = u.userRepo.Update(user)
	if err != nil {
		return errors.New("gagal memperbarui password")
	}

	// 5. Hapus OTP dari Redis & Database
	if u.redisRepo != nil {
		_ = u.redisRepo.DeleteCache("reset_code:" + email)
	}
	_ = u.userRepo.ClearResetToken(user.ID)

	// 6. Kirim email konfirmasi bahwa password berhasil direset (asynchronous goroutine)
	go utils.SendPasswordChangedSuccessEmail(user.Email, user.Name)

	return nil
}

// AdminResetUserPassword mereset password pengguna secara langsung oleh Super Admin tanpa OTP.
func (u *userUsecase) AdminResetUserPassword(input entity.AdminResetPasswordInput) error {
	user, err := u.userRepo.FindByID(input.UserID)
	if err != nil || user == nil {
		return errors.New("pengguna target tidak ditemukan")
	}

	hashedPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		return errors.New("gagal mengamankan password baru")
	}

	user.Password = hashedPassword

	err = u.userRepo.Update(user)
	if err != nil {
		return errors.New("gagal memperbarui password pengguna")
	}

	if u.redisRepo != nil {
		_ = u.redisRepo.DeleteRefreshToken(input.UserID)
	}

	// Kirim email konfirmasi bahwa password berhasil direset oleh Superadmin
	go utils.SendPasswordChangedSuccessEmail(user.Email, user.Name)

	return nil
}

// LogActivity mencatat aktivitas pengguna secara mandiri ke tabel AuditLog.
func (u *userUsecase) LogActivity(userID uint, userName, userRole, action, ip, userAgent string) {
	if u.auditRepo == nil {
		return
	}
	log := entity.AuditLog{
		UserID:    userID,
		UserName:  userName,
		UserRole:  userRole,
		Action:    action,
		IPAddress: ip,
		UserAgent: userAgent,
	}
	_ = u.auditRepo.Create(&log)
}

// GetMyAuditLogs mengambil riwayat aktivitas milik akun yang sedang login.
func (u *userUsecase) GetMyAuditLogs(userID uint) ([]entity.AuditLogResponse, error) {
	if u.auditRepo == nil {
		return []entity.AuditLogResponse{}, nil
	}
	logs, err := u.auditRepo.FindByUserID(userID)
	if err != nil {
		return nil, errors.New("gagal mengambil catatan aktivitas")
	}
	var res []entity.AuditLogResponse
	for _, l := range logs {
		res = append(res, entity.FormatAuditLog(l))
	}
	return res, nil
}

// GetAllAuditLogs mengambil seluruh catatan aktivitas sistem (Khusus Superadmin & Owner).
func (u *userUsecase) GetAllAuditLogs() ([]entity.AuditLogResponse, error) {
	if u.auditRepo == nil {
		return []entity.AuditLogResponse{}, nil
	}
	logs, err := u.auditRepo.FindAll()
	if err != nil {
		return nil, errors.New("gagal mengambil seluruh catatan aktivitas")
	}
	var res []entity.AuditLogResponse
	for _, l := range logs {
		res = append(res, entity.FormatAuditLog(l))
	}
	return res, nil
}

// GetTrashedUsers mengambil seluruh data pengguna yang di-soft delete (Khusus Superadmin & Owner).
func (u *userUsecase) GetTrashedUsers() ([]entity.UserResponse, error) {
	users, err := u.userRepo.FindTrashedUsers()
	if err != nil {
		return nil, errors.New("gagal mengambil data pengguna terhapus")
	}
	var res []entity.UserResponse
	for _, user := range users {
		res = append(res, entity.FormatUser(user))
	}
	return res, nil
}

// RestoreUser memulihkan pengguna yang terhapus secara soft delete.
func (u *userUsecase) RestoreUser(targetUserID uint) error {
	err := u.userRepo.RestoreUser(targetUserID)
	if err != nil {
		return errors.New("gagal memulihkan pengguna")
	}
	return nil
}

// PermanentDeleteUser menghapus pengguna secara permanen dari PostgreSQL.
func (u *userUsecase) PermanentDeleteUser(targetUserID uint) error {
	err := u.userRepo.PermanentDeleteUser(targetUserID)
	if err != nil {
		return errors.New("gagal menghapus pengguna secara permanen")
	}
	return nil
}
