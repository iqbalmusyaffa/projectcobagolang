package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"projectgolangnyoba/entity"
	"projectgolangnyoba/repository"
	"projectgolangnyoba/utils"
)

// hashToken membuat hash SHA-256 untuk memetakan refresh token ke sesi database.
func hashToken(token string) string {
	h := sha256.Sum256([]byte(token))
	return hex.EncodeToString(h[:])
}

// UserUsecase mendefinisikan kontrak logika bisnis terkait User, Aktivasi Akun, & Sesi Aktif.
type UserUsecase interface {
	Register(input entity.RegisterInput) (entity.UserResponse, error)
	VerifyEmail(input entity.VerifyEmailInput) error
	ResendVerification(input entity.ResendVerificationInput) error
	AdminCreateUser(input entity.AdminCreateUserInput) (entity.UserResponse, error)
	AdminUpdateUser(input entity.AdminUpdateUserInput) (entity.UserResponse, error)
	AdminDeleteUser(targetUserID uint) error
	Login(input entity.LoginInput, ipAddress, userAgent string) (entity.TokenPairResponse, error)
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
	GetActiveSessions(userID uint, currentRefreshToken string) ([]entity.UserSessionResponse, error)
	RevokeSession(userID, sessionID uint) error
	RevokeOtherSessions(userID uint, currentRefreshToken string) error
}

// userUsecase implementasi dari UserUsecase yang bergantung pada Repository User, Redis, AuditLog, & Session.
type userUsecase struct {
	userRepo    repository.UserRepository
	redisRepo   repository.RedisRepository
	auditRepo   repository.AuditLogRepository
	sessionRepo repository.SessionRepository
}

// NewUserUsecase adalah konstruktor untuk membuat instance UserUsecase.
func NewUserUsecase(
	userRepo repository.UserRepository,
	redisRepo repository.RedisRepository,
	auditRepo repository.AuditLogRepository,
	sessionRepo repository.SessionRepository,
) UserUsecase {
	return &userUsecase{
		userRepo:    userRepo,
		redisRepo:   redisRepo,
		auditRepo:   auditRepo,
		sessionRepo: sessionRepo,
	}
}

// Register memproses pendaftaran akun pengguna baru dengan status belum diverifikasi dan mengirim OTP aktivasi.
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

	// 3. Generate kode OTP 6-digit untuk verifikasi akun
	otpCode, err := utils.GenerateOTP()
	if err != nil {
		otpCode = "123456"
	}
	expiresAt := time.Now().Add(15 * time.Minute)

	role := input.Role
	if role == "" {
		role = "user"
	}

	user := entity.User{
		Name:                       input.Name,
		Email:                      input.Email,
		Password:                   hashedPassword,
		Role:                       role,
		IsEmailVerified:            false,
		VerificationToken:          otpCode,
		VerificationTokenExpiresAt: &expiresAt,
	}

	// 4. Simpan ke database melalui repository
	err = u.userRepo.Create(&user)
	if err != nil {
		return entity.UserResponse{}, err
	}

	// 5. Simpan OTP ke Redis (TTL 15 Menit)
	if u.redisRepo != nil {
		_ = u.redisRepo.SetCache("verify_code:"+user.Email, otpCode, 15*time.Minute)
	}

	// 6. Kirim email aktivasi secara asynchronous
	go utils.SendVerificationEmail(user.Email, otpCode)

	return entity.FormatUser(user), nil
}

