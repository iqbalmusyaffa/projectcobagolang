# Golang REST API & Vue 3 TailAdmin Dashboard (Clean Architecture)

Proyek fullstack modern berbasis **Golang REST API** (Gin Framework, GORM PostgreSQL, Redis) menggunakan **Clean Architecture** dan **Vue 3 Frontend (Tailwind CSS / TailAdmin)**.

---

## 🚀 Teknologi Utama (Tech Stack)

- **Backend:** Golang (Gin Gonic Framework)
- **Database:** PostgreSQL (GORM ORM dengan Auto-Migration)
- **Caching & Keamanan:** Redis (JWT Token Blacklisting, Refresh Tokens, & OTP Caching)
- **Pengiriman Email:** Standard Go `net/smtp` (HTML Email Template dengan Tailwind CSS & Fallback Log Dev)
- **Frontend:** Vue 3 (Vite, Vue Router, Axios, Tailwind CSS / TailAdmin)
- **Arsitektur:** Clean Architecture (Entity, Repository, Usecase, Handler/Controller, Middleware, Utils)

---

## 📁 Struktur Project

```text
projectgolangnyoba/
├── config/                # Koneksi Database PostgreSQL & Redis Client
│   ├── database.go        # GORM Init & AutoMigrate
│   ├── redis.go           # Redis Client & Blacklist Token Store
│   └── seed.go            # Seeder Akun Default
├── entity/                # Struct Model GORM & DTO (Request/Response)
│   └── user.go            # User entity, Reset OTP DTOs, Admin DTOs
├── repository/            # Data Access Layer / Database Queries
│   ├── user_repository.go # Repository PostgreSQL (Users & Reset Tokens)
│   └── redis_repository.go# Repository Redis (Blacklist & Cache OTP)
├── usecase/               # Business Logic Layer
│   └── user_usecase.go    # Logika Bisnis (Auth, Profile, Forgot Password, Admin CRUD)
├── handler/               # Presentation Layer / HTTP Handlers
│   └── user_handler.go    # Gin HTTP Controllers (Auth, Forgot Password, Profile, Admin)
├── middleware/            # Auth & RBAC Middleware
│   ├── auth_middleware.go # Validasi Token JWT & Blacklist Check
│   ├── cors_middleware.go # Dynamic CORS Handling
│   └── role_middleware.go # Role-Based Access Control (RBAC)
├── utils/                 # Helper Functions
│   ├── password.go        # Hash & Compare Bcrypt
│   ├── jwt.go             # Generate & Verify Access/Refresh JWT
│   └── email.go           # Generate OTP 6-Digit & SMTP HTML Mailer (Tailwind CSS)
├── images/                # Penyimpanan lokal foto profil (Avatar)
├── frontend/              # Single Page Application Vue 3 + Vite
│   ├── src/
│   │   ├── views/
│   │   │   ├── LoginView.vue
│   │   │   ├── RegisterView.vue
│   │   │   ├── ForgotPasswordView.vue # Halaman Request OTP & Reset Password
│   │   │   └── DashboardView.vue      # Dashboard Utama & Halaman Manajemen Admin
│   │   ├── router/
│   │   │   └── index.js    # Navigation Guards & Routes
│   │   └── services/
│   │       └── api.js      # Axios Client dengan Authorization Header Interceptor
├── .env.example           # Contoh Konfigurasi Environment Variable
├── go.mod                 # Dependency Go
└── main.go                # Entry Point & Route Engine
```

---

## 🌟 Fitur Unggulan

### 1. Autentikasi & Keamanan Tingkat Lanjut
- **JWT Access Token & Refresh Token:** Autentikasi terpisah dengan token refresh 7 hari.
- **Redis Token Blacklisting:** Token otomatis dimasukkan ke daftar hitam Redis saat pengguna melakukan Logout atau Soft Delete Akun.
- **RBAC (Role-Based Access Control):** Peran hirarki (`superadmin`, `owner`, `admin`, `user`).
- **Toggle Password Visibility:** Ikon mata interaktif (Eye / Eye-Slash) di seluruh form (Login, Register, Ganti Password, Reset Password).

