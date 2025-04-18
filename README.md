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

### Setup Database
Pastikan PostgreSQL sudah berjalan dan environment di .env sudah sesuai. Contoh variabel penting:
```bash
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=tokokita_db
```

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