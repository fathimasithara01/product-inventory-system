# Product Inventory System

A backend system for managing products, variants, sub-variants, and stock transactions for an online store.  
Built with **Golang**, **GORM**, **PostgreSQL**, and **Gin**.  

This project demonstrates **Clean Architecture**, **high-concurrency safe stock updates**, and **real-world backend patterns**.

---

## 🛠 Features

- Product CRUD with variants and sub-variants
- Stock management (IN / OUT) with **row-level locking**
- Stock reports with date filtering
- Pagination support for listing products
- Decimal-safe stock handling
- Clean, modular architecture (Handlers → Services → Repositories → Models)
- Middleware for logging & panic recovery
- JSON API responses standardized

---

## 📦 Tech Stack

- **Backend:** Golang, Gin, GORM
- **Database:** PostgreSQL
- **Middleware:** Logging, Recovery
- **Utils:** JSON response, pagination, validation
- **Models:** Product, Variant, VariantOption, SubVariant, StockTransaction
- **Version Control:** GitHub

---

## ⚙️ Prerequisites

- Golang >= 1.21
- PostgreSQL >= 13
- Git

---

## 🚀 Setup Instructions

1. **Clone repository**

```bash
git clone https://github.com/fathimasithara01/product-inventory.git
cd product-inventory/backend

2. **Create PostgreSQL database**

CREATE DATABASE product_inventory;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


3. **Set environment variables (or use .env)**

export DB_HOST=localhost
export DB_PORT=5432
export DB_USER=postgres
export DB_PASSWORD=postgres
export DB_NAME=product_inventory
export APP_PORT=8080


Run migrations

psql -U postgres -d product_inventory -f backend/migrations/001_init_tables.sql


Run server

cd cmd/server
go run main.go


Server will start on port 8080

🚀 Server running on port 8080

📌 API Endpoints
Products

POST /api/products → Create product with variants and sub-variants

GET /api/products?page=1&limit=10 → List products with pagination

Stock

POST /api/stock/in → Add stock to sub-variant

POST /api/stock/out → Remove stock from sub-variant

Reports

GET /api/reports/stock?from=YYYY-MM-DD&to=YYYY-MM-DD → Stock report

💻 JSON Examples

Create Product:

{
  "product": {
    "product_id": 101,
    "product_code": "PROD-101",
    "product_name": "T-Shirt",
    "product_image": "https://example.com/image.png",
    "created_user": "uuid-string"
  },
  "variants": [
    {"name": "Color"},
    {"name": "Size"}
  ],
  "options": [
    {"variant_id": "uuid-string", "value": "Red"},
    {"variant_id": "uuid-string", "value": "Blue"}
  ],
  "sub_variants": [
    {"option_ids": ["Red", "M"], "sku": "TSHIRT-RED-M", "stock": "50"}
  ]
}


Add Stock:

{
  "product_id": "uuid-string",
  "sub_variant_id": "uuid-string",
  "quantity": "10"
}


Remove Stock:

{
  "product_id": "uuid-string",
  "sub_variant_id": "uuid-string",
  "quantity": "5"
}

📌 Notes

Stock updates are transactional and thread-safe

Decimal is used for stock to avoid floating-point errors

All API responses follow the standard JSON format

✅ Author

Fathima Sithara

GitHub

LinkedIn