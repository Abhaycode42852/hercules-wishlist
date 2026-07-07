<div align="center">

# Bond Wishlist Service

### A high-performance Go backend for managing bond wishlists.

Built using **Go**, **Gin**, and **PostgreSQL**.

</div>

A RESTful backend service built with Go, Gin, and PostgreSQL that allows users to create wishlists and manage bonds inside them.

---

## Tech Stack

- Go
- Gin Framework
- PostgreSQL

---

## Project Structure

```text
wishlist-backend/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── config/
│   ├── handlers/
│   ├── models/
│   ├── repository/
│   ├── routes/
│   └── services/
│
├── .env.example
├── .gitignore
├── setup.sql
├── go.mod
└── README.md
```

---
# Architecture

```text
HTTP Request
     │
     ▼
 Gin Handler
     │
     ▼
 Service Layer
     │
     ▼
 Repository Layer
     │
     ▼
 PostgreSQL
```
---

## Setup Instructions

### 1. Clone Repository

```bash
git clone https://github.com/Abhaycode42852/wishlist-backend.git
cd wishlist-backend
```

### 2. Database Setup

# PostgreSQL Installation & Setup Guide

This project uses **PostgreSQL** as its database. Follow the steps below to install PostgreSQL, create the database, and configure the application.

---

# 1. Download PostgreSQL

Visit the official PostgreSQL website:

🔗 https://www.postgresql.org/download/

Select your operating system:

- Windows
- macOS
- Linux

For Windows users, click:

```text
Download the installer certified by EDB
```

This will redirect you to:

🔗 https://www.enterprisedb.com/downloads/postgres-postgresql-downloads

---

# 2. Install PostgreSQL

Run the downloaded installer.

During installation:

### Components

Keep the default selections:

```text
✓ PostgreSQL Server
✓ pgAdmin 4
✓ Stack Builder
✓ Command Line Tools
```

### Installation Directory

Use the default location unless you have a specific preference.

### Set Database Password

You will be asked to create a password for the PostgreSQL superuser (`postgres`).

Example:

```text
Username: postgres
Password: your_password
```

⚠️ Remember this password. It will be required later in the `.env` file.

### Port

Keep the default PostgreSQL port:

```text
5432
```

### Locale

Use the default locale and continue.

Finish the installation.

---

# 3. Verify Installation

Open:

```text
pgAdmin 4
```

You should see:

```text
Servers
└── PostgreSQL
```

If pgAdmin opens successfully, PostgreSQL has been installed correctly.

---

# 4. Create the Database

Open pgAdmin.

Right-click:

```text
Databases
```

Select:

```text
Create → Database
```

Enter:

```text
Database Name: wishlist_db
```

Click:

```text
Save
```

The database should now appear under:

```text
Databases
└── wishlist_db
```

---

# 5. Execute Database Schema

Open:

```text
wishlist_db
```

Select:

```text
Tools → Query Tool
```

Open the project's:

```text
setup.sql
```

Execute the script.

You should see tables similar to:

```text
public
├── bonds
├── wishlists
└── wishlist_bonds
```

---

### 3. Configure Environment Variables

Create:

```text
.env
```

using:

```text
.env.example
```

Example:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=wishlist_db
```

### 4. Install Dependencies

```bash
go mod tidy
```

### 6. Run Application

```bash
go run cmd/main.go
```

Server starts on:

```text
http://localhost:8080
```
---

## 5. Add Database Schema Section


# Database Schema

## Bonds

```text
bonds
├── b_id (UUID)
├── name
├── yield
├── coupon_rate
├── rating
├── frequency
├── min_investment
├── min_units
├── max_units
├── face_value
├── maturity_date
├── issuer
├── sector
├── logo_url
└── isin
```
## Wishlists
```text
wishlists
├── w_id (UUID)
├── name
├── size
└── created_at
```
## Wishlist_Bonds
```text
wishlist_bonds
├── w_id
└── b_id
```

---

## Business Constraints

### Wishlists

- Maximum 5 wishlists allowed.
- Wishlist names must be unique.
- Wishlist names cannot be empty.
- Wishlist names cannot exceed 30 characters.

### Bonds

- Maximum 10 bonds per wishlist.
- Duplicate bonds are not allowed in the same wishlist.

---

## API Endpoints

### Get All Bonds

```http
GET /api/v1/all-bonds
```

Query Parameters:

| Parameter | Description |
|------------|------------|
| page | Page number |
| limit | Number of records |
| sort | name, yield, frequency, rating, min_units, tenure |
| order | asc, desc |

Example:

```http
GET /api/v1/all-bonds?page=1&limit=10&sort=yield&order=desc
```

---

### Create Wishlist

```http
POST /api/v1/wishlist
```

Body:

```json
{
  "name": "Retirement Bonds"
}
```

---

### Get All Wishlists

```http
GET /api/v1/wishlist
```

---

### Get Wishlist Details

Returns wishlist information and all bonds belonging to the wishlist.

```http
GET /api/v1/wishlist/:id
```

---

### Update Wishlist

```http
PUT /api/v1/wishlist/:id
```

Body:

```json
{
  "name": "Long Term Bonds"
}
```

---

### Delete Wishlist

```http
DELETE /api/v1/wishlist/:id
```

---

### Add Bond To Wishlist

```http
POST /api/v1/wishlist/:id/add-bonds
```

Body:

```json
{
  "bond_id": "<bond-uuid>"
}
```

---

### Remove Bond From Wishlist

```http
DELETE /api/v1/wishlist/:id/bonds/:bondId
```
---
## Sample Responses

```md
GET /api/v1/all-bonds
```

```json
{
  "page": 1,
  "limit": 10,
  "total": 5,
  "data": [
    {
      "b_id": "...",
      "name": "Government Bond",
      "yield": 7.5,
      "coupon_rate": 7.0,
      "rating": "AAA",
      "sector": "Government",
      "face_value": 1000,
      "logo_url": "..."
    }
  ]
}
```
```md
GET /api/v1/wishlist
```

```json
[{
  "w_id": <UUID>,
  "name": "Wishlist",
  "size": 5
    },
{
  "w_id": <UUID>,
  "name": "Wishlist2",
  "size": 9
    }
  ]

```
---

---

## Features

- RESTful API design
- Layered Architecture
  - Handler
  - Service
  - Repository
- PostgreSQL persistence
- Pagination
- Sorting
- Transactions
- Validation
- Constraint enforcement

---

## Future Improvements

- Authentication & Authorization
- Database migrations
- Docker support
