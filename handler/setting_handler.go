package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"projectgolangnyoba/entity"
	"projectgolangnyoba/usecase"

	"github.com/gin-gonic/gin"
)

// SettingHandler menangani permintaan HTTP terkait pengaturan sistem dan koneksi SMTP.
type SettingHandler struct {
	settingUsecase usecase.SettingUsecase
	userUsecase    usecase.UserUsecase
}

// NewSettingHandler membuat instance konstruktor untuk SettingHandler.
func NewSettingHandler(settingUsecase usecase.SettingUsecase, userUsecase usecase.UserUsecase) *SettingHandler {
	return &SettingHandler{
		settingUsecase: settingUsecase,
		userUsecase:    userUsecase,
	}
}

// GetSettings handler untuk GET /api/settings (Protected - Ambil Pengaturan Aplikasi)
func (h *SettingHandler) GetSettings(c *gin.Context) {
	setting, err := h.settingUsecase.GetSettings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	maskedSetting := *setting
	if maskedSetting.SMTPSenderPassword != "" {
		maskedSetting.SMTPSenderPassword = "••••••••"
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Pengaturan sistem berhasil diambil",
		"data":    maskedSetting,
	})
}

// UpdateSettings handler untuk PUT /api/admin/settings (Khusus Superadmin & Owner)
func (h *SettingHandler) UpdateSettings(c *gin.Context) {
	var input entity.SystemSettingInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Input tidak valid",
			"error":   err.Error(),
		})
		return
	}

	setting, err := h.settingUsecase.UpdateSettings(input)
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
		go h.userUsecase.LogActivity(aid, admin.Name, admin.Role, "Memperbarui Pengaturan Sistem & Server SMTP", c.ClientIP(), c.Request.UserAgent())
	}

	maskedSetting := *setting
	if maskedSetting.SMTPSenderPassword != "" {
		maskedSetting.SMTPSenderPassword = "••••••••"
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Pengaturan sistem berhasil diperbarui",
		"data":    maskedSetting,
	})
}

// TestEmail handler untuk POST /api/admin/settings/test-email (Khusus Superadmin & Owner)
func (h *SettingHandler) TestEmail(c *gin.Context) {
	var input entity.TestEmailInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Alamat email tujuan uji coba tidak valid",
			"error":   err.Error(),
		})
		return
	}

	err := h.settingUsecase.TestSMTPEmail(input.TargetEmail)
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
		go h.userUsecase.LogActivity(aid, admin.Name, admin.Role, fmt.Sprintf("Mengirimkan tes email SMTP ke %s", input.TargetEmail), c.ClientIP(), c.Request.UserAgent())
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": fmt.Sprintf("Email uji coba SMTP berhasil terkirim ke '%s'!", input.TargetEmail),
	})
}

// UploadLogo handler untuk POST /api/admin/settings/logo (Khusus Superadmin & Owner)
func (h *SettingHandler) UploadLogo(c *gin.Context) {
	file, err := c.FormFile("logo")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Berkas logo tidak ditemukan dalam request",
		})
		return
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" && ext != ".svg" && ext != ".webp" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Format gambar logo tidak didukung. Harap unggah PNG, JPG, JPEG, SVG, atau WEBP",
		})
		return
	}

	filename := fmt.Sprintf("logo_%d%s", time.Now().Unix(), ext)
	dst := filepath.Join("images", filename)

	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Gagal menyimpan berkas logo",
		})
		return
	}

	logoPath := fmt.Sprintf("images/%s", filename)
	setting, err := h.settingUsecase.UpdateLogo(logoPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Logo sistem berhasil diperbarui",
		"data":    setting,
	})
}
