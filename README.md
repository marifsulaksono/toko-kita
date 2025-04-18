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

### Flow Bussiness
- Alur proses FIFO saat penjualan.

![Diagram Alur proses FIFO saat penjualan](https://uml.planttext.com/plantuml/png/NP5DImCn48Rl-HNlILhOW-SMjAKWYWSB5K4Gbkdkj3jkDjqcCuZ_tko77hnEtZplayacAKMrpZiLzOojjuXzl2HEU7XwUC61IB7dy6daAJWtvjtgbQDqXCXNuwJeVmDhjteZBpKbWOJU58kabnQLjTONkEetxd1ReFP-cRFRxYWUYJiZv5pLbaVIgi57W0Mr5gBSGJcs93fC5nOpixTODPIuBhFSzh3BA1UvPr8d9omcHkIE94WubtFh_HBQS6Qyqz81gsUB5CxoBpXUNW_dEzAC7sOrHS6VMSj8v3bDhkMWMDQi93wj7OTQJhCClXh9ErsCMFOslYIloXh9Bf8c4xTw5BLmXbwKF6wGfEifJ1t5N3NpYycYz1BQN9PNrhjDhXm-y0S0)

- Alur proses pencatatan pembelian dan penjualan.

![Diagram Alur proses pencatatan pembelian dan penjualan](https://uml.planttext.com/plantuml/png/TP71IWGn38RlVOgSbXNs2HHX5s7HamEA9s78JCDCd4odswO74T_TJ5t15RoqIV_axvTsdnMJbbcSiWh1GKg29YsPCZGqToIC0JOMXxU2Wi6vsk5Sj9MLb_2hxiC1N3zJuXaXqQbGLeW_wiqglg2mnyp08HQ5xKdVTp1Y6d07NkIg5Ztn7Crj8iYwz5FReVQZ0Rq6s7eAWsc9PkJ0OMZD0MuXbjIFcCclkRbyNWAUuVy1FAkNW7RrL4yESyS2bG7iMGeo71-EZrMxGyNgC9I45Mnj0s_1VTFy7GR5XjXyxwqdHPPAQZtHxIduzH--0000)

- Mekanisme pembuatan laporan laba bulanan.

![Diagram Mekanisme pembuatan laporan laba bulanan](https://uml.planttext.com/plantuml/png/NP31JiCm38RlUGgVTXhYlgbY824cs64WZaYLsvercfjaZdFea_10l1Zk2XhQAPBp-RF-VvCLHSl0JhqomJqNWRt4J1bscA9WiBT1U2YC0ODpvmMtUgDJeQpmzFWEjq96wjqGW_RLOYIKQkrMhuLaX8niheQamunoDDz7W6QomG8K-n8CVOu-m95ckEv8qNE6ReRdOOFQzhLN6lx-RM_hjOG3Q5JaUhvKmKb7-Fw2JCm-7EealgLhja_fGgOYRO-Phj4ayTEhSV_zDnU3aIyqDCX-YcK6prFGaX3OMpVdbCkPN_xQrbPWJ6v3ePjQ_m40)

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