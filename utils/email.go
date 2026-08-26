package utils

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"math/big"
	"net/smtp"
	"os"
)

// GenerateOTP menghasilkan kode OTP 6 digit angka secara acak & aman.
func GenerateOTP() (string, error) {
	const digits = "0123456789"
	otp := make([]byte, 6)
	for i := 0; i < 6; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		otp[i] = digits[num.Int64()]
	}
	return string(otp), nil
}

// sendMailSSL mengirim email khusus untuk SMTP dengan Port 465 (Implicit SSL/TLS Connection).
func sendMailSSL(addr, host string, auth smtp.Auth, from string, to []string, msg []byte) error {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         host,
	}

	conn, err := tls.Dial("tcp", addr, tlsConfig)
	if err != nil {
		return fmt.Errorf("gagal melakukan tls.Dial ke %s: %w", addr, err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, host)
	if err != nil {
		return fmt.Errorf("gagal membuat smtp client SSL: %w", err)
	}
	defer client.Quit()

	if auth != nil {
		if ok, _ := client.Extension("AUTH"); ok {
			if err = client.Auth(auth); err != nil {
				return fmt.Errorf("gagal autentikasi SSL SMTP: %w", err)
			}
		}
	}

	if err = client.Mail(from); err != nil {
		return fmt.Errorf("gagal menetapkan sender email: %w", err)
	}

	for _, k := range to {
		if err = client.Rcpt(k); err != nil {
			return fmt.Errorf("gagal menetapkan recipient email: %w", err)
		}
	}

	w, err := client.Data()
	if err != nil {
		return fmt.Errorf("gagal membuat data writer: %w", err)
	}

	_, err = w.Write(msg)
	if err != nil {
		return fmt.Errorf("gagal menulis isi email: %w", err)
	}

	err = w.Close()
	if err != nil {
		return fmt.Errorf("gagal menutup writer email: %w", err)
	}

	return nil
}

