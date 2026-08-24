package usecase

import (
	"errors"

	"projectgolangnyoba/entity"
	"projectgolangnyoba/repository"
	"projectgolangnyoba/utils"
)

// UserUsecase mendefinisikan kontrak logika bisnis terkait User.
type UserUsecase interface {
	Register(input entity.RegisterInput) (entity.UserResponse, error)
	Login(input entity.LoginInput) (string, error)
	GetProfile(userID uint) (entity.UserResponse, error)
	UpdateProfile(userID uint, input entity.UpdateProfileInput) (entity.UserResponse, error)
	ChangePassword(userID uint, input entity.ChangePasswordInput) error
}

// userUsecase implementasi dari UserUsecase yang bergantung pada UserRepository.
type userUsecase struct {
	userRepo repository.UserRepository
}

// NewUserUsecase adalah konstruktor untuk membuat instance UserUsecase.
func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: userRepo}
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

	// 3. Buat entity User baru
	user := entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
	}

	// 4. Simpan ke database melalui repository
	err = u.userRepo.Create(&user)
	if err != nil {
		return entity.UserResponse{}, err
	}

	// 5. Kembalikan data user dalam format UserResponse (tanpa password)
	return entity.FormatUser(user), nil
}

// Login memverifikasi kredensial email & password dan menghasilkan JWT Token.
func (u *userUsecase) Login(input entity.LoginInput) (string, error) {
	// 1. Cari user berdasarkan email
	user, err := u.userRepo.FindByEmail(input.Email)
	if err != nil {
		return "", errors.New("email atau password salah")
	}

	// 2. Verifikasi kesesuaian password
	if !utils.CheckPasswordHash(input.Password, user.Password) {
		return "", errors.New("email atau password salah")
	}

	// 3. Generate JWT Token
	token, err := utils.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", errors.New("gagal membuat token autentikasi")
	}

	return token, nil
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

