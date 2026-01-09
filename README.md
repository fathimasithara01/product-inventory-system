Product Inventory System

A full-stack Product Inventory Management System with a Golang backend and a Svelte frontend.
The system manages products, variants, sub-variants, and stock transactions with transactional safety and concurrency control.

Tech Stack
Backend

Language: Go (Golang)

Framework: Gin

ORM: GORM

Database: PostgreSQL

UUID: google/uuid

Decimal Handling: shopspring/decimal

Frontend

Framework: Svelte

Build Tool: Vite

Language: JavaScript

Project Structure
product-inventory-system/
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   ├── migrate.go
│   │   └── postgres.go
│   ├── internal/
│   │   ├── handler/
│   │   ├── service/
│   │   ├── repository/
│   │   ├── model/
│   │   ├── dto/
│   │   ├── middleware/
│   │   └── utils/
│   ├── migrations/
│   ├── .env
│   ├── go.mod
│   └── go.sum
│
├── frontend/
│   ├── src/
│   ├── public/
│   ├── package.json
│   └── vite.config.js
│
├── postman/
│   └── ProductInventorySystem.postman_collection.json
│
└── README.md

Backend Setup
Environment Variables

Create a .env file inside backend/:

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=inventory_db

Run Backend Server
cd backend
go mod tidy
go run cmd/server/main.go


Backend runs on:

http://localhost:8080

Frontend Setup
cd frontend
npm install
npm run dev


Frontend runs on:

http://localhost:5173


Note:
The frontend is a basic implementation to demonstrate integration with the backend APIs.
Some UI refinements and validations are pending.

API Endpoints
Product APIs

Create Product
POST /api/products

List Products
GET /api/products?page=1&limit=10

Stock APIs

Stock IN
POST /api/stock/in

Stock OUT
POST /api/stock/out

Stock Report

Get Stock Transactions
GET /api/stock/report?from=YYYY-MM-DD&to=YYYY-MM-DD

Business Logic Highlights

Atomic database transactions

Row-level locking using SELECT FOR UPDATE

Concurrent-safe stock updates

Auto-calculated product total stock

Flexible variant and sub-variant modeling

Decimal-safe stock calculations

Data Integrity Rules

Sub-variant stock cannot go below zero

Stock updates are concurrency-safe

Product total stock always reflects sub-variant totals

Database Tables

products

variants

variant_options

sub_variants

stock_transactions

Testing with Postman

Recommended order:

Create Product

Copy sub_variant_id

Stock IN

Stock OUT

Stock Report

Postman collection:

postman/ProductInventorySystem.postman_collection.json

Future Enhancements

Pagination for stock reports

Product-wise stock reports

CSV / Excel export

Authentication (JWT)

Soft delete support

Author

Fathima Sithara
Backend Developer (Golang | Distributed Systems)
Kerala, India
GitHub: https://github.com/fathimasithara01

LinkedIn: https://www.linkedin.com/in/fathimasithara01/
