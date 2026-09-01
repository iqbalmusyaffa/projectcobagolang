package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"projectgolangnyoba/entity"
	"projectgolangnyoba/usecase"
	"projectgolangnyoba/utils"

	"github.com/gin-gonic/gin"
)

// UserHandler menangani permintaan HTTP terkait user dan meneruskannya ke Usecase.
type UserHandler struct {
	userUsecase usecase.UserUsecase
}

// NewUserHandler konstruktor untuk membuat instance UserHandler.
func NewUserHandler(userUsecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

// Register handler untuk endpoint POST /api/register
func (h *UserHandler) Register(c *gin.Context) {
	var input entity.RegisterInput

	// 1. Parsing & Validasi JSON Body request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
		return
	}

	// 2. Panggil Usecase untuk registrasi
	userResponse, err := h.userUsecase.Register(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	// 3. Kembalikan respons sukses 201 Created
	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Registrasi akun berhasil! Kode OTP aktivasi telah dikirim ke email Anda.",
		"data":    userResponse,
	})
}

// VerifyEmail handler untuk endpoint POST /api/verify-email
func (h *UserHandler) VerifyEmail(c *gin.Context) {
	var input entity.VerifyEmailInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Email dan kode OTP 6-digit wajib diisi",
			"error":   err.Error(),
		})
		return
	}

	err := h.userUsecase.VerifyEmail(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Akun Anda berhasil diverifikasi! Silakan login untuk melanjutkan.",
	})
}

// ResendVerification handler untuk endpoint POST /api/resend-verification
func (h *UserHandler) ResendVerification(c *gin.Context) {
	var input entity.ResendVerificationInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Alamat email wajib diisi",
			"error":   err.Error(),
		})
		return
	}

	err := h.userUsecase.ResendVerification(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Kode OTP aktivasi baru berhasil dikirim ke email Anda.",
	})
}

// Login handler untuk endpoint POST /api/login
func (h *UserHandler) Login(c *gin.Context) {
	var input entity.LoginInput

	// 1. Parsing & Validasi JSON Body request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
		return
	}

	ipAddress := c.ClientIP()
	userAgent := c.Request.UserAgent()

	// 2. Panggil Usecase untuk login & dapatkan Access & Refresh Token
	tokenPair, err := h.userUsecase.Login(input, ipAddress, userAgent)
	if err != nil {
		isUnverified := strings.Contains(err.Error(), "UNVERIFIED_EMAIL")
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":        "error",
			"message":       err.Error(),
			"is_unverified": isUnverified,
			"email":         input.Email,
		})
		return
	}

	// Log activity
	claims, _ := utils.ValidateToken(tokenPair.AccessToken)
	if claims != nil {
		user, err := h.userUsecase.GetProfile(claims.UserID)
		if err == nil {
			go h.userUsecase.LogActivity(user.ID, user.Name, user.Role, "Berhasil Login ke sistem", ipAddress, userAgent)
		}
	}

	// 3. Kembalikan respons sukses 200 OK dengan Access Token & Refresh Token
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Login berhasil",
		"data": gin.H{
			"token":         tokenPair.AccessToken,  // Backward compatibility
			"access_token":  tokenPair.AccessToken,
			"refresh_token": tokenPair.RefreshToken,
		},
	})
}

// RefreshToken handler untuk POST /api/refresh-token
func (h *UserHandler) RefreshToken(c *gin.Context) {
	var input entity.RefreshTokenInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Refresh token wajib diisi",
			"error":   err.Error(),
		})
		return
	}

	newAccessToken, err := h.userUsecase.RefreshToken(input.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Access token baru berhasil diterbitkan",
		"data": gin.H{
			"access_token": newAccessToken,
		},
	})
}

// ForgotPassword handler untuk POST /api/forgot-password
func (h *UserHandler) ForgotPassword(c *gin.Context) {
	var input entity.ForgotPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Alamat email tidak valid",
			"error":   err.Error(),
		})
		return
	}

	err := h.userUsecase.RequestPasswordReset(input.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Kode OTP reset password berhasil dikirim ke email Anda",
	})
}

