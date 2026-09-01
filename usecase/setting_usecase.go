package usecase

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/smtp"
	"os"
	"projectgolangnyoba/entity"
	"projectgolangnyoba/repository"
)

// SettingUsecase mendefinisikan logika bisnis terkait pengaturan sistem & pengujian SMTP.
type SettingUsecase interface {
	GetSettings() (*entity.SystemSetting, error)
	UpdateSettings(input entity.SystemSettingInput) (*entity.SystemSetting, error)
	UpdateLogo(logoPath string) (*entity.SystemSetting, error)
	TestSMTPEmail(targetEmail string) error
}

type settingUsecase struct {
	settingRepo repository.SettingRepository
	auditRepo   repository.AuditLogRepository
}

// NewSettingUsecase adalah konstruktor untuk menginisialisasi usecase pengaturan sistem.
func NewSettingUsecase(settingRepo repository.SettingRepository, auditRepo repository.AuditLogRepository) SettingUsecase {
	return &settingUsecase{
		settingRepo: settingRepo,
		auditRepo:   auditRepo,
	}
}

// GetSettings mengambil pengaturan sistem saat ini.
func (u *settingUsecase) GetSettings() (*entity.SystemSetting, error) {
	return u.settingRepo.GetSettings()
}

// UpdateSettings memperbarui pengaturan umum dan kredensial SMTP server di database.
func (u *settingUsecase) UpdateSettings(input entity.SystemSettingInput) (*entity.SystemSetting, error) {
	setting, err := u.settingRepo.GetSettings()
	if err != nil {
		return nil, errors.New("gagal mengambil pengaturan sistem")
	}

	if input.AppName != "" {
		setting.AppName = input.AppName
	}
	setting.AppDescription = input.AppDescription
	if input.DefaultLanguage != "" {
		setting.DefaultLanguage = input.DefaultLanguage
	}
	if input.Timezone != "" {
		setting.Timezone = input.Timezone
	}
	setting.SMTPHost = input.SMTPHost
	if input.SMTPPort != "" {
		setting.SMTPPort = input.SMTPPort
	}
	setting.SMTPSenderEmail = input.SMTPSenderEmail
	if input.SMTPSenderPassword != "" && input.SMTPSenderPassword != "••••••••" {
		setting.SMTPSenderPassword = input.SMTPSenderPassword
	}
	setting.SMTPUser = input.SMTPUser
	setting.EnableEmailNotification = input.EnableEmailNotification

	err = u.settingRepo.UpdateSettings(setting)
	if err != nil {
		return nil, errors.New("gagal menyimpan pengaturan sistem")
	}

	// Perbarui variabel environment sistem secara dinamis jika diatur di database
	if setting.SMTPHost != "" {
		_ = os.Setenv("SMTP_HOST", setting.SMTPHost)
		_ = os.Setenv("SMTP_PORT", setting.SMTPPort)
		_ = os.Setenv("SMTP_SENDER_EMAIL", setting.SMTPSenderEmail)
		if setting.SMTPSenderPassword != "" {
			_ = os.Setenv("SMTP_SENDER_PASSWORD", setting.SMTPSenderPassword)
		}
		_ = os.Setenv("SMTP_USER", setting.SMTPUser)
	}

	return setting, nil
}

// UpdateLogo memperbarui lokasi file logo sistem di database.
func (u *settingUsecase) UpdateLogo(logoPath string) (*entity.SystemSetting, error) {
	err := u.settingRepo.UpdateLogo(logoPath)
	if err != nil {
		return nil, errors.New("gagal memperbarui logo aplikasi")
	}
	return u.settingRepo.GetSettings()
}

