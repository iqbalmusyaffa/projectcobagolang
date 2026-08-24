# Golang REST API & Vue 3 TailAdmin Dashboard (Clean Architecture)

Proyek fullstack modern berbasis **Golang REST API** (Gin Framework, GORM PostgreSQL, Redis) menggunakan **Clean Architecture** dan **Vue 3 Frontend (Tailwind CSS / TailAdmin)**.

---

## 🚀 Teknologi Utama (Tech Stack)

- **Backend:** Golang (Gin Gonic Framework)
- **Database:** PostgreSQL (GORM ORM dengan Auto-Migration)
- **Caching & Keamanan:** Redis (JWT Token Blacklisting & Refresh Tokens)
- **Frontend:** Vue 3 (Vite, Vue Router, Axios, Tailwind CSS / TailAdmin)
- **Arsitektur:** Clean Architecture (Entity, Repository, Usecase, Handler/Controller, Middleware)

---

## 📁 Struktur Project

```text
projectgolangnyoba/
├── config/                # Koneksi Database PostgreSQL & Redis Client
│   ├── database.go        # GORM Init & AutoMigrate
│   └── redis.go           # Redis Client & Blacklist Token Store
├── entity/                # Struct Model GORM & DTO (Request/Response)
│   └── user.go            # User entity, DTOs (Profile, Admin, Role, Auth)
├── repository/            # Data Access Layer / Database Queries
│   ├── user_repository.go # Repository Interface & Implementasi PostgreSQL
│   └── redis_repository.go# Repository Blacklist Token Redis
├── usecase/               # Business Logic Layer
│   └── user_usecase.go    # Logika Bisnis (Register, Login, Avatar, Admin CRUD)
├── handler/               # Presentation Layer / HTTP Handlers
│   └── user_handler.go    # Gin HTTP Controllers (Auth, Profile, Admin)
├── middleware/            # Auth & RBAC Middleware
│   └── auth_middleware.go # Validasi Token JWT & Role-Based Access Control
├── utils/                 # Helper Functions
│   ├── password.go        # Hash & Compare Bcrypt
│   └── jwt.go             # Generate & Verify Access/Refresh JWT
├── images/                # Penyimpanan lokal foto profil (Avatar)
├── frontend/              # Single Page Application Vue 3 + Vite
│   ├── src/
│   │   ├── views/
│   │   │   ├── LoginView.vue
│   │   │   ├── RegisterView.vue
│   │   │   └── DashboardView.vue # Dashboard Utama & Halaman Profil Dedicated
│   │   └── services/
│   │       └── api.js      # Axios Client dengan Authorization Header Interceptor
├── go.mod                 # Dependency Go
└── main.go                # Entry Point & Route Engine
```

---

## 🌟 Fitur Unggulan

### 1. Autentikasi & Keamanan Tingkat Lanjut
- **JWT Access Token & Refresh Token:** Autentikasi terpisah dengan token refresh.
- **Redis Token Blacklisting:** Token otomatis dimasukkan ke dalam daftar hitam Redis saat pengguna melakukan Logout atau Soft Delete Akun.
- **RBAC (Role-Based Access Control):** Peran hirarki (`superadmin`, `owner`, `admin`, `user`).

### 2. Manajemen Profil Pengguna (User Profile)
- **Biodata Lengkap:** Menyimpan Nama, Email, Nomor Telepon/HP (`phone`), Jenis Kelamin (`gender`), Tanggal Lahir (`birth_date`), Alamat Lengkap (`address`), Bio (`bio`), serta Timestamp `created_at` & `updated_at`.
- **Upload Foto Profil (Avatar):** Mendukung pengunggahan foto avatar (`JPG`, `PNG`, `WEBP` maks 2MB) disajikan via static file `/images`.
- **Hapus Akun Mandiri (Soft Delete):** Fitur *Danger Zone* untuk menonaktifkan akun sendiri.

### 3. Fitur Khusus Super Admin (User Management)
- **Lihat Semua Pengguna:** Menampilkan tabel seluruh pengguna terdaftar.
- **Tambah User Baru:** Membuat akun pengguna langsung oleh Super Admin.
- **Edit Data Pengguna Lain:** Super Admin dapat mengedit Nama, Email, Peran, No HP, Gender, Tgl Lahir, Alamat, dan Bio pengguna lain.
- **Hapus Pengguna Lain (Soft Delete):** Menghapus pengguna lain dari sistem serta mencabut token refresh di Redis.
- **Instant Role Switcher:** Mengubah hak akses pengguna secara cepat (`admin`, `owner`, `superadmin`, `user`).

### 4. Dashboard Analytics & Visualisasi Grafik (Chart.js)
- **Realtime Stats Cards:** Menampilkan 5 Card Ringkasan Statistik (Total Pengguna, Super Admin, Owner, Admin Standar, dan User Biasa).
- **Grafik Interaktif Chart.js:**
  - **Doughnut Chart:** Visualisasi persentase distribusi 4 peran pengguna (RBAC).
  - **Bar Chart:** Komparasi jumlah akun aktif berdasarkan kategori peran.