### 2. Email Verification & Aktivasi Akun (OTP 6-Digit)
- **Registrasi dengan Status Verifikasi:** Pengguna baru yang mendaftar memiliki status `is_email_verified = false`.
- **Pengiriman Kode Aktivasi:** Kode OTP 6-digit dikirimkan melalui email HTML Tailwind CSS (dan dicetak di terminal server jika mode Dev).
- **Halaman Aktivasi (`/verify-email`):** Input 6-digit OTP responsif, countdown timer 60 detik kirim ulang kode, dan auto-verifikasi via URL query (`?email=...&otp=...`).
- **Proteksi Login:** Akun yang belum diverifikasi otomatis ditolak saat login dengan tombol instan menuju halaman aktivasi.

### 3. Manajemen Sesi & Perangkat Aktif (Active Sessions & Devices)
- **Pencatatan Perangkat:** Setiap login mencatat nama peramban, sistem operasi, jenis perangkat (Desktop, Mobile, Tablet), IP address, dan timestamp aktivitas.
- **Badge Sesi Utama:** Penanda *"Perangkat Ini"* (*Current Session*) pada sesi yang sedang digunakan.
- **Revoke Sesi:** Pengguna dapat memutuskan sesi tertentu (*"Putuskan Sesi"*) atau memutuskan seluruh perangkat lain sekaligus (*"Putuskan Semua Sesi Lain"*).

### 4. Fitur Lupa Password & Reset via Email (OTP 6-Digit)
- **Request OTP 6-Digit:** Pengguna meminta OTP yang dikirimkan ke email terdaftar, berlaku selama 15 Menit (disimpan di Redis & PostgreSQL fallback).
- **Pengiriman Email SMTP / Dev Log:** Menggunakan email HTML Tailwind CSS modern via `net/smtp`. Jika SMTP belum di-set di `.env`, kode OTP otomatis dicetak ke console log terminal untuk kemudahan testing.
- **Reset Password Form:** Halaman 2-Langkah modern `/forgot-password` untuk menukarkan kode OTP 6-digit dengan kata sandi baru.

### 5. Fitur Khusus Super Admin & Bantuan User
- **Lihat & Kelola Seluruh Pengguna:** Menampilkan tabel seluruh pengguna terdaftar.
- **Kirim OTP Reset Password (✉️):** Superadmin dapat memicu pengiriman kode OTP reset password ke email pengguna secara otomatis dari tabel.
- **Reset Password Langsung (🔒):** Memungkinkan Superadmin **langsung mereset password pengguna secara instan tanpa OTP** untuk membantu pengguna yang mengalami kendala/bug pada email.
- **Tambah & Edit Pengguna:** Membuat dan mengubah data pengguna lain (Nama, Email, Peran, No HP, Gender, Tgl Lahir, Alamat, Bio).
- **Hapus Pengguna (Soft Delete) & Instant Role Switcher:** Menghapus pengguna (Soft Delete) serta mengubah peran pengguna secara cepat.

### 6. Manajemen Profil & Dashboard Analytics (Chart.js)
- **Biodata Lengkap & Upload Avatar:** Mengelola biodata dan mengunggah foto avatar disajikan via static `/images`.
- **Badge Status Verifikasi:** Status badge hijau *"Email Terverifikasi"* di bagian header profil.
- **Realtime Stats Cards & Grafik Interaktif:** Doughnut Chart dan Bar Chart untuk visualisasi statistik distribusi peran pengguna.

---

## 🔗 Dokumentasi REST API

