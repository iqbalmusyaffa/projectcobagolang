package utils

import (
	"crypto/rand"
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

// SendResetPasswordEmail mengirimkan email dengan Tailwind CSS Styling & Inline Fallback.
// Jika SMTP belum dikonfigurasi, OTP akan dicetak ke console terminal.
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
			<script src="https://cdn.tailwindcss.com"></script>
		</head>
		<body class="bg-slate-100 p-4 sm:p-6 font-sans text-slate-800 antialiased" style="background-color: #f1f5f9; padding: 24px; font-family: ui-sans-serif, system-ui, sans-serif; color: #1e293b; margin: 0;">
			<div class="max-w-md mx-auto bg-white rounded-xl shadow-xl border border-slate-200 overflow-hidden" style="max-width: 480px; margin: 0 auto; background-color: #ffffff; border-radius: 16px; border: 1px solid #e2e8f0; overflow: hidden; box-shadow: 0 20px 25px -5px rgba(0,0,0,0.1), 0 8px 10px -6px rgba(0,0,0,0.1);">
				
				<!-- Brand Header dengan Tailwind Indigo Gradient -->
				<div class="bg-gradient-to-r from-indigo-600 to-indigo-700 p-8 text-center text-white" style="background: linear-gradient(135deg, #4f46e5 0%%, #4338ca 100%%); padding: 32px; text-align: center; color: #ffffff;">
					<div class="inline-block bg-white/20 text-white text-xs font-bold px-3 py-1 rounded-full tracking-wider uppercase mb-3" style="display: inline-block; background-color: rgba(255,255,255,0.2); color: #ffffff; font-size: 11px; font-weight: 700; padding: 4px 12px; border-radius: 9999px; letter-spacing: 0.05em; text-transform: uppercase; margin-bottom: 12px;">
						TailAdmin Vue 3
					</div>
					<h1 class="text-2xl font-bold tracking-tight text-white m-0" style="font-size: 22px; font-weight: 700; color: #ffffff; margin: 0;">
						Reset Kata Sandi Akun
					</h1>
				</div>

				<!-- Content Area -->
				<div class="p-6 sm:p-8 space-y-5" style="padding: 28px;">
					<p class="text-sm text-slate-700 m-0" style="font-size: 14px; color: #334155; margin: 0;">
						Halo <b class="text-slate-900" style="color: #0f172a;">%s</b>,
					</p>
					
					<p class="text-sm text-slate-600 leading-relaxed m-0" style="font-size: 14px; color: #475569; line-height: 1.6; margin-top: 12px;">
						Kami menerima permintaan untuk mereset kata sandi akun Anda. Gunakan kode OTP 6-digit di bawah ini untuk mereset password:
					</p>

					<!-- OTP Card Tailwind Styled -->
					<div class="my-6 p-5 bg-indigo-50 border-2 border-dashed border-indigo-500 rounded-xl text-center" style="margin: 24px 0; padding: 20px; background-color: #e0e7ff; border: 2px dashed #6366f1; border-radius: 12px; text-align: center;">
						<div class="text-xs uppercase font-bold text-indigo-600 tracking-wider mb-1" style="font-size: 11px; text-transform: uppercase; font-weight: 700; color: #4f46e5; letter-spacing: 0.1em; margin-bottom: 4px;">
							Kode OTP Verifikasi Anda
						</div>
						<div class="text-3xl font-extrabold text-indigo-950 tracking-[8px] font-mono" style="font-size: 32px; font-weight: 800; color: #1e1b4b; letter-spacing: 8px; font-family: monospace;">
							%s
						</div>
					</div>

					<!-- Tombol Tailwind Button -->
					<div class="text-center my-6" style="text-align: center; margin: 24px 0;">
						<a href="%s" class="inline-block px-6 py-3 bg-indigo-600 hover:bg-indigo-700 text-white font-semibold text-sm rounded-lg shadow-md transition duration-200 no-underline" style="display: inline-block; padding: 12px 24px; background-color: #4f46e5; color: #ffffff !important; font-weight: 600; font-size: 14px; border-radius: 8px; text-decoration: none; box-shadow: 0 4px 6px -1px rgba(79, 70, 229, 0.3);">
							Buka Halaman Reset Password &rarr;
						</a>
					</div>

					<!-- Tailwind Alert Box -->
					<div class="p-4 bg-amber-50 border-l-4 border-amber-500 rounded-r-lg text-xs text-amber-900 leading-relaxed" style="padding: 14px; background-color: #fffbeb; border-left: 4px solid #f59e0b; border-radius: 0 8px 8px 0; font-size: 12px; color: #78350f; line-height: 1.5; margin-top: 20px;">
						<strong class="font-bold block mb-1" style="display: block; font-weight: 700; margin-bottom: 4px;">🔒 Informasi Keamanan:</strong>
						• Kode OTP berlaku selama <b class="font-bold" style="font-weight: 700;">15 Menit</b>.<br>
						• Jangan pernah memberitahukan kode ini kepada siapa pun.<br>
						• Jika Anda tidak meminta reset password, abaikan email ini.
					</div>
				</div>

				<!-- Footer Tailwind -->
				<div class="bg-slate-50 p-4 text-center border-t border-slate-100 text-xs text-slate-400" style="background-color: #f8fafc; padding: 16px; text-align: center; border-top: 1px solid #f1f5f9; font-size: 11px; color: #94a3b8;">
					Vue 3 + TailAdmin Dashboard &bull; Golang Clean Architecture &copy; 2026
				</div>
			</div>
		</body>
		</html>
	`, toEmail, otpCode, resetURL)

	msg := []byte(subject + mime + body)
	auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpHost)
	addr := fmt.Sprintf("%s:%s", smtpHost, smtpPort)

	err := smtp.SendMail(addr, auth, senderEmail, []string{toEmail}, msg)
	if err != nil {
		log.Printf("⚠️ Gagal mengirim email via SMTP (%s): %v", addr, err)
		return nil
	}

	log.Printf("📧 Email reset password berhasil dikirim ke %s via SMTP (%s)", toEmail, addr)
	return nil
}
