# Kasir API

API backend untuk aplikasi kasir (Point of Sale) yang dibangun menggunakan Go dan PostgreSQL.

> 📖 **Dokumentasi API Lengkap**: Lihat [API_DOCUMENTATION.md](API_DOCUMENTATION.md) untuk dokumentasi endpoint yang lebih detail dengan contoh request/response lengkap.

> 📦 **Postman Collection**: Import file [Kasir_API.postman_collection.json](Kasir_API.postman_collection.json) ke Postman untuk testing API.

## 📋 Daftar Isi

- [Fitur](#fitur)
- [Teknologi](#teknologi)
- [Prasyarat](#prasyarat)
- [Instalasi](#instalasi)
- [Konfigurasi](#konfigurasi)
- [Menjalankan Aplikasi](#menjalankan-aplikasi)
- [Testing](#testing)
- [API Endpoints](#api-endpoints)
- [Struktur Database](#struktur-database)
- [Struktur Project](#struktur-project)

## ✨ Fitur

- ✅ Manajemen Kategori Produk (CRUD)
- ✅ Manajemen Produk (CRUD)
- ✅ Koneksi Database PostgreSQL dengan Supabase
- ✅ Auto Migration Database
- ✅ Health Check Endpoint
- ✅ Environment Configuration dengan Viper
- ✅ Clean Architecture (Handler → Service → Repository)

## 🛠 Teknologi

- **Language**: Go 1.25.5
- **Database**: PostgreSQL (Supabase)
- **Libraries**:
  - `github.com/jackc/pgx/v5` - PostgreSQL driver (modern, high-performance)
  - `github.com/spf13/viper` - Configuration management
- **Architecture**: Clean Architecture Pattern

## 📦 Prasyarat

Sebelum menjalankan aplikasi, pastikan Anda sudah menginstal:

- Go 1.25.5 atau lebih baru
- PostgreSQL (atau akses ke Supabase)
- Git

## 🚀 Instalasi

1. Clone repository ini:
```bash
git clone <repository-url>
cd kasir-api
```

2. Install dependencies:
```bash
go mod download
```

3. Copy file `.env.example` menjadi `.env`:
```bash
copy .env.example .env
```

4. Edit file `.env` dan sesuaikan konfigurasi database Anda.

## ⚙️ Konfigurasi

### Environment Variables

Buat file `.env` di root directory dengan konfigurasi berikut:

```env
PORT=8080
DB_CONN=postgresql://username:password@host:port/database?options
```

### Contoh Konfigurasi Database

#### Supabase (Production) - Recommended
```env
DB_CONN=postgres://postgres.xxx:password@aws-1-ap-south-1.pooler.supabase.com:6543/postgres
```

#### Supabase (Direct Connection)
```env
DB_CONN=postgres://postgres.xxx:password@aws-1-ap-south-1.pooler.supabase.com:5432/postgres
```

#### PostgreSQL Lokal
```env
DB_CONN=postgres://postgres:password@localhost:5432/kasir_db
```

### Penjelasan Connection String

Format: `postgres://[user]:[password]@[host]:[port]/[database]`

- **user**: Username database
- **password**: Password database
- **host**: Host database (localhost atau remote)
- **port**: Port database (default PostgreSQL: 5432, Supabase Pooler: 6543)
- **database**: Nama database
- **options**: Parameter tambahan (contoh: `sslmode=disable`, `pgbouncer=true`)

## 🏃 Menjalankan Aplikasi

### Development Mode

```bash
go run main.go
```

### Build dan Run

```bash
# Build executable
go build -o kasir-api.exe

# Jalankan executable
./kasir-api.exe
```

### Menggunakan Helper Scripts (Recommended)

Untuk memudahkan development, gunakan script helper yang sudah disediakan:

**Start Server (dengan auto-stop server lama):**
```bash
start_server.bat
```

Script ini akan:
- ✅ Mengecek apakah port 8080 sudah digunakan
- ✅ Menghentikan proses lama secara otomatis
- ✅ Menjalankan server baru

**Stop Server:**
```bash
stop_server.bat
```

Script ini akan menghentikan semua proses yang menggunakan port 8080.

Aplikasi akan berjalan di `http://localhost:8080` (atau port yang Anda konfigurasi).

## 🧪 Testing

### Test Koneksi Database

Jalankan script test koneksi database menggunakan PGX driver:

```bash
go run test_pgx.go
```

Script ini akan:
- ✅ Memuat konfigurasi dari `.env`
- ✅ Mencoba koneksi ke database
- ✅ Menampilkan informasi database (version, user, schema)
- ✅ Menjalankan query test
- ✅ Memeriksa tabel yang ada dan jumlah rows

**Output yang diharapkan:**
```
===========================================
  KASIR API - PGX DATABASE TEST
===========================================

📋 Step 1: Loading configuration from .env file...
✅ Configuration loaded successfully!

🔌 Step 2: Testing database connection with PGX...
✅ Database connection successful!

🔍 Step 3: Testing query - Get PostgreSQL version...
   ✅ PostgreSQL Version: PostgreSQL 17.6...

ℹ️  Step 4: Retrieving database information...
   - Current Database: postgres
   - Current User: postgres
   - Current Schema: public

🧪 Step 5: Testing simple calculation query...
   ✅ Test query executed successfully (1 + 1 = 2)

📊 Step 6: Checking existing tables in database...
   ✅ Found 2 table(s):
      - categories (2 rows)
      - products (2 rows)

===========================================
  ✅ ALL TESTS COMPLETED SUCCESSFULLY!
===========================================

🎉 Your database connection is working perfectly!
```

### Test API Endpoints

#### Health Check
```bash
curl http://localhost:8080/health
```

Response:
```json
{
  "status": "OK",
  "message": "API Running"
}
```

#### Get All Categories
```bash
curl http://localhost:8080/api/categories
```

#### Get All Products
```bash
curl http://localhost:8080/api/produk
```

## 📡 API Endpoints

Base URL: `http://localhost:8080`

---

### 🏠 System Endpoints

#### 1. Home Endpoint
**Endpoint:** `GET /`

**Deskripsi:** Endpoint untuk mengecek apakah server berjalan.

**Request:**
```bash
curl http://localhost:8080/
```

**Response:**
```
Running on port :8080
```

**Status Code:** `200 OK`

---

#### 2. Health Check
**Endpoint:** `GET /health`

**Deskripsi:** Endpoint untuk health check API.

**Request:**
```bash
curl http://localhost:8080/health
```

**Response:**
```json
{
  "status": "OK",
  "message": "API Running"
}
```

**Status Code:** `200 OK`

---

### 📂 Categories Endpoints

#### 1. Get All Categories
**Endpoint:** `GET /api/categories`

**Deskripsi:** Mendapatkan semua data kategori.

**Request:**
```bash
curl http://localhost:8080/api/categories
```

**Response Success:**
```json
[
  {
    "id": 1,
    "name": "Makanan",
    "description": "Kategori makanan dan snack"
  },
  {
    "id": 2,
    "name": "Minuman",
    "description": "Kategori minuman"
  }
]
```

**Status Code:** `200 OK`

**Response Empty:**
```json
[]
```

---

#### 2. Get Category by ID
**Endpoint:** `GET /api/categories/{id}`

**Deskripsi:** Mendapatkan detail kategori berdasarkan ID.

**Request:**
```bash
curl http://localhost:8080/api/categories/1
```

**Response Success:**
```json
{
  "id": 1,
  "name": "Makanan",
  "description": "Kategori makanan dan snack"
}
```

**Status Code:** `200 OK`

**Response Error (Not Found):**
```json
{
  "error": "Category not found"
}
```

**Status Code:** `404 Not Found`

---

#### 3. Create New Category
**Endpoint:** `POST /api/categories`

**Deskripsi:** Membuat kategori baru.

**Request Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Elektronik",
  "description": "Kategori produk elektronik"
}
```

**Request Example:**
```bash
curl -X POST http://localhost:8080/api/categories \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Elektronik",
    "description": "Kategori produk elektronik"
  }'
```

**Response Success:**
```json
{
  "id": 3,
  "name": "Elektronik",
  "description": "Kategori produk elektronik"
}
```

**Status Code:** `201 Created`

**Response Error (Invalid Input):**
```json
{
  "error": "Name is required"
}
```

**Status Code:** `400 Bad Request`

**Validasi:**
- `name` (required): Nama kategori, tidak boleh kosong
- `description` (optional): Deskripsi kategori

---

#### 4. Update Category
**Endpoint:** `PUT /api/categories/{id}`

**Deskripsi:** Mengupdate data kategori berdasarkan ID.

**Request Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Elektronik Updated",
  "description": "Kategori produk elektronik dan gadget"
}
```

**Request Example:**
```bash
curl -X PUT http://localhost:8080/api/categories/3 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Elektronik Updated",
    "description": "Kategori produk elektronik dan gadget"
  }'
```

**Response Success:**
```json
{
  "id": 3,
  "name": "Elektronik Updated",
  "description": "Kategori produk elektronik dan gadget"
}
```

**Status Code:** `200 OK`

**Response Error (Not Found):**
```json
{
  "error": "Category not found"
}
```

**Status Code:** `404 Not Found`

---

#### 5. Delete Category
**Endpoint:** `DELETE /api/categories/{id}`

**Deskripsi:** Menghapus kategori berdasarkan ID.

**Request:**
```bash
curl -X DELETE http://localhost:8080/api/categories/3
```

**Response Success:**
```json
{
  "message": "Category deleted successfully"
}
```

**Status Code:** `200 OK`

**Response Error (Not Found):**
```json
{
  "error": "Category not found"
}
```

**Status Code:** `404 Not Found`

**⚠️ Catatan:** 
- Jika kategori memiliki produk yang terkait, `category_id` pada produk tersebut akan di-set menjadi `NULL` (ON DELETE SET NULL)

---

### 🛍️ Products Endpoints

#### 1. Get All Products
**Endpoint:** `GET /api/produk`

**Deskripsi:** Mendapatkan semua data produk.

**Request:**
```bash
curl http://localhost:8080/api/produk
```

**Response Success:**
```json
[
  {
    "id": 1,
    "name": "Nasi Goreng",
    "price": 15000,
    "stock": 50,
    "category_id": 1
  },
  {
    "id": 2,
    "name": "Es Teh Manis",
    "price": 5000,
    "stock": 100,
    "category_id": 2
  }
]
```

**Status Code:** `200 OK`

**Response Empty:**
```json
[]
```

---

#### 2. Get Product by ID
**Endpoint:** `GET /api/produk/{id}`

**Deskripsi:** Mendapatkan detail produk berdasarkan ID.

**Request:**
```bash
curl http://localhost:8080/api/produk/1
```

**Response Success:**
```json
{
  "id": 1,
  "name": "Nasi Goreng",
  "price": 15000,
  "stock": 50,
  "category_id": 1
}
```

**Status Code:** `200 OK`

**Response Error (Not Found):**
```json
{
  "error": "Product not found"
}
```

**Status Code:** `404 Not Found`

---

#### 3. Create New Product
**Endpoint:** `POST /api/produk`

**Deskripsi:** Membuat produk baru.

**Request Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Mie Goreng",
  "price": 12000,
  "stock": 30,
  "category_id": 1
}
```

**Request Example:**
```bash
curl -X POST http://localhost:8080/api/produk \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Mie Goreng",
    "price": 12000,
    "stock": 30,
    "category_id": 1
  }'
```

**Response Success:**
```json
{
  "id": 3,
  "name": "Mie Goreng",
  "price": 12000,
  "stock": 30,
  "category_id": 1
}
```

**Status Code:** `201 Created`

**Response Error (Invalid Input):**
```json
{
  "error": "Name, price, and stock are required"
}
```

**Status Code:** `400 Bad Request`

**Validasi:**
- `name` (required): Nama produk, tidak boleh kosong
- `price` (required): Harga produk, harus berupa angka positif
- `stock` (required): Stok produk, harus berupa angka positif atau 0
- `category_id` (optional): ID kategori, harus valid (ada di tabel categories)

---

#### 4. Update Product
**Endpoint:** `PUT /api/produk/{id}`

**Deskripsi:** Mengupdate data produk berdasarkan ID.

**Request Headers:**
```
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Mie Goreng Spesial",
  "price": 15000,
  "stock": 25,
  "category_id": 1
}
```

**Request Example:**
```bash
curl -X PUT http://localhost:8080/api/produk/3 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Mie Goreng Spesial",
    "price": 15000,
    "stock": 25,
    "category_id": 1
  }'
```

**Response Success:**
```json
{
  "id": 3,
  "name": "Mie Goreng Spesial",
  "price": 15000,
  "stock": 25,
  "category_id": 1
}
```

**Status Code:** `200 OK`

**Response Error (Not Found):**
```json
{
  "error": "Product not found"
}
```

**Status Code:** `404 Not Found`

---

#### 5. Delete Product
**Endpoint:** `DELETE /api/produk/{id}`

**Deskripsi:** Menghapus produk berdasarkan ID.

**Request:**
```bash
curl -X DELETE http://localhost:8080/api/produk/3
```

**Response Success:**
```json
{
  "message": "Product deleted successfully"
}
```

**Status Code:** `200 OK`

**Response Error (Not Found):**
```json
{
  "error": "Product not found"
}
```

**Status Code:** `404 Not Found`

---

### 📊 HTTP Status Codes

API ini menggunakan status code HTTP standar:

| Status Code | Deskripsi |
|-------------|-----------|
| `200 OK` | Request berhasil |
| `201 Created` | Resource berhasil dibuat |
| `400 Bad Request` | Request tidak valid (validasi error) |
| `404 Not Found` | Resource tidak ditemukan |
| `405 Method Not Allowed` | HTTP method tidak diizinkan |
| `500 Internal Server Error` | Error pada server |

---

### 🔍 Testing Endpoints dengan Postman

1. **Import Collection**: Buat collection baru di Postman
2. **Set Base URL**: `http://localhost:8080`
3. **Test Endpoints**: Gunakan contoh request di atas

**Contoh Testing Flow:**
```
1. GET /health                    → Cek API berjalan
2. GET /api/categories            → Lihat semua kategori
3. POST /api/categories           → Buat kategori baru
4. GET /api/categories/1          → Lihat detail kategori
5. POST /api/produk               → Buat produk baru
6. GET /api/produk                → Lihat semua produk
7. PUT /api/produk/1              → Update produk
8. DELETE /api/produk/1           → Hapus produk
```

## 🗄️ Struktur Database

### Table: categories

| Column | Type | Constraint |
|--------|------|------------|
| id | SERIAL | PRIMARY KEY |
| name | VARCHAR(255) | NOT NULL |
| description | TEXT | - |

### Table: products

| Column | Type | Constraint |
|--------|------|------------|
| id | SERIAL | PRIMARY KEY |
| name | VARCHAR(255) | NOT NULL |
| price | INT | NOT NULL |
| stock | INT | NOT NULL |
| category_id | INT | FOREIGN KEY → categories(id) |

## 📁 Struktur Project

```
kasir-api/
├── database/
│   └── database.go          # Database initialization & migration
├── handlers/
│   ├── category_handler.go  # HTTP handlers untuk categories
│   └── product_handler.go   # HTTP handlers untuk products
├── models/
│   ├── category.go          # Model Category
│   └── product.go           # Model Product
├── repositories/
│   ├── category_repository.go  # Database operations untuk categories
│   └── product_repository.go   # Database operations untuk products
├── response/
│   └── response.go          # Response helper
├── services/
│   ├── category_service.go  # Business logic untuk categories
│   └── product_service.go   # Business logic untuk products
├── .env                     # Environment configuration (tidak di-commit)
├── .env.example             # Template environment configuration
├── .gitignore              # Git ignore rules
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksums
├── main.go                 # Application entry point
├── start_server.bat        # Helper script untuk start server
├── stop_server.bat         # Helper script untuk stop server
└── README.md               # Dokumentasi ini
```

## 🏗️ Architecture Pattern

Project ini menggunakan **Clean Architecture** dengan layer:

1. **Handler Layer** (`handlers/`)
   - Menerima HTTP request
   - Validasi input
   - Memanggil service layer
   - Mengembalikan HTTP response

2. **Service Layer** (`services/`)
   - Business logic
   - Validasi business rules
   - Memanggil repository layer

3. **Repository Layer** (`repositories/`)
   - Database operations (CRUD)
   - Query execution
   - Data mapping

4. **Model Layer** (`models/`)
   - Data structures
   - Domain entities

## 🔧 Troubleshooting

### Database Connection Error

**Problem**: `Failed to initialize database`

**Solution**:
1. Pastikan connection string di `.env` sudah benar
2. Cek apakah database server berjalan
3. Verifikasi username, password, dan nama database
4. Untuk Supabase, pastikan menggunakan port 6543 dengan `pgbouncer=true`

### Migration Error

**Problem**: `Migration warning`

**Solution**:
1. Pastikan database user memiliki permission untuk CREATE TABLE
2. Cek apakah tabel sudah ada sebelumnya
3. Jalankan migration manual jika diperlukan

### Port Already in Use

**Problem**: `bind: Only one usage of each socket address (protocol/network address/port) is normally permitted`

**Penyebab**: Port 8080 sudah digunakan oleh aplikasi lain atau instance aplikasi yang masih berjalan.

**Solution 1: Hentikan Proses yang Menggunakan Port**

1. Cek proses yang menggunakan port 8080:
```bash
netstat -ano | findstr :8080
```

Output akan menampilkan:
```
TCP    0.0.0.0:8080           0.0.0.0:0              LISTENING       21764
```

2. Catat PID (Process ID) dari output di atas (contoh: 21764)

3. Hentikan proses tersebut:
```bash
taskkill /F /PID 21764
```

4. Jalankan aplikasi kembali:
```bash
go run main.go
```

**Solution 2: Ubah Port Aplikasi**

1. Edit file `.env`:
```env
PORT=8081
```

2. Jalankan aplikasi:
```bash
go run main.go
```

3. Aplikasi akan berjalan di port 8081:
```
http://localhost:8081
```

**Solution 3: Gunakan Script Helper**

Buat file `stop_server.bat`:
```batch
@echo off
for /f "tokens=5" %%a in ('netstat -ano ^| findstr :8080') do (
    taskkill /F /PID %%a
)
echo Server stopped!
```

Jalankan script sebelum start server:
```bash
stop_server.bat
go run main.go
```

## 📝 Development Guidelines

### Menambah Endpoint Baru

1. Buat model di `models/`
2. Buat repository di `repositories/`
3. Buat service di `services/`
4. Buat handler di `handlers/`
5. Register route di `main.go`
6. Update migration di `database/database.go` jika perlu tabel baru

### Code Style

- Gunakan `gofmt` untuk formatting
- Follow Go naming conventions
- Tambahkan comment untuk exported functions
- Error handling yang proper

## 📄 License

[Tentukan license Anda di sini]

## 👥 Contributors

[Daftar kontributor]

## 📞 Contact

Untuk pertanyaan atau dukungan, silakan hubungi [contact information]

---

**Happy Coding! 🚀**
