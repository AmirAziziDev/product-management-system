# Product Management System

A simple **Product Management System** with a **Vue + Vuetify** frontend and a **Go + Gin** backend. It lets you create products and list them with their types and colors.

---

## Contents

- [Overview](#overview)
- [Tech Stack](#tech-stack)
- [Quick Start](#quick-start)
- [Configuration](#configuration)
- [Running Locally (no Docker)](#running-locally-no-docker)
- [Docker](#docker)
- [Project Structure](#project-structure)
- [Scripts](#scripts)
- [API](#api)
    - [Products](#products)
    - [Product Types](#product-types)
    - [Colors](#colors)
    - [Health](#health)
    - [Error Format](#error-format)
- [Development Notes](#development-notes)
- [Troubleshooting](#troubleshooting)
- [License](#license)

---

## Overview

The system lets you:

- Create a product with a **code**, **name**, **type**, and **one or more colors** (description is optional).
- View a paginated list of products with their **type** and **colors**.
- Run everything with **Docker Compose** or run **frontend** and **backend** separately.

If you are here for the task description, see **[`task.md`](task.md)**.

---

## Tech Stack

- **Frontend:** Vue 3, Vite, Vuetify
- **Backend:** Go (Gin)
- **Database:** PostgreSQL
- **Container:** Docker & Docker Compose

---

## Quick Start

1) **Clone the repo**
```bash
git clone https://github.com/AmirAziziDev/product-management-system.git
cd product-management-system
```

2) **Create environment file**
```bash
cp .env.example .env
```

3) **Start with Docker (recommended)**
```bash
docker compose up --build
```

Services (default):
- Frontend: `http://localhost:3000`
- Backend API: `http://localhost:8080`
- PostgreSQL: `localhost:5432`

> If you prefer running without Docker, see the next sections.

---

## Configuration

Create a `.env` file in the project root. Minimum variables:

```env
POSTGRES_DB=product_management
POSTGRES_USER=postgres
POSTGRES_PASSWORD=password

# Optional overrides
BACKEND_PORT=8080
FRONTEND_PORT=3000
DB_HOST=localhost
DB_PORT=5432
```

> The Docker Compose file will read from `.env`.  
> When running locally without Docker, make sure PostgreSQL is running and these values match your setup.

---

## Running Locally (no Docker)

### 1) Database
- Install PostgreSQL.
- Create database `product_management` (or use your own name and update `.env`).

### 2) Backend
```bash
cd backend
go mod tidy
go run main.go
```
Default backend URL: `http://localhost:8080`

### 3) Frontend
```bash
cd frontend
npm install
npm run dev
```
Default frontend URL: `http://localhost:3000`

---

## Docker

1) Make sure you have **Docker** and **Docker Compose** installed.
2) From the project root:
```bash
docker compose up --build
```
3) Open:
- Frontend: `http://localhost:3000`
- Backend:  `http://localhost:8080`

To stop:
```bash
docker compose down
```

---

## Project Structure

```
product-management-system/
├─ backend/
│  ├─ handlers/            # HTTP handlers (Gin)
│  ├─ middleware/          # Validation, CORS, logging, etc.
│  ├─ models/              # Domain models
│  ├─ repositories/        # Data access (interfaces + SQL)
│  ├─ routes/              # Router setup
│  ├─ main.go
│  └─ go.mod / go.sum
├─ frontend/
│  ├─ src/                 # Vue 3 app
│  ├─ vite.config.ts       # Vite config
│  └─ package.json
├─ docker-compose.yml
├─ .env.example
├─ .env                    # (not committed)
├─ task.md
└─ README.md
```

---

## Scripts

### Frontend (Vue)
- `npm run dev` — Start development server
- `npm run build` — Build for production
- `npm run preview` — Preview production build

### Backend (Go)
- `go run main.go` — Start the backend server
- `go test ./...` — Run all tests
- `go build` — Build the application

---

## API

**Base URL:** `http://localhost:8080`

### Products

#### Create product
`POST /products`

**Request (JSON):**
```json
{
  "code": 101,
  "name": "Bookcase",
  "description": "Perfect for organizing books",
  "product_type_id": 1,
  "color_ids": [1, 3, 5]
}
```

**Responses:**
- `201 Created` with the created product
- `400 Bad Request` for validation problems
- `404 Not Found` when a referenced entity does not exist
- `409 Conflict` when `code` or `name` already exists

**Example 409 (unique violation):**
```json
{ "error": "code already exists" }
```

#### List products (paginated)
`GET /products?page=1&page_size=20`

**Response (JSON):**
```json
{
  "data": [
    {
      "id": 1,
      "code": 101,
      "name": "Bookcase",
      "description": "Perfect for organizing books",
      "product_type": {
        "id": 1,
        "code": 100,
        "name": "Storage & Organization",
        "created_at": "2025-08-25T15:33:08.919692Z"
      },
      "colors": [
        { "id": 1, "code": 10, "name": "White", "hex": "#FFFFFF", "created_at": "2025-08-25T15:33:08.919692Z" },
        { "id": 3, "code": 12, "name": "Oak", "hex": "#C3B091", "created_at": "2025-08-25T15:33:08.919692Z" }
      ],
      "created_at": "2025-08-25T15:33:08.919692Z"
    }
  ],
  "page": 1,
  "page_size": 20,
  "total": 42
}
```

### Product Types
`GET /product-types` — List all product types

**Response:**
```json
[
  { "id": 1, "code": 100, "name": "Storage & Organization", "created_at": "2025-08-25T15:33:08.919692Z" }
]
```

### Colors
`GET /colors` — List all colors

**Response:**
```json
[
  { "id": 1, "code": 10, "name": "White", "hex": "#FFFFFF", "created_at": "2025-08-25T15:33:08.919692Z" }
]
```

### Health
`GET /healthz` — Health check endpoint

**Response:**
```json
{
  "service": "product-management-api",
  "status": "healthy"
}
```

### Error Format

Use a **minimal and consistent** format:

- **Validation errors (per-field):**
```json
{
  "errors": {
    "product_type_id": "product type does not exist",
  }
}
```

> This format is easy to handle in Vue + Vuetify forms.  
> Show messages near the related inputs and a general message at the top if needed.

---

## Development Notes

- **Sorting:** Product list is ordered by `created_at DESC` in the backend.
- **Pagination:** Use `page` and `page_size` query params.
- **CORS:** Enabled to allow frontend requests during development.
- **Logging:** Backend logs requests and errors (Gin + zap).
- **Testing:** Run `go test ./...` inside `backend`.

---

## Troubleshooting

- **Ports already in use**: stop other apps or change `BACKEND_PORT` / `FRONTEND_PORT` in `.env`.
- **DB connection fails**: check `POSTGRES_*` variables and container logs.
- **Migrations**: if you added new tables, make sure the DB schema matches your code.
- **Frontend cannot reach backend**: verify CORS and that the API URL matches your frontend config.

---

