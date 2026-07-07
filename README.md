# Wishlist Backend Service

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

## Setup Instructions

### 1. Clone Repository

```bash
git clone <repository-url>
cd wishlist-backend
```

### 2. Create Database

Create a PostgreSQL database named:

```text
wishlist_db
```

### 3. Run Database Script

Execute:

```text
setup.sql
```

using pgAdmin or psql.

### 4. Configure Environment Variables

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

### 5. Install Dependencies

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