// ResetPassword handler untuk POST /api/reset-password
func (h *UserHandler) ResetPassword(c *gin.Context) {
	var input entity.ResetPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Input tidak valid. Pastikan OTP 6-digit dan password minimal 6 karakter",
			"error":   err.Error(),
		})
		return
	}

	err := h.userUsecase.ResetPassword(input.Email, input.OTP, input.NewPassword)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Password Anda berhasil diperbarui. Silakan login dengan password baru.",
	})
}

// Profile handler untuk endpoint GET /api/profile (Protected Endpoint)
func (h *UserHandler) Profile(c *gin.Context) {
	// 1. Ambil userID yang diset oleh AuthMiddleware
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Pengguna tidak terautentikasi",
		})
		return
	}

	userID := userIDVal.(uint)

	// 2. Panggil Usecase untuk mendapatkan profil user
	userResponse, err := h.userUsecase.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	// 3. Kembalikan data profil
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Berhasil mengambil profil",
		"data":    userResponse,
	})
}

// UpdateProfile handler untuk endpoint PUT /api/profile (Protected Endpoint)
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Pengguna tidak terautentikasi",
		})
		return
	}

	var input entity.UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
		return
	}

	userID := userIDVal.(uint)
	userResponse, err := h.userUsecase.UpdateProfile(userID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Profil berhasil diperbarui",
		"data":    userResponse,
	})
}

// ChangePassword handler untuk endpoint PUT /api/change-password (Protected Endpoint)
func (h *UserHandler) ChangePassword(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Pengguna tidak terautentikasi",
		})
		return
	}

	var input entity.ChangePasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
		return
	}

	userID := userIDVal.(uint)
	err := h.userUsecase.ChangePassword(userID, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Password berhasil diperbarui",
	})
}

// GetAllUsers handler untuk GET /api/admin/users (Khusus Superadmin & Owner)
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.userUsecase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Berhasil mengambil daftar pengguna",
		"data":    users,
	})
}

// ChangeRole handler untuk PUT /api/superadmin/change-role (Khusus Superadmin)
func (h *UserHandler) ChangeRole(c *gin.Context) {
	var input entity.ChangeRoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
		return
	}

	userResponse, err := h.userUsecase.ChangeUserRole(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Peran pengguna berhasil diperbarui",
		"data":    userResponse,
	})
}

// CreateUserByAdmin handler untuk POST /api/superadmin/users (Khusus Superadmin)
func (h *UserHandler) CreateUserByAdmin(c *gin.Context) {
	var input entity.AdminCreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
		return
	}

	userResponse, err := h.userUsecase.AdminCreateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Pengguna baru berhasil dibuat oleh Super Admin",
		"data":    userResponse,
	})
}

// Logout handler untuk POST /api/logout (Protected Endpoint, Memasukkan Token ke Redis Blacklist)
func (h *UserHandler) Logout(c *gin.Context) {
	tokenVal, exists := c.Get("currentToken")
	if exists {
		tokenStr := tokenVal.(string)
		_ = h.userUsecase.LogoutToken(tokenStr)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Logout berhasil (Token di-blacklist di Redis)",
	})
}

// UploadAvatar handler untuk POST /api/profile/avatar (Protected Endpoint, Upload File Multipart)
func (h *UserHandler) UploadAvatar(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Pengguna tidak terautentikasi",
		})
		return
	}
	userID := userIDVal.(uint)

	// 1. Ambil file dari form multipart 'avatar'
	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "File foto profil (avatar) wajib diunggah",
			"error":   err.Error(),
		})
		return
	}

	// 2. Validasi ekstensi file (.jpg, .jpeg, .png, .webp)
	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".webp" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Format file tidak didukung! Harap unggah file berformat .jpg, .jpeg, .png, atau .webp",
		})
		return
	}

	// 3. Validasi ukuran file (maksimal 2MB)
	if file.Size > 2*1024*1024 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Ukuran file terlalu besar! Maksimal ukuran file adalah 2MB",
		})
		return
	}

	// 4. Buat nama file unik dan simpan ke folder images/
	fileName := fmt.Sprintf("avatar-%d-%d%s", userID, time.Now().Unix(), ext)
	dst := filepath.Join("images", fileName)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Gagal menyimpan file avatar di server",
			"error":   err.Error(),
		})
		return
	}

	// 5. Update avatar path di database via Usecase
	avatarPath := "images/" + fileName
	userResponse, err := h.userUsecase.UploadAvatar(userID, avatarPath)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Foto profil berhasil diperbarui",
		"data":    userResponse,
	})
}