- **API Endpoint:** `GET /api/dashboard/stats` menyajikan statistik agregasi SQL PostgreSQL.

### 5. Interaktivitas UI & Notifikasi Real-Time
- **Sistem Notifikasi Lonceng Header:** Popover notifikasi modern dengan badge unread counter, penanda kategori (Security, Profile, System), notifikasi sesi dinamis per pengguna, serta fitur *Tandai Dibaca* & *Hapus Notifikasi*.
- **Tombol Sinkron Data (Data Sync):** Fitur sinkronisasi instan dari database PostgreSQL dari Topbar Header tanpa merefresh browser.
- **Pratinjau Foto Profil & Download Avatar:** Modal lightbox foto profil resolusi tinggi yang dilengkapi tombol **Unduh Foto Profil**.

---

## 🔗 Dokumentasi REST API

| Method | Endpoint | Akses | Deskripsi |
| :--- | :--- | :--- | :--- |
| `POST` | `/api/register` | Publik | Registrasi akun baru (Default Role: `user`) |
| `POST` | `/api/login` | Publik | Login & mendapatkan JWT Access/Refresh Token |
| `POST` | `/api/refresh` | Publik | Memperbarui Access Token dengan Refresh Token |
| `GET` | `/api/profile` | Protected | Mengambil data profil login saat ini |
| `PUT` | `/api/profile` | Protected | Memperbarui biodata profil saat ini |
| `POST` | `/api/profile/avatar` | Protected | Mengunggah foto profil avatar baru |
| `DELETE` | `/api/profile` | Protected | Menghapus akun sendiri (Soft Delete) |
| `PUT` | `/api/change-password` | Protected | Mengubah kata sandi akun |
| `POST` | `/api/logout` | Protected | Logout & blacklist token ke Redis |
| `GET` | `/api/dashboard/stats` | Protected | Mengambil data statistik agregasi dashboard & grafik |
| `GET` | `/api/admin/users` | Admin & Superadmin | Mengambil daftar seluruh pengguna |
| `POST` | `/api/superadmin/users` | Khusus Superadmin | Membuat pengguna baru (`superadmin`, `owner`, `admin`, `user`) |
| `PUT` | `/api/superadmin/users` | Khusus Superadmin | Memperbarui data pengguna lain |
| `DELETE` | `/api/superadmin/users/:id` | Khusus Superadmin | Menghapus pengguna lain (Soft Delete) |
| `PUT` | `/api/superadmin/change-role` | Khusus Superadmin | Mengubah Peran (Role) pengguna |

---

## 🛠️ Cara Menjalankan Aplikasi

### 📥 1. Clone & Setup Repository
```bash
# 1. Clone repository dari GitHub
git clone https://github.com/iqbalmusyaffa/projectcobagolang.git

# 2. Masuk ke direktori proyek
cd projectcobagolang

# 3. Pindah ke branch fitur terbaru
git checkout feature/add-user-role
```

### ⚡ 2. Jalankan Backend (Golang REST API)
```bash
# 1. Download dependency Go
go mod tidy

# 2. Pastikan PostgreSQL & Redis aktif di komputer Anda

# 3. Jalankan server Golang (Server otomatis melakukan GORM AutoMigrate & Seeding)
go run main.go
```
*Backend berjalan di: `http://localhost:8080`*

### 💻 3. Jalankan Frontend (Vue 3 TailAdmin)
```bash
# 1. Masuk ke folder frontend
cd frontend

# 2. Install dependency Node.js
npm install

# 3. Menjalankan server lokal Vite
npm run dev
```
*Frontend berjalan di: `http://localhost:5173`*

---

## 👤 Akun Default (Database Seeder)

Aplikasi secara otomatis menyisipkan 4 akun default beserta biodata awal setiap kali server Golang dinyalakan (`go run main.go`):

| Role | Email | Password | Hak Akses |
| :--- | :--- | :--- | :--- |
| **SUPERADMIN** | `faridwimansyah8@gmail.com` | `password123` | Hak Akses Penuh Sistem (CRUD User, Switch Role) |
| **OWNER** | `owner@demo.com` | `password123` | Akses Laporan & Manajemen User Read-only |
| **ADMIN** | `admin@demo.com` | `password123` | Akses Admin Standar Operasional |
| **USER** | `user@demo.com` | `password123` | **(BARU)** Akses Pengguna Biasa / Regular User |

---

## ⚡ Cara Kerja Auto-Migrate & Seeder

1. **Auto-Migrate GORM:** `config.InitDB()` memanggil `db.AutoMigrate(&entity.User{})`. GORM akan otomatis membuat tabel `users` di PostgreSQL dan mendeteksi perubahan kolom baru (`phone`, `gender`, `birth_date`, `address`, `bio`, `avatar`, `deleted_at`).
2. **Database Seeding:** Setelah migrasi skema tabel selesai, `config.SeedSuperAdmin(db)` otomatis memeriksa ketersediaan email seeder. Jika belum ada di database, keempat akun di atas akan disisipkan secara otomatis.