| Method | Endpoint | Akses | Deskripsi |
| :--- | :--- | :--- | :--- |
| `POST` | `/api/register` | Publik | Registrasi akun baru (Status: Unverified) |
| `POST` | `/api/verify-email` | Publik | Verifikasi & aktivasi akun via OTP 6-digit |
| `POST` | `/api/resend-verification` | Publik | Mengirim ulang kode OTP aktivasi akun |
| `POST` | `/api/login` | Publik | Login & mendapatkan JWT Access/Refresh Token |
| `POST` | `/api/refresh-token` | Publik | Memperbarui Access Token dengan Refresh Token |
| `POST` | `/api/forgot-password` | Publik | Meminta kode OTP reset password via email |
| `POST` | `/api/reset-password` | Publik | Mereset password dengan kode OTP 6-digit |
| `GET` | `/api/profile` | Protected | Mengambil data profil login saat ini |
| `PUT` | `/api/profile` | Protected | Memperbarui biodata profil saat ini |
| `POST` | `/api/profile/avatar` | Protected | Mengunggah foto profil avatar baru |
| `DELETE` | `/api/profile` | Protected | Menghapus akun sendiri (Soft Delete) |
| `PUT` | `/api/change-password` | Protected | Mengubah kata sandi akun |
| `GET` | `/api/sessions` | Protected | Mengambil daftar seluruh sesi & perangkat aktif |
| `DELETE` | `/api/sessions/:id` | Protected | Mencabut sesi pada perangkat tertentu |
| `POST` | `/api/sessions/revoke-others` | Protected | Mencabut seluruh sesi pada perangkat lain |
| `POST` | `/api/logout` | Protected | Logout & blacklist token ke Redis |
| `GET` | `/api/dashboard/stats` | Protected | Mengambil data statistik agregasi dashboard & grafik |
| `GET` | `/api/admin/users` | Admin & Superadmin | Mengambil daftar seluruh pengguna |
| `POST` | `/api/superadmin/users` | Khusus Superadmin | Membuat pengguna baru (`superadmin`, `owner`, `admin`, `user`) |
| `PUT` | `/api/superadmin/users` | Khusus Superadmin | Memperbarui data pengguna lain |
| `DELETE` | `/api/superadmin/users/:id` | Khusus Superadmin | Menghapus pengguna lain (Soft Delete) |
| `PUT` | `/api/superadmin/change-role` | Khusus Superadmin | Mengubah Peran (Role) pengguna |
| `PUT` | `/api/superadmin/users/reset-password` | Khusus Superadmin | Reset password pengguna secara langsung tanpa OTP |

---

## 🛠️ Konfigurasi Environment (`.env`)

Salin file `.env.example` menjadi `.env` dan atur konfigurasi environment Anda:

```env
# Database PostgreSQL Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_db_password
DB_NAME=golang_db
DB_SSLMODE=disable

# Redis Configuration (Caching & Token Blacklist)
REDIS_HOST=localhost
REDIS_PORT=6379
REDIS_PASSWORD=

# JWT Secret Key
JWT_SECRET=your_jwt_secret_key_here

# Application Port & Frontend URL
PORT=8080
FRONTEND_URL=http://localhost:5173

# SMTP Email Configuration (Reset Password - Opsional, jika kosong OTP akan dicetak di Log Terminal Server)
SMTP_HOST=smtp.gmail.com
SMTP_PORT=587
SMTP_SENDER_EMAIL=your_email@gmail.com
SMTP_SENDER_PASSWORD=your_app_password
```

---

## 🛠️ Cara Menjalankan Aplikasi

### ⚡ 1. Jalankan Backend (Golang REST API)
```bash
# 1. Download dependency Go
go mod tidy

# 2. Pastikan PostgreSQL & Redis aktif

# 3. Jalankan server Golang
go run main.go
```
*Backend berjalan di: `http://localhost:8080`*

### 💻 2. Jalankan Frontend (Vue 3 TailAdmin)
```bash
# 1. Masuk ke folder frontend
cd frontend

# 2. Install dependency Node.js
npm install

# 3. Menjalankan server dev Vite
npm run dev
```
*Frontend berjalan di: `http://localhost:5173`*

---

## 👤 Akun Default (Database Seeder)

Aplikasi secara otomatis menyisipkan 4 akun default setiap kali server Golang dinyalakan (`go run main.go`):

| Role | Email | Password | Hak Akses |
| :--- | :--- | :--- | :--- |
| **SUPERADMIN** | `faridwimansyah8@gmail.com` | `password123` | Hak Akses Penuh Sistem (CRUD User, Direct Reset Password, Switch Role) |
| **OWNER** | `owner@demo.com` | `password123` | Akses Laporan & Manajemen User Read-only |
| **ADMIN** | `admin@demo.com` | `password123` | Akses Admin Standar Operasional |
| **USER** | `user@demo.com` | `password123` | Akses Pengguna Biasa / Regular User |