// DeleteAccount handler untuk DELETE /api/profile (Protected Endpoint, Soft Delete)
func (h *UserHandler) DeleteAccount(c *gin.Context) {
	userIDVal, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Pengguna tidak terautentikasi",
		})
		return
	}
	userID := userIDVal.(uint)

	tokenStr := ""
	if tokenVal, exists := c.Get("currentToken"); exists {
		tokenStr = tokenVal.(string)
	}

	err := h.userUsecase.DeleteAccount(userID, tokenStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Akun Anda berhasil dihapus (Soft Delete)",
	})
}

// AdminUpdateUser handler untuk PUT /api/superadmin/users (Khusus Superadmin)
func (h *UserHandler) AdminUpdateUser(c *gin.Context) {
	var input entity.AdminUpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
		return
	}

	userResponse, err := h.userUsecase.AdminUpdateUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Data pengguna berhasil diperbarui oleh Super Admin",
		"data":    userResponse,
	})
}

// AdminDeleteUser handler untuk DELETE /api/superadmin/users/:id (Khusus Superadmin)
func (h *UserHandler) AdminDeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ID pengguna tidak valid",
		})
		return
	}

	err = h.userUsecase.AdminDeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Pengguna berhasil dihapus oleh Super Admin (Soft Delete)",
	})
}

// AdminResetPassword handler untuk PUT /api/superadmin/users/reset-password (Khusus Superadmin)
func (h *UserHandler) AdminResetPassword(c *gin.Context) {
	var input entity.AdminResetPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Input reset password tidak valid. Pastikan new_password minimal 6 karakter.",
			"error":   err.Error(),
		})
		return
	}

	err := h.userUsecase.AdminResetUserPassword(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Password pengguna berhasil diperbarui secara langsung oleh Super Admin",
	})
}

// GetDashboardStats handler untuk GET /api/dashboard/stats (Protected Endpoint)
func (h *UserHandler) GetDashboardStats(c *gin.Context) {
	stats, err := h.userUsecase.GetDashboardStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Gagal mengambil statistik dashboard",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Statistik dashboard berhasil diambil",
		"data":    stats,
	})
}

// GetMyAuditLogs handler untuk GET /api/audit-logs/my (Protected - Riwayat Pengguna Sendiri)
func (h *UserHandler) GetMyAuditLogs(c *gin.Context) {
	userID, exists := c.Get("currentUser")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Pengguna tidak terautentikasi",
		})
		return
	}

	uid, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "ID Pengguna tidak valid",
		})
		return
	}

	logs, err := h.userUsecase.GetMyAuditLogs(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Riwayat aktivitas pengguna berhasil diambil",
		"data":    logs,
	})
}

// GetAllAuditLogs handler untuk GET /api/admin/audit-logs (Khusus Superadmin & Owner)
func (h *UserHandler) GetAllAuditLogs(c *gin.Context) {
	logs, err := h.userUsecase.GetAllAuditLogs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Seluruh catatan aktivitas sistem berhasil diambil",
		"data":    logs,
	})
}

// GetTrashedUsers handler untuk GET /api/admin/users/trashed (Khusus Superadmin & Owner)
func (h *UserHandler) GetTrashedUsers(c *gin.Context) {
	users, err := h.userUsecase.GetTrashedUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Daftar pengguna terhapus berhasil diambil",
		"data":    users,
	})
}

