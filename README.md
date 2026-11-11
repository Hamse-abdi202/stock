 Inventory Management System

A RESTful Inventory Management System built with **Golang, Gin, GORM, and PostgreSQL**. It helps small businesses manage products, categories, stock movements, users, and inventory analytics.

==> Features

### Authentication & Authorization
- User registration and login
- JWT authentication
- Password hashing
- Role-Based Access Control (RBAC)
- Admin and Employee roles
- Token refresh and logout

### User Management
- Create and manage users
- Retrieve user information
- Role management
- Admin-only user operations

### Category Management
- Create categories
- Get all categories
- Get category by ID
- Update categories
- Delete categories

### Product Management
- Create products
- Get all products
- Get product by ID
- Update products
- Delete products
- Manage product prices and quantities
- Low-stock threshold management

### Stock Management
- Stock IN operations
- Stock OUT operations
- Track stock movements
- Record movement notes
- Track users responsible for stock changes

### Dashboard & Analytics
- Total products
- Total stock quantity
- Low-stock products
- Recent stock movements
- Total inventory value

## Technology Stack

 ✔️Golang | Backend programming language |
 ✔️Technology | Purpose |
 ✔️Gin | HTTP web framework |
 ✔️GORM | ORM |
 ✔️PostgreSQL | Database |
 ✔️JWT | Authentication |
 ✔️bcrypt | Password hashing |
 ✔️REST API | API architecture |

## Project Structure

InventoryManagementSystem/
│
├── config/
├── constants/
├── controllers/
├── dtos/
├── helpers/
├── middleware/
├── models/
├── repositories/
├── routes/
├── services/
├── main.go
├── go.mod
└── go.sum

## Getting Started

### 1. Clone Repository

git clone https://github.com/Hamse-abdi202/stock.git

### 2. Navigate to Project

cd stock

### 3. Install Dependencies

go mod tidy
 4. Configure Environment Variables
Create a `.env` file in the project root:

PORT=5666

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=inventory_db
DB_SSLMODE=disable

JWT_SECRET=your_secret_key
```

Configure the database connection according to your application configuration.

### 5. Run the Application

go run main.go

The backend will run at:

http://localhost:5666

## API Endpoints

### Authentication

| Method | Endpoint | Description |
|---|---|---|
| POST | /api/users | Register user |
| POST | /api/users/login | Login |
| GET | /api/users/whoami | Get authenticated user |
| POST | /api/users/refresh-token | Refresh access token |

### Categories

| Method | Endpoint | Description |
|---|---|---|
| POST | /api/categories | Create category |
| GET | /api/categories | Get all categories |
| GET | /api/categories/:id | Get category |
| PUT | /api/categories/:id | Update category |
| DELETE | /api/categories/:id | Delete category |

### Products

| Method | Endpoint | Description |
|---|---|---|
| POST | /api/products | Create product |
| GET | /api/products | Get all products |
| GET | /api/products/:id | Get product |
| PUT | /api/products/:id | Update product |
| DELETE | /api/products/:id | Delete product |

*Stock movement and dashboard routes depend on the route configuration in your application.*

## User Roles

**ADMIN**
- Full product and category management
- Manage users
- View dashboard and analytics
- Manage stock movements

**EMPLOYEE**
- View products
- Update stock quantities
- Record stock movements
- No product deletion permissions

## Database

The application uses PostgreSQL with GORM for database operations.

Main entities:
- User
- Product
- Category
- StockMovement

## Contributing

Contributions, issues, and feature requests are welcome.

1. Fork the repository.
2. Create a new branch.
3. Make your changes.
4. Commit your changes.
5. Push to your branch.
6. Open a Pull Request.

## Author

**Hamse Abdi**

GitHub: https://github.com/Hamse-abdi202

---

⭐ If you find this project useful, consider giving it a star!