// SendResetPasswordEmail mengirimkan email dengan Tailwind CSS Styling & Inline Fallback.
// Mendukung Port 587 (STARTTLS) & Port 465 (Implicit SSL/TLS).
func SendResetPasswordEmail(toEmail, otpCode string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	senderEmail := os.Getenv("SMTP_SENDER_EMAIL")
	senderPassword := os.Getenv("SMTP_SENDER_PASSWORD")

	// Selalu tampilkan di log server agar dev/testing sangat mudah dilakukan
	log.Printf("==================================================")
	log.Printf("🔑 [RESET PASSWORD OTP] Email: %s | OTP Code: %s", toEmail, otpCode)
	log.Printf("==================================================")

	// Jika SMTP credentials tidak lengkap, gunakan mode DEV (Console Log Only)
	if smtpHost == "" || smtpPort == "" || senderEmail == "" || senderPassword == "" {
		log.Printf("ℹ️ SMTP tidak dikonfigurasi penuh di .env. Menggunakan mode Dev Log.")
		return nil
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}
	resetURL := fmt.Sprintf("%s/forgot-password", frontendURL)

	subject := "Subject: 🔑 Kode Reset Password Anda - TailAdmin Vue 3\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="id">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Reset Password</title>
		</head>
		<body style="background-color: #f1f5f9; padding: 24px; font-family: ui-sans-serif, system-ui, sans-serif; color: #1e293b; margin: 0;">
			<div style="max-width: 480px; margin: 0 auto; background-color: #ffffff; border-radius: 16px; border: 1px solid #e2e8f0; overflow: hidden; box-shadow: 0 20px 25px -5px rgba(0,0,0,0.1);">
				<div style="background: linear-gradient(135deg, #4f46e5 0%%, #4338ca 100%%); padding: 32px; text-align: center; color: #ffffff;">
					<div style="display: inline-block; background-color: rgba(255,255,255,0.2); color: #ffffff; font-size: 11px; font-weight: 700; padding: 4px 12px; border-radius: 9999px; letter-spacing: 0.05em; text-transform: uppercase; margin-bottom: 12px;">
						TailAdmin Vue 3
					</div>
					<h1 style="font-size: 22px; font-weight: 700; color: #ffffff; margin: 0;">
						Reset Kata Sandi Akun
					</h1>
				</div>

				<div style="padding: 28px;">
					<p style="font-size: 14px; color: #334155; margin: 0;">
						Halo <b style="color: #0f172a;">%s</b>,
					</p>
					
					<p style="font-size: 14px; color: #475569; line-height: 1.6; margin-top: 12px;">
						Kami menerima permintaan untuk mereset kata sandi akun Anda. Gunakan kode OTP 6-digit di bawah ini untuk mereset password:
					</p>

					<div style="margin: 24px 0; padding: 20px; background-color: #e0e7ff; border: 2px dashed #6366f1; border-radius: 12px; text-align: center;">
						<div style="font-size: 11px; text-transform: uppercase; font-weight: 700; color: #4f46e5; letter-spacing: 0.1em; margin-bottom: 4px;">
							Kode OTP Verifikasi Anda
						</div>
						<div style="font-size: 32px; font-weight: 800; color: #1e1b4b; letter-spacing: 8px; font-family: monospace;">
							%s
						</div>
					</div>

					<div style="text-align: center; margin: 24px 0;">
						<a href="%s" style="display: inline-block; padding: 12px 24px; background-color: #4f46e5; color: #ffffff !important; font-weight: 600; font-size: 14px; border-radius: 8px; text-decoration: none;">
							Buka Halaman Reset Password &rarr;
						</a>
					</div>

					<div style="padding: 14px; background-color: #fffbeb; border-left: 4px solid #f59e0b; border-radius: 0 8px 8px 0; font-size: 12px; color: #78350f; line-height: 1.5; margin-top: 20px;">
						<strong style="display: block; font-weight: 700; margin-bottom: 4px;">🔒 Informasi Keamanan:</strong>
						• Kode OTP berlaku selama <b>15 Menit</b>.<br>
						• Jangan pernah memberitahukan kode ini kepada siapa pun.<br>
						• Jika Anda tidak meminta reset password, abaikan email ini.
					</div>
				</div>

				<div style="background-color: #f8fafc; padding: 16px; text-align: center; border-top: 1px solid #f1f5f9; font-size: 11px; color: #94a3b8;">
					Vue 3 + TailAdmin Dashboard &bull; Golang Clean Architecture &copy; 2026
				</div>
			</div>
		</body>
		</html>
	`, toEmail, otpCode, resetURL)

	smtpUser := os.Getenv("SMTP_USER")
	if smtpUser == "" {
		smtpUser = senderEmail
	}

	msg := []byte(subject + mime + body)
	auth := smtp.PlainAuth("", smtpUser, senderPassword, smtpHost)
	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	var err error
	if smtpPort == "465" {
		err = sendMailSSL(addr, smtpHost, auth, senderEmail, []string{toEmail}, msg)
	} else {
		err = smtp.SendMail(addr, auth, senderEmail, []string{toEmail}, msg)
	}

	if err != nil {
		log.Printf("⚠️ Gagal mengirim email via SMTP (%s): %v", addr, err)
		return nil
	}

	log.Printf("📧 Email reset password berhasil dikirim ke %s via SMTP (%s)", toEmail, addr)
	return nil
}

// SendPasswordChangedSuccessEmail mengirimkan email konfirmasi bahwa password telah berhasil diubah/direset.
func SendPasswordChangedSuccessEmail(toEmail, userName string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	senderEmail := os.Getenv("SMTP_SENDER_EMAIL")
	senderPassword := os.Getenv("SMTP_SENDER_PASSWORD")

	if userName == "" {
		userName = toEmail
	}

	// Selalu tampilkan di log server agar dev/testing mudah dilakukan
	log.Printf("==================================================")
	log.Printf("✅ [PASSWORD CHANGED CONFIRMATION] Email: %s (%s)", toEmail, userName)
	log.Printf("==================================================")

	// Jika SMTP credentials tidak lengkap, gunakan mode DEV (Console Log Only)
	if smtpHost == "" || smtpPort == "" || senderEmail == "" || senderPassword == "" {
		log.Printf("ℹ️ SMTP tidak dikonfigurasi penuh di .env. Menggunakan mode Dev Log.")
		return nil
	}

	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:5173"
	}
	loginURL := fmt.Sprintf("%s/login", frontendURL)

	subject := "Subject: ✅ Kata Sandi Akun Anda Berhasil Diperbarui - TailAdmin Vue 3\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="id">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Kata Sandi Berhasil Diperbarui</title>
		</head>
		<body style="background-color: #f1f5f9; padding: 24px; font-family: ui-sans-serif, system-ui, sans-serif; color: #1e293b; margin: 0;">
			<div style="max-width: 480px; margin: 0 auto; background-color: #ffffff; border-radius: 16px; border: 1px solid #e2e8f0; overflow: hidden; box-shadow: 0 20px 25px -5px rgba(0,0,0,0.1);">
				
				<!-- Brand Header Emerald Success -->
				<div style="background: linear-gradient(135deg, #10b981 0%%, #047857 100%%); padding: 32px; text-align: center; color: #ffffff;">
					<div style="display: inline-block; background-color: rgba(255,255,255,0.2); color: #ffffff; font-size: 11px; font-weight: 700; padding: 4px 12px; border-radius: 9999px; letter-spacing: 0.05em; text-transform: uppercase; margin-bottom: 12px;">
						TailAdmin Vue 3
					</div>
					<h1 style="font-size: 22px; font-weight: 700; color: #ffffff; margin: 0;">
						Kata Sandi Berhasil Diperbarui!
					</h1>
				</div>

				<div style="padding: 28px;">
					<p style="font-size: 14px; color: #334155; margin: 0;">
						Halo <b style="color: #0f172a;">%s</b>,
					</p>
					
					<p style="font-size: 14px; color: #475569; line-height: 1.6; margin-top: 12px;">
						Email ini mengonfirmasi bahwa kata sandi untuk akun Anda (<b style="color: #0f172a;">%s</b>) telah <b>berhasil diperbarui</b>.
					</p>

					<div style="text-align: center; margin: 24px 0;">
						<a href="%s" style="display: inline-block; padding: 12px 24px; background-color: #10b981; color: #ffffff !important; font-weight: 600; font-size: 14px; border-radius: 8px; text-decoration: none;">
							Masuk Ke Akun Anda &rarr;
						</a>
					</div>

					<div style="padding: 14px; background-color: #fef2f2; border-left: 4px solid #ef4444; border-radius: 0 8px 8px 0; font-size: 12px; color: #991b1b; line-height: 1.5; margin-top: 20px;">
						<strong style="display: block; font-weight: 700; margin-bottom: 4px;">⚠️ Peringatan Keamanan:</strong>
						Jika Anda <b>TIDAK</b> merasa melakukan perubahan kata sandi ini, akun Anda mungkin dalam bahaya. Segera gunakan fitur Lupa Password atau hubungi tim dukungan kami.
					</div>
				</div>

				<div style="background-color: #f8fafc; padding: 16px; text-align: center; border-top: 1px solid #f1f5f9; font-size: 11px; color: #94a3b8;">
					Vue 3 + TailAdmin Dashboard &bull; Golang Clean Architecture &copy; 2026
				</div>
			</div>
		</body>
		</html>
	`, userName, toEmail, loginURL)

	smtpUser := os.Getenv("SMTP_USER")
	if smtpUser == "" {
		smtpUser = senderEmail
	}

	msg := []byte(subject + mime + body)
	auth := smtp.PlainAuth("", smtpUser, senderPassword, smtpHost)
	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	var err error
	if smtpPort == "465" {
		err = sendMailSSL(addr, smtpHost, auth, senderEmail, []string{toEmail}, msg)
	} else {
		err = smtp.SendMail(addr, auth, senderEmail, []string{toEmail}, msg)
	}

	if err != nil {
		log.Printf("⚠️ Gagal mengirim email konfirmasi ganti password via SMTP (%s): %v", addr, err)
		return nil
	}

	log.Printf("📧 Email konfirmasi ganti password berhasil dikirim ke %s via SMTP (%s)", toEmail, addr)
	return nil
}