// RestoreUser handler untuk PUT /api/superadmin/users/:id/restore (Khusus Superadmin)
func (h *UserHandler) RestoreUser(c *gin.Context) {
	idParam := c.Param("id")
	targetID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ID Pengguna tidak valid",
		})
		return
	}

	err = h.userUsecase.RestoreUser(uint(targetID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	// Log activity
	adminID, _ := c.Get("currentUser")
	if aid, ok := adminID.(uint); ok {
		admin, _ := h.userUsecase.GetProfile(aid)
		go h.userUsecase.LogActivity(aid, admin.Name, admin.Role, fmt.Sprintf("Memulihkan pengguna: #%d", targetID), c.ClientIP(), c.Request.UserAgent())
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Pengguna berhasil dipulihkan dari tempat sampah",
	})
}

// PermanentDeleteUser handler untuk DELETE /api/superadmin/users/:id/permanent (Khusus Superadmin)
func (h *UserHandler) PermanentDeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	targetID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ID Pengguna tidak valid",
		})
		return
	}

	err = h.userUsecase.PermanentDeleteUser(uint(targetID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	// Log activity
	adminID, _ := c.Get("currentUser")
	if aid, ok := adminID.(uint); ok {
		admin, _ := h.userUsecase.GetProfile(aid)
		go h.userUsecase.LogActivity(aid, admin.Name, admin.Role, fmt.Sprintf("Menghapus permanen pengguna: #%d", targetID), c.ClientIP(), c.Request.UserAgent())
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Pengguna berhasil dihapus secara permanen dari database",
	})
}

// GetActiveSessions handler untuk endpoint GET /api/sessions
func (h *UserHandler) GetActiveSessions(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Pengguna tidak terautentikasi",
		})
		return
	}

	currentRefreshToken := c.GetHeader("X-Refresh-Token")
	if currentRefreshToken == "" {
		currentRefreshToken = c.Query("refresh_token")
	}

	sessions, err := h.userUsecase.GetActiveSessions(userID.(uint), currentRefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Daftar sesi aktif berhasil diambil",
		"data":    sessions,
	})
}

// RevokeSession handler untuk endpoint DELETE /api/sessions/:id
func (h *UserHandler) RevokeSession(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Pengguna tidak terautentikasi",
		})
		return
	}

	sessionIDStr := c.Param("id")
	sessionID, err := strconv.ParseUint(sessionIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ID sesi tidak valid",
		})
		return
	}

	err = h.userUsecase.RevokeSession(userID.(uint), uint(sessionID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	// Log activity
	user, errProfile := h.userUsecase.GetProfile(userID.(uint))
	if errProfile == nil {
		go h.userUsecase.LogActivity(user.ID, user.Name, user.Role, fmt.Sprintf("Memutuskan sesi perangkat #%d", sessionID), c.ClientIP(), c.Request.UserAgent())
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Sesi perangkat berhasil diputuskan",
	})
}

// RevokeOtherSessions handler untuk endpoint POST /api/sessions/revoke-others
func (h *UserHandler) RevokeOtherSessions(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": "Pengguna tidak terautentikasi",
		})
		return
	}

	currentRefreshToken := c.GetHeader("X-Refresh-Token")
	if currentRefreshToken == "" {
		currentRefreshToken = c.Query("refresh_token")
	}

	err := h.userUsecase.RevokeOtherSessions(userID.(uint), currentRefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	// Log activity
	user, errProfile := h.userUsecase.GetProfile(userID.(uint))
	if errProfile == nil {
		go h.userUsecase.LogActivity(user.ID, user.Name, user.Role, "Memutuskan seluruh sesi perangkat lain", c.ClientIP(), c.Request.UserAgent())
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Semua sesi perangkat lain berhasil diputuskan",
	})
}

// ShowAPIStatus merender halaman status REST API Golang


func (h *UserHandler) ShowAPIStatus(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// ShowLoginPage merender halaman status API
func (h *UserHandler) ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// ShowRegisterPage merender halaman status API
func (h *UserHandler) ShowRegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// ShowDashboardPage merender halaman status API
func (h *UserHandler) ShowDashboardPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}



