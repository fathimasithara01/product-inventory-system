Product Inventory System – Backend

A backend service built with Golang + Gin + GORM to manage products, variants, sub-variants, and stock transactions with full transactional safety.

🚀 Tech Stack

Language: Go (Golang)

Framework: Gin

ORM: GORM

Database: PostgreSQL

Decimal Handling: shopspring/decimal

UUID: google/uuid

📁 Project Structure
backend/
├── cmd/
│   └── server/
│       └── main.go
│── config/
│   └── config.go
│
│── database/
│   └── migrate.go
│   └── postgres.go
│
├── internal/
│   ├── handler/
│   │   ├── product_handler.go
│   │   ├── stock_handler.go
│   │   └── report_handler.go
│   ├── service/
│   │   ├── product_service.go
│   │   └── stock_service.go
│   ├── repository/
│   │   ├── product_repo.go
│   │   ├── variant_repo.go
│   │   ├── subvariant_repo.go
│   │   └── stock_repo.go
│   ├── model/
│   │   ├── product.go
│   │   ├── variant.go
│   │   ├── subvariant.go
│   │   └── stock_transaction.go
│   ├── dto/
│   │    └── product_dto.go
│   ├── middleware/
│   │    ├── logger.go
│   │    ├── recover.go
│   ├── utils/
│       ├── pagination.go
│       ├── response.go
│       └── validator.go
├── migrations/
│   └── 00001_create_products.sql
│   └── 00002_create_variants.sql
│   └── 00003_create_variant_options.sql
│   └── 00004_create_sub_variants.sql
│   └── 00005_create_stock_transactions.sql
│
├──.env
│
├── go.mod
├── go.sum
└── README.md

⚙️ Environment Variables

Create a .env file or export variables:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=inventory_db

▶️ Run the Server
go mod tidy
go run cmd/server/main.go


Server starts on:

http://localhost:8080

📦 API Endpoints
🔹 Product APIs
Create Product
POST /api/products


Request JSON

{
  "product_id": 1,
  "product_code": "P001",
  "product_name": "T-Shirt",
  "product_image": "https://example.com/tshirt.png",
  "created_user": "c8bcdcb4-2113-4f1f-b861-5f57fe0dcad7",
  "hsn_code": "HSN100",
  "variants": [
    { "name": "Size" },
    { "name": "Color" }
  ],
  "options": [
    { "variant_index": 0, "value": "S" },
    { "variant_index": 0, "value": "M" },
    { "variant_index": 0, "value": "L" },
    { "variant_index": 1, "value": "Red" },
    { "variant_index": 1, "value": "Blue" }
  ],
  "sub_variants": [
    {
      "option_indices": [0, 3],
      "sku": "TS-S-RED",
      "stock": "50"
    }
  ]
}

List Products
GET /api/products?page=1&limit=10

🔹 Stock APIs
Stock IN
POST /api/stock/in

{
  "product_id": "2eff1762-da7b-4eeb-96dc-022eebd0a4ca",
  "sub_variant_id": "e0848ff9-7cc7-4174-8ba0-4e57835b0354",
  "quantity": "10"
}

Stock OUT
POST /api/stock/out

{
  "product_id": "2eff1762-da7b-4eeb-96dc-022eebd0a4ca",
  "sub_variant_id": "e0848ff9-7cc7-4174-8ba0-4e57835b0354",
  "quantity": "5"
}

🔹 Stock Report
Get Stock Transactions
GET /api/stock/report?from=2026-01-01&to=2026-01-31

🧠 Business Logic Highlights

✅ Atomic transactions using DB transactions

✅ Row-level locking (SELECT FOR UPDATE) for stock safety

✅ Accurate total stock auto-calculated from sub-variants

✅ Variant & SubVariant flexible modeling

✅ Decimal-safe stock calculations

🔐 Data Integrity Rules

Sub-variant stock can never go below zero

Stock updates are concurrent-safe

Product total_stock always reflects sum of sub-variant stock

📊 Database Tables (High Level)

products

variants

variant_options

sub_variants

stock_transactions

🧪 Testing with Postman

Recommended order:

Create Product

Copy sub_variant_id

Stock IN

Stock OUT

Stock Report

📌 Future Enhancements

Pagination for stock report

Product-wise stock report

CSV / Excel export

Authentication (JWT)

Soft delete support

👩‍💻 Author

Fathima Sithara
Backend Developer (Golang | Distributed Systems)