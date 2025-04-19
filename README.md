# 🏪 TokoKita - Backend FIFO Inventory & Sales Management

Sistem backend REST API untuk mengelola pembelian, penjualan, dan laporan laba menggunakan metode FIFO berbasis Golang + PostgreSQL.

---

## 🚀 Instruksi Menjalankan Server
### Persiapan Awal
1. Clone repository ini
2. Copy file `.env.example` menjadi `.env` dan sesuaikan konfigurasi database

```bash
cp .env.example .env
```

3. Buat database PostgreSQL sesuai konfigurasi .env
4. Jalankan seeder untuk membuat akun super admin:
- jika menggunakan make: ```make run-seed```
- jika tidak menggunakan make: ```go run cmd/api/main.go --seed```
5. Jalankan server API:
- jika menggunakan make: ```make run-api```
- jika tidak menggunakan make: ```go run cmd/api/main.go```

> [!NOTE]
> Migrasi database dilakukan secara otomatis menggunakan AutoMigrate dari GORM, cukup sesuaikan model yang ingin dimigrasi pada fungsi AutoMigrare di internal/contract/common/common.go

### Setup Database
Pastikan PostgreSQL sudah berjalan dan environment di .env sudah sesuai. Contoh variabel penting:
```bash
# konfigurasi variabel untuk koneksi ke database
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=tokokita_db

# data seeder
ROLE_NAME_SEEDER="Super Admin"
USER_NAME_SEEDER="Muhammad Arif Sulaksono"
USER_EMAIL_SEEDER="marifsulaksono@gmail.com"
USER_PASSWORD_SEEDER="Password123"
```

### Database Design
- Entity Relation Diagram

![ERD](https://drive.google.com/file/d/1hXQoQ-dWClWyUUfDAf0aR6wTgRD3oyxs/view?usp=sharing)

### Flow Bussiness
- Alur proses FIFO saat penjualan.

![Diagram Alur proses FIFO saat penjualan](https://drive.google.com/file/d/1G30LYELVy49X-3OoGg4isV0YmOv5VACu/view?usp=sharing)

- Alur proses pencatatan pembelian dan penjualan.

![Diagram Alur proses pencatatan pembelian dan penjualan](https://drive.google.com/file/d/1LZC1eZY5rX7RbSwC8seDlFHC3WLmDqoW/view?usp=sharing)

- Mekanisme pembuatan laporan laba bulanan.

![Diagram Mekanisme pembuatan laporan laba bulanan](https://drive.google.com/file/d/1WC7_hZ8EEiUemvJDlsXGgVFpIimzeGpk/view?usp=sharing)

### Testing Endpoint
Gunakan dokumentasi Postman berikut untuk mencoba semua endpoint API: [👉 Dokumentasi Postman](https://documenter.getpostman.com/view/30332593/2sB2cd5JDR)

### Tech & Stack
- Go 1.21+
- PostgreSQL
- GORM
- JWT
- Redis
- Makefile untuk kemudahan dev (opsional)

### More Information
Contact me: [marifsulaksono@gmail.com](mailto:marifsulaksono@gmail.com)