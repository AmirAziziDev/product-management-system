# Product Management System — Agile-lite Roadmap

This document defines the **versioned, agile-lite implementation roadmap** for the Product Management System demo application.  
It showcases **Go + Gin + PostgreSQL + Vue/Vuetify** with clear vertical slices.

---

## Release Roadmap (v0.1.0 → v0.7.0)

Each version increment is a *vertical slice* (DB → API → UI).  
Tags: `v0.1.0`, `v0.2.0`, … `v0.7.0`.

---

### v0.1.0 — Minimal list (id, code, name, description, created_at)

**Goal**  
Read-only product list. **Order is fixed** (newest first by `created_at DESC`).

**DB/Migrations**
- `products(id SERIAL PK, code INTEGER NOT NULL UNIQUE CHECK (code >= 0), name TEXT NOT NULL UNIQUE, description TEXT NULL, created_at TIMESTAMPTZ DEFAULT now())`
- Seed **50** products (some with `description`, some `NULL`).

**Comments (SKU segments)**
```sql
COMMENT ON COLUMN products.code IS
  'Stable business code (unsigned int). Used as the second part of SKU.';
```

**API**
- `GET /healthz`
- `GET /api/v1/products?page=1&page_size=20`  
  Response:
  ```json
  {
    "data": [
      { "id": 1, "code": 20, "name": "Alpha", "description": null, "created_at": "2025-08-24T08:00:00Z" }
    ],
    "meta": {
      "total": 50,
      "page": 1,
      "page_size": 20
    }
  }
  ```

**UI**
- `/products` table: **ID, Code, Name, Description, Created At**.
- Server-side pagination.
- **SKU display note:** when composing SKU parts later, format numeric codes with zero‑padding as needed (e.g., `020`).

---

### v0.2.0 — Add Product Type (code + segment prefix)

**Goal**  
Introduce product types and display on list.

**DB/Migrations**
- `product_types(id SERIAL PK, code INTEGER NOT NULL UNIQUE CHECK (code >= 0), name TEXT UNIQUE, created_at TIMESTAMPTZ)`
- `products.product_type_id INT NOT NULL REFERENCES product_types(id)`

**Comments (SKU segments)**
```sql
COMMENT ON COLUMN product_types.code IS
  'Stable business code (unsigned int). Used as the first part of SKU.';
```

**Seeding**
- Seed **~10 product types** and backfill `product_type_id` for existing products.

**API**
- `GET /api/v1/product-types`
- `GET /api/v1/products?page=1&page_size=20` includes product type info.

**UI**
- Add **Product Type (code + name)** column.
- Show **SKU preview** prefix: `product_type.code` + `.` + `products.code` (colors added v0.3.0).
- **SKU display note:** zero‑pad codes for consistent width (e.g., `804.020`).

---

### v0.3.0 — Add Colors (many-to-many) & full SKU

**Goal**  
Attach colors; display comma-separated + full SKU.

**DB/Migrations**
- `colors(id SERIAL PK, code INTEGER NOT NULL UNIQUE CHECK (code >= 0), name TEXT UNIQUE, hex CHAR(7) NOT NULL CHECK (hex ~ '^#[0-9A-Fa-f]{6}$'))`
- `products_colors(product_id INT REFERENCES products(id), color_id INT REFERENCES colors(id), PRIMARY KEY(product_id, color_id))`

**Comments (SKU segments)**
```sql
COMMENT ON COLUMN colors.code IS
  'Stable business code (unsigned int). Used as the third part of SKU.';
```

**Seeding**
- Seed **12 colors** with valid hex codes.

**API**
- `GET /api/v1/colors`
- `GET /api/v1/products?page=1&page_size=20` includes `colors: []` for each product.

**UI**
- Add **Colors** column (comma-separated names).
- **SKU** shown as `product_types_code.products_code.colors_code` (e.g., `804.020.023`). For multi-color products, show comma-separated SKUs.
- **SKU display note:** zero‑pad each numeric segment to your chosen width (e.g., 3 digits).

---

### v0.4.0 — Create Product (single-page + live SKU preview)

**Goal**  
Create products with type and colors. **SKU** is auto-generated from numeric codes (type.code + product.code + color.code).

**API**
- `POST /api/v1/products` with:
  ```json
  { "code": 20, "name": "Alpha", "description": "optional", "product_type_id": 804, "color_ids": [23, 24] }
  ```
- Transactional insert.

**UI**
- **Single-page form**: Name, Description (optional), Product Type (select), **Product Code (unsigned int)**, Colors (multi).
- **Live SKU preview** list: e.g., `804.020.023`, `804.020.024` (UI zero‑pads as configured).
- Primary action: **Create** → **confirm dialog** → submit.

---

### v0.5.0 — Filter by Product Type

**Goal**  
Filter list by product type.

**API**
- `GET /api/v1/products?product_type_id=...&page=...&page_size=...`

**UI**
- Type filter (select).
- Pagination based on `meta.total/page/page_size` (ordering still fixed by backend).

---

### v0.6.0 — Soft Delete + status filter

**Goal**  
Soft delete products; extend list filter with `status`.

**DB/Migrations**
- Add `deleted_at TIMESTAMPTZ NULL` to `products`.

**API**
- `DELETE /api/v1/products/{id}` → set `deleted_at`.
- **List with status filter** (replaces `include_deleted`):  
  `GET /api/v1/products?status=active|deleted|all&page=...&page_size=...`
    - default: `status=active`

**UI**
- Delete button with confirm.
- **Status** dropdown: *Active / Deleted / All*.
- “Deleted” badge for rows where `deleted_at` is set.

---

### v0.7.0 — Update Product (PATCH; name/description only)

**Goal**  
Partial update of product **without** changing SKU (codes), type, or colors.

**API**
- `PATCH /api/v1/products/{id}` with:
  ```json
  { "name": "New Name", "description": "Updated description" }
  ```
- Only `name` and `description` are updatable at this stage. `code`, `product_type_id`, and `color_ids` are immutable post-creation.

**UI**
- Edit dialog: fields for **Name** and **Description** only.

---