// VerifyEmail memverifikasi akun pengguna menggunakan kode OTP 6-digit.
func (u *userUsecase) VerifyEmail(input entity.VerifyEmailInput) error {
	user, err := u.userRepo.FindByEmail(input.Email)
	if err != nil || user == nil {
		return errors.New("email tidak terdaftar di sistem")
	}

	if user.IsEmailVerified {
		return errors.New("email sudah diverifikasi sebelumnya, silakan langsung login")
	}

	var isValidOTP bool
	// 1. Cek dari Redis Cache terlebih dahulu
	if u.redisRepo != nil {
		cachedOTP, err := u.redisRepo.GetCache("verify_code:" + input.Email)
		if err == nil && cachedOTP == input.OTP {
			isValidOTP = true
		}
	}

	// 2. Fallback cek ke PostgreSQL jika Redis tidak ada
	if !isValidOTP {
		if user.VerificationToken != "" && user.VerificationToken == input.OTP {
			if user.VerificationTokenExpiresAt != nil && user.VerificationTokenExpiresAt.After(time.Now()) {
				isValidOTP = true
			}
		}
	}

	if !isValidOTP {
		return errors.New("kode OTP aktivasi tidak valid atau sudah kadaluwarsa (15 menit)")
	}

	// 3. Update status email verified di database
	err = u.userRepo.MarkEmailVerified(user.ID)
	if err != nil {
		return errors.New("gagal memverifikasi akun")
	}

	// 4. Hapus OTP dari Redis
	if u.redisRepo != nil {
		_ = u.redisRepo.DeleteCache("verify_code:" + input.Email)
	}

	return nil
}

// ResendVerification membuat dan mengirimkan kembali kode OTP aktivasi baru via email.
func (u *userUsecase) ResendVerification(input entity.ResendVerificationInput) error {
	user, err := u.userRepo.FindByEmail(input.Email)
	if err != nil || user == nil {
		return errors.New("email tidak terdaftar di sistem")
	}

	if user.IsEmailVerified {
		return errors.New("email sudah diverifikasi sebelumnya, silakan langsung login")
	}

	otpCode, err := utils.GenerateOTP()
	if err != nil {
		return errors.New("gagal membuat kode verifikasi baru")
	}

	expiresAt := time.Now().Add(15 * time.Minute)

	if u.redisRepo != nil {
		_ = u.redisRepo.SetCache("verify_code:"+user.Email, otpCode, 15*time.Minute)
	}
	_ = u.userRepo.SaveVerificationToken(user.Email, otpCode, expiresAt)

	go utils.SendVerificationEmail(user.Email, otpCode)

	return nil
}

// AdminCreateUser memproses pendaftaran pengguna baru khusus oleh Super Admin dengan menentukan role (langsung aktif/terverifikasi).
func (u *userUsecase) AdminCreateUser(input entity.AdminCreateUserInput) (entity.UserResponse, error) {
	existingUser, _ := u.userRepo.FindByEmail(input.Email)
	if existingUser != nil {
		return entity.UserResponse{}, errors.New("email sudah terdaftar, silakan gunakan email lain")
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return entity.UserResponse{}, errors.New("gagal mengamankan password")
	}

	role := input.Role
	if role != "superadmin" && role != "owner" && role != "admin" && role != "user" {
		role = "user"
	}

	now := time.Now()
	user := entity.User{
		Name:            input.Name,
		Email:           input.Email,
		Password:        hashedPassword,
		Role:            role,
		IsEmailVerified: true,
		EmailVerifiedAt: &now,
	}

	err = u.userRepo.Create(&user)
	if err != nil {
		return entity.UserResponse{}, err
	}

	return entity.FormatUser(user), nil
}

