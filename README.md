# 🚀 MentalArts REST API  

MentalArts REST API is a robust and scalable RESTful API designed to manage a book library system. This project is built using **Go** with the **Gin-Gonic** framework, leveraging **PostgreSQL** for database management and **Docker** for containerization.  

The API supports user authentication, role-based access control, rate limiting, caching, structured logging, and API monitoring.  

## 📌 Project Overview  

- **Programming Language:** Go  
- **Framework:** Gin-Gonic  
- **Database:** PostgreSQL (with GORM ORM)  
- **Containerization:** Docker & Docker Compose  
- **Authentication:** JWT-based authentication  
- **Documentation:** Swagger  

---

## 📚 API Features  

This API is designed to manage books, authors, and reviews while ensuring secure and efficient operations.  

### 🏛️ Core Entities & Relationships  

- **Books**: title, author, ISBN, publication year, description  
- **Authors**: name, biography, birth date  
- **Reviews**: rating, comment, date posted  

📌 **Relationships:**  

- One **Author** can have many **Books** (1:N)  
- One **Book** can have many **Reviews** (1:N)  
- Books and Authors have a bidirectional relationship  

---

## 🔑 Admin Account  

A default admin account is created during system initialization.  

- **Email:** `admin@gmail.com`  
- **Password:** `adminpassword`
- **Role:** `admin`  

📌 **Admin Privileges:**  

- Manage all books (CRUD operations)  
- Manage all authors (CRUD operations)  
- Manage all reviews (CRUD operations)  
- View all users  
- Delete users  
- Assign roles to users  
- All **POST**, **PUT**, **DELETE** requests are restricted to Admin users  

---

## 📌 API Endpoints  

### 📖 Books  

- `GET /api/v1/books` → List all books with pagination  
- `GET /api/v1/books/:id` → Get book details with author and reviews  
- `POST /api/v1/books` → Create a new book  
- `PUT /api/v1/books/:id` → Update book details  
- `DELETE /api/v1/books/:id` → Delete a book  

### ✍️ Authors  

- `GET /api/v1/authors` → List all authors with their books  
- `GET /api/v1/authors/:id` → Get author details  
- `POST /api/v1/authors` → Create a new author  
- `PUT /api/v1/authors/:id` → Update author details  
- `DELETE /api/v1/authors/:id` → Delete an author  

### ⭐ Reviews  

- `GET /api/v1/books/:id/reviews` → Get all reviews for a book  
- `POST /api/v1/books/:id/reviews` → Add a review to a book  
- `PUT /api/v1/reviews/:id` → Update a review  
- `DELETE /api/v1/reviews/:id` → Delete a review  

### 🔐 Authentication  

- `POST /api/v1/auth/register` → User registration  
- `POST /api/v1/auth/login` → User login  
- `POST /api/v1/auth/refresh-token` → Refresh JWT token  

---

## 🛠️ Development & Deployment  

### 1️⃣ Clone the Repository  

```sh
git clone https://github.com/MentalArts/go-rest-api-abdullah-karakoc.git
cd go-rest-api-abdullah-karakoc
```

### 2️⃣ Setup Environment Variables  

Create a `.env` file in the root directory:  

```env
DB_HOST=db
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=library
DB_PORT=5432

REDIS_HOST=redis
REDIS_PORT=6379
```

### 3️⃣ Install Dependencies  

```sh
go mod tidy
```

### 4️⃣ Run the API Locally  

```sh
go run main.go
```

### 5️⃣ Run with Docker  

```sh
docker-compose up --build
```

---

## 📖 API Documentation  

To generate the API documentation:  

```sh
swag init --parseDependency --parseInternal
```

Access Swagger UI:  

```
http://localhost:8080/swagger/index.html
```

---

## 📌 Example API Requests  

### 🔐 User Registration  

```sh
curl -X POST "http://localhost:8080/api/v1/auth/register" \
-H "Content-Type: application/json" \
-d '{
  "name": "John Doe",
  "email": "johndoe@example.com",
  "password": "SecurePassword123"
}'
```

### 🔐 User Login  

```sh
curl -X POST "http://localhost:8080/api/v1/auth/login" \
-H "Content-Type: application/json" \
-d '{
  "email": "johndoe@example.com",
  "password": "SecurePassword123"
}'
```

#### 📌 Response  

```json
{
  "token": "your-jwt-token"
}
```

---

### 📖 Add a New Book (Admin Only)  

```sh
curl -X POST "http://localhost:8080/api/v1/books" \
-H "Authorization: Bearer your-jwt-token" \
-H "Content-Type: application/json" \
-d '{
  "title": "Go Programming",
  "author_id": 1,
  "isbn": "123-456-789",
  "publication_year": 2024,
  "description": "A book about Go programming"
}'
```

---

### 📖 Get All Books  

```sh
curl -X GET "http://localhost:8080/api/v1/books" \
-H "Content-Type: application/json"
```

---

### ⭐ Add a Review to a Book  

```sh
curl -X POST "http://localhost:8080/api/v1/books/1/reviews" \
-H "Authorization: Bearer your-jwt-token" \
-H "Content-Type: application/json" \
-d '{
  "rating": 5,
  "comment": "Great book!",
  "date_posted": "2025-03-09"
}'
```

---

### 🔄 Refresh Token  

```sh
curl -X POST "http://localhost:8080/api/v1/auth/refresh-token" \
-H "Authorization: Bearer your-refresh-token" \
-H "Content-Type: application/json"
```

---

## 🎯 Extra Features Implemented  

### ✅ Authentication & Authorization  

- **JWT Authentication:** Secure token-based authentication  
- **Role-Based Access Control (RBAC):** User & Admin roles with restricted actions  

### ✅ Rate Limiting  

- **Per User/IP-based rate limiting** to prevent abuse  
- Implemented using **gin-contrib/limiter**  

### ✅ Caching  

- **Redis caching** for frequently accessed endpoints  
- Improves performance and reduces database queries  

### ✅ Input Validation & Error Handling  

- **Gin validation** for request payloads  
- **Custom error responses** for consistency  

---



## 📜 License  

This project is licensed under the **MIT License**.  

---

## 📧 Contact  

For questions or contributions, please contact:  

📩 **abdullahkrkc1453@gmail.com**  