// TestSMTPEmail mengirimkan email pengujian HTML secara realtime untuk memverifikasi koneksi SMTP.
func (u *settingUsecase) TestSMTPEmail(targetEmail string) error {
	setting, err := u.settingRepo.GetSettings()
	if err != nil || setting.SMTPHost == "" || setting.SMTPSenderEmail == "" || setting.SMTPSenderPassword == "" {
		// Fallback ke ENV jika di DB belum terisi
		host := os.Getenv("SMTP_HOST")
		port := os.Getenv("SMTP_PORT")
		sender := os.Getenv("SMTP_SENDER_EMAIL")
		pass := os.Getenv("SMTP_SENDER_PASSWORD")
		if host == "" || sender == "" || pass == "" {
			return errors.New("pengaturan SMTP belum diisi secara lengkap di sistem atau .env")
		}
		setting.SMTPHost = host
		setting.SMTPPort = port
		setting.SMTPSenderEmail = sender
		setting.SMTPSenderPassword = pass
		setting.SMTPUser = os.Getenv("SMTP_USER")
	}

	smtpUser := setting.SMTPUser
	if smtpUser == "" {
		smtpUser = setting.SMTPSenderEmail
	}

	subject := "Subject: 🚀 Uji Coba Koneksi SMTP Server - " + setting.AppName + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<body style="font-family: Arial, sans-serif; background-color: #f4f6f8; margin: 0; padding: 20px;">
			<div style="max-width: 600px; margin: 0 auto; background: #ffffff; border-radius: 12px; overflow: hidden; border: 1px solid #e2e8f0; box-shadow: 0 4px 12px rgba(0,0,0,0.05);">
				<div style="background: linear-gradient(135deg, #2563eb, #4f46e5); color: #ffffff; padding: 24px; text-align: center;">
					<h2 style="margin: 0; font-size: 20px; font-weight: 700;">🚀 Tes Koneksi SMTP Berhasil!</h2>
					<p style="margin: 4px 0 0 0; font-size: 13px; opacity: 0.9;">%s</p>
				</div>
				<div style="padding: 24px; color: #334155; line-height: 1.6; font-size: 14px;">
					<p style="margin-top: 0;">Halo,</p>
					<p>Email ini dikirimkan untuk menguji pengiriman email dari server SMTP <b>%s</b> pada port <b>%s</b>.</p>
					<div style="padding: 12px 16px; background-color: #f1f5f9; border-radius: 8px; font-family: monospace; font-size: 12px; margin: 16px 0;">
						• Host: %s<br>
						• Port: %s<br>
						• Sender Email: %s<br>
						• Target Email: %s
					</div>
					<p style="margin-bottom: 0;">Server SMTP Anda siap digunakan untuk notifikasi email & reset OTP password!</p>
				</div>
			</div>
		</body>
		</html>
	`, setting.AppName, setting.SMTPHost, setting.SMTPPort, setting.SMTPHost, setting.SMTPPort, setting.SMTPSenderEmail, targetEmail)

	msg := []byte(subject + mime + body)
	auth := smtp.PlainAuth("", smtpUser, setting.SMTPSenderPassword, setting.SMTPHost)
	addr := fmt.Sprintf("%s:%s", setting.SMTPHost, setting.SMTPPort)

	if setting.SMTPPort == "465" {
		conn, dialErr := tls.Dial("tcp", addr, &tls.Config{InsecureSkipVerify: true, ServerName: setting.SMTPHost})
		if dialErr != nil {
			return fmt.Errorf("gagal terhubung ke SSL Port 465 (%s): %v", addr, dialErr)
		}
		defer conn.Close()

		client, clientErr := smtp.NewClient(conn, setting.SMTPHost)
		if clientErr != nil {
			return fmt.Errorf("gagal membuat SMTP Client SSL: %v", clientErr)
		}
		defer client.Quit()

		if auth != nil {
			if ok, _ := client.Extension("AUTH"); ok {
				if authErr := client.Auth(auth); authErr != nil {
					return fmt.Errorf("autentikasi SSL SMTP gagal: %v", authErr)
				}
			}
		}

		if mailErr := client.Mail(setting.SMTPSenderEmail); mailErr != nil {
			return mailErr
		}
		if rcptErr := client.Rcpt(targetEmail); rcptErr != nil {
			return rcptErr
		}

		w, wErr := client.Data()
		if wErr != nil {
			return wErr
		}
		_, _ = w.Write(msg)
		_ = w.Close()
	} else {
		err = smtp.SendMail(addr, auth, setting.SMTPSenderEmail, []string{targetEmail}, msg)
		if err != nil {
			return fmt.Errorf("gagal mengirim tes email via SMTP (%s): %v", addr, err)
		}
	}

	return nil
}