// Login memverifikasi kredensial email & password, memvalidasi status verifikasi email, serta mencatat sesi aktif.
func (u *userUsecase) Login(input entity.LoginInput, ipAddress, userAgent string) (entity.TokenPairResponse, error) {
	// 1. Cari user berdasarkan email
	user, err := u.userRepo.FindByEmail(input.Email)
	if err != nil {
		return entity.TokenPairResponse{}, errors.New("email atau password salah")
	}

	// 2. Verifikasi kesesuaian password
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return entity.TokenPairResponse{}, errors.New("email atau password salah")
	}

	// 3. Validasi status verifikasi email
	if !user.IsEmailVerified {
		return entity.TokenPairResponse{}, errors.New("UNVERIFIED_EMAIL: Akun Anda belum diverifikasi. Silakan aktivasi melalui link/kode OTP yang telah dikirim ke email Anda.")
	}

	role := user.Role
	if role == "" {
		role = "user"
	}

	// 4. Generate Access Token (15 Menit) & Refresh Token (7 Hari)
	accessToken, err := utils.GenerateAccessToken(user.ID, user.Email, role)
	if err != nil {
		return entity.TokenPairResponse{}, errors.New("gagal membuat access token")
	}

	refreshToken, err := utils.GenerateRefreshToken(user.ID, user.Email)
	if err != nil {
		return entity.TokenPairResponse{}, errors.New("gagal membuat refresh token")
	}

	// 5. Simpan Refresh Token ke Redis dengan TTL 7 Hari (168 Jam)
	if u.redisRepo != nil {
		_ = u.redisRepo.StoreRefreshToken(user.ID, refreshToken, 7*24*time.Hour)
	}

	// 6. Catat Sesi Perangkat Aktif di PostgreSQL
	if u.sessionRepo != nil {
		deviceName, browser, os, deviceType := utils.ParseUserAgent(userAgent)
		refreshHash := hashToken(refreshToken)

		session := entity.UserSession{
			UserID:           user.ID,
			RefreshTokenHash: refreshHash,
			IPAddress:        ipAddress,
			UserAgent:        userAgent,
			DeviceName:       deviceName,
			DeviceType:       deviceType,
			Browser:          browser,
			OS:               os,
			LastActiveAt:     time.Now(),
			ExpiresAt:        time.Now().Add(7 * 24 * time.Hour),
		}
		_ = u.sessionRepo.Create(&session)
	}

	return entity.TokenPairResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// RefreshToken memvalidasi Refresh Token, memperbarui keaktifan sesi, dan mengembalikan Access Token baru.
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
		role = "user"
	}

	// 4. Terbitkan Access Token 15 Menit yang baru
	newAccessToken, err := utils.GenerateAccessToken(user.ID, user.Email, role)
	if err != nil {
		return "", errors.New("gagal membuat access token baru")
	}

	// 5. Perbarui last_active_at pada data sesi aktif
	if u.sessionRepo != nil {
		hash := hashToken(refreshTokenStr)
		sess, err := u.sessionRepo.FindByRefreshTokenHash(hash)
		if err == nil && sess != nil {
			_ = u.sessionRepo.UpdateLastActive(sess.ID)
		}
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
	if input.Role != "superadmin" && input.Role != "owner" && input.Role != "admin" && input.Role != "user" {
		return entity.UserResponse{}, errors.New("role tidak valid (pilih: superadmin, owner, admin, atau user)")
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

// LogoutToken memasukkan token JWT ke Redis Blacklist, menghapus Refresh Token & menghapus sesi aktif.
func (u *userUsecase) LogoutToken(token string) error {
	if token != "" {
		claims, err := utils.ValidateToken(token)
		if err == nil && claims.UserID != 0 {
			if u.redisRepo != nil {
				_ = u.redisRepo.DeleteRefreshToken(claims.UserID)
			}
		}
		if u.redisRepo != nil {
			return u.redisRepo.BlacklistToken(token, 24*time.Hour)
		}
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

// DeleteAccount menghapus akun secara Soft Delete, meng-invalidasi token, dan menghapus seluruh sesi.
func (u *userUsecase) DeleteAccount(userID uint, token string) error {
	user, err := u.userRepo.FindByID(userID)
	if err != nil {
		return errors.New("user tidak ditemukan")
	}

	err = u.userRepo.Delete(user.ID)
	if err != nil {
		return errors.New("gagal menghapus akun")
	}

	if u.sessionRepo != nil {
		_ = u.sessionRepo.DeleteByUserID(userID)
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

	if u.sessionRepo != nil {
		_ = u.sessionRepo.DeleteByUserID(targetUserID)
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
	user, err := u.userRepo.FindByEmail(email)
	if err != nil || user == nil {
		return errors.New("email tidak terdaftar di sistem")
	}

	otpCode, err := utils.GenerateOTP()
	if err != nil {
		return errors.New("gagal menghasilkan kode OTP")
	}

	expiresAt := time.Now().Add(15 * time.Minute)

	if u.redisRepo != nil {
		_ = u.redisRepo.SetCache("reset_code:"+email, otpCode, 15*time.Minute)
	}

	_ = u.userRepo.SaveResetToken(email, otpCode, expiresAt)

	return utils.SendResetPasswordEmail(email, otpCode)
}

// ResetPassword memverifikasi OTP dan mengganti password pengguna dengan yang baru.
func (u *userUsecase) ResetPassword(email, otp, newPassword string) error {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil || user == nil {
		return errors.New("email tidak terdaftar di sistem")
	}

	var isValidOTP bool
	if u.redisRepo != nil {
		cachedOTP, err := u.redisRepo.GetCache("reset_code:" + email)
		if err == nil && cachedOTP == otp {
			isValidOTP = true
		}
	}

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

	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		return errors.New("gagal mengamankan password baru")
	}

	user.Password = hashedPassword

	err = u.userRepo.Update(user)
	if err != nil {
		return errors.New("gagal memperbarui password")
	}

	if u.redisRepo != nil {
		_ = u.redisRepo.DeleteCache("reset_code:" + email)
	}
	_ = u.userRepo.ClearResetToken(user.ID)

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

	if u.sessionRepo != nil {
		_ = u.sessionRepo.DeleteByUserID(input.UserID)
	}

	if u.redisRepo != nil {
		_ = u.redisRepo.DeleteRefreshToken(input.UserID)
	}

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

// GetActiveSessions mengambil daftar seluruh sesi & perangkat aktif pengguna.
func (u *userUsecase) GetActiveSessions(userID uint, currentRefreshToken string) ([]entity.UserSessionResponse, error) {
	if u.sessionRepo == nil {
		return []entity.UserSessionResponse{}, nil
	}

	sessions, err := u.sessionRepo.FindByUserID(userID)
	if err != nil {
		return nil, errors.New("gagal mengambil daftar sesi aktif")
	}

	currentHash := ""
	if currentRefreshToken != "" {
		currentHash = hashToken(currentRefreshToken)
	}

	var response []entity.UserSessionResponse
	for _, sess := range sessions {
		isCurrent := false
		if currentHash != "" && sess.RefreshTokenHash == currentHash {
			isCurrent = true
		}
		response = append(response, entity.FormatUserSession(sess, isCurrent))
	}

	return response, nil
}

// RevokeSession mencabut sesi spesifik tertentu (Logout perangkat tunggal).
func (u *userUsecase) RevokeSession(userID, sessionID uint) error {
	if u.sessionRepo == nil {
		return nil
	}

	sess, err := u.sessionRepo.FindByID(sessionID)
	if err != nil || sess == nil {
		return errors.New("sesi tidak ditemukan")
	}

	if sess.UserID != userID {
		return errors.New("tidak memiliki izin untuk mencabut sesi ini")
	}

	return u.sessionRepo.Delete(sessionID)
}

// RevokeOtherSessions mencabut seluruh sesi pengguna selain perangkat yang sedang digunakan saat ini.
func (u *userUsecase) RevokeOtherSessions(userID uint, currentRefreshToken string) error {
	if u.sessionRepo == nil {
		return nil
	}

	var currentSessionID uint = 0
	if currentRefreshToken != "" {
		hash := hashToken(currentRefreshToken)
		sess, err := u.sessionRepo.FindByRefreshTokenHash(hash)
		if err == nil && sess != nil {
			currentSessionID = sess.ID
		}
	}

	if currentSessionID != 0 {
		return u.sessionRepo.DeleteByUserIDExcept(userID, currentSessionID)
	}

	return u.sessionRepo.DeleteByUserID(userID)
}
