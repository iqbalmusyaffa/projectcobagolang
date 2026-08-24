package handler

import (
	"net/http"

	"projectgolangnyoba/entity"
	"projectgolangnyoba/usecase"

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
		"message": "Registrasi akun berhasil",
		"data":    userResponse,
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

	// 2. Panggil Usecase untuk login & dapatkan Access & Refresh Token
	tokenPair, err := h.userUsecase.Login(input)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
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



