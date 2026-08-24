# Golang Login & Register REST API + TailAdmin Frontend (Clean Architecture)

Project REST API menggunakan **Golang** & **PostgreSQL** dengan **Clean Architecture (Layered Architecture)** dan dilengkapi **Frontend TailAdmin (Tailwind CSS)**.

## 📁 Struktur Project

```text
projectgolangnyoba/
├── config/
│   └── database.go        # Koneksi & Migrasi PostgreSQL via GORM
├── entity/
│   └── user.go            # Model database & Struct DTO (Request/Response)
├── repository/
│   └── user_repository.go # Interface & Query Database PostgreSQL
├── usecase/
│   └── user_usecase.go    # Logika Bisnis (Register, Hash Bcrypt, Login, JWT)
├── handler/
│   └── user_handler.go    # HTTP Controller & HTML Template Renderer
├── middleware/
│   └── auth_middleware.go # Validasi Token JWT untuk Endpoint Terproteksi
├── utils/
│   ├── password.go       # Helper Hash & Compare Password (bcrypt)
│   └── jwt.go            # Helper Generate & Validate JWT Token
├── views/                 # Template Frontend Tailwind CSS / TailAdmin
│   ├── login.html        # Form Login TailAdmin
│   ├── register.html     # Form Register TailAdmin
│   └── dashboard.html    # Dashboard TailAdmin Free Template UI
├── .env.example           # Contoh variabel lingkungan
├── go.mod                 # Dependency Go
└── main.go                # Entry point aplikasi, Web Routes, & REST API
```

---

## 🛠️ Prasyarat (Prerequisites)

1. **Golang** (v1.20+) terinstal di komputer Anda.
2. **PostgreSQL** sudah berjalan di komputer lokal (port `5432`).
3. Buat database baru di PostgreSQL bernama `golang_db`:
   ```sql
   CREATE DATABASE golang_db;
   ```

---

## 🚀 Cara Menjalankan Aplikasi

1. **Download Dependency:**
   ```bash
   go mod tidy
   ```

2. **Jalankan Aplikasi:**
   ```bash
   go run main.go
   ```

3. **Buka di Browser:**
   - **Login Page:** `http://localhost:8080/login`
   - **Register Page:** `http://localhost:8080/register`
   - **Dashboard Page:** `http://localhost:8080/dashboard`

---

## 📌 Fitur & Alur Autentikasi Frontend

1. **Registrasi Akun:**
   Buka `http://localhost:8080/register`, isi Nama, Email, dan Password. Form akan mengirim `POST /api/register` lalu meredirect ke halaman Login.

2. **Login Akun:**
   Buka `http://localhost:8080/login`, isi Email dan Password. Form akan mengirim `POST /api/login`, menerima Token JWT, menyimpannya di `localStorage`, dan meredirect ke Dashboard.

3. **Dashboard TailAdmin:**
   Buka `http://localhost:8080/dashboard`. Javascript akan mengambil token dari `localStorage` dan memanggil `GET /api/profile` secara otomatis untuk merender Nama & Email Pengguna secara dinamis.

4. **Logout:**
   Klik tombol **Keluar (Logout)** di Sidebar / Header. Token JWT akan dihapus dari `localStorage` dan halaman dikembalikan ke Login.
