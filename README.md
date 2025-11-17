# 📚 PowerBook — Reading Tracker & User Authentication Service

A modular Go project built with **DDD + Hexagonal Architecture**.

---

## 📘 Overview

**PowerBook** is a backend service designed to support user registration, authentication, and secure access to user-specific data.

### Current Features

- ✅ User registration (`/users/register`)
- ✅ User login with JWT authentication (`/users/login`)
- ✅ Secure endpoint to fetch the authenticated user (`/users/me`)
- ✅ Swagger documentation with automatic OpenAPI generation
- ✅ Proper project structure using DDD (Domain-Driven Design) and Hexagonal Architecture (Ports & Adapters)
- ✅ Fully functioning development server

This repository represents the foundation for a larger system (reading logs, competitions, leaderboards, badges, etc.), but focuses on authentication and architecture setup at this stage.

---

## 🧱 Architecture

The project is built around:

### 🟦 Hexagonal Architecture (Ports & Adapters)

Hexagonal architecture enforces strict separation:

```
┌─────────────────────────────────────────┐
│             Delivery Layer              │
│   (HTTP handlers, Swagger, Middleware)  │
└───────────────────┬─────────────────────┘
                    │
┌───────────────────▼─────────────────────┐
│            Application Layer            │
│        (Commands, Use Cases, DTOs)      │
└───────────────────┬─────────────────────┘
                    │
┌───────────────────▼─────────────────────┐
│              Domain Layer               │
│ (Entities, Aggregates, Domain Logic)    │
└───────────────────┬─────────────────────┘
                    │
┌───────────────────▼─────────────────────┐
│          Infrastructure Layer           │
│    (Repositories, JWT, Config, etc.)    │
└─────────────────────────────────────────┘
```

#### Core Principles

- Domain logic stays pure (no HTTP, DB, JSON, or framework imports)
- Application layer orchestrates use cases
- Adapters implement ports
- Everything depends inward (toward the domain)

### 🟩 Domain-Driven Design

The project already includes complete DDD structure:

- **✔ Aggregates**: User
- **✔ Value Objects**: PasswordHash (embedded in entity via bcrypt)
- **✔ Domain Services**: Not yet implemented (planned for reading logic)
- **✔ Domain Events**: Not implemented yet (future competition module)

Even the simple auth module follows DDD rules and keeps core logic isolated.

---

## 📂 Project Structure

```
powerbook/
│
├── cmd/
│   └── powerbook/
│       └── main.go               # Application entry point
│
├── internal/
│   ├── config/                   # Viper configuration loader
│   │   ├── config.go
│   │   └── load.go
│   │
│   ├── domain/                   # Pure DDD domain model
│   │   └── user/
│   │       └── user.go           # User aggregate
│   │
│   ├── application/              # Use cases (business actions)
│   │   └── user/
│   │       ├── register_user.go
│   │       └── login_user.go
│   │
│   ├── ports/                    # Interfaces for adapters
│   │   ├── user_repository.go
│   │   ├── token_service.go
│   │   └── auth_service.go
│   │
│   ├── adapters/
│   │   ├── http/                 # HTTP delivery layer
│   │   │   ├── handlers/         # REST controllers
│   │   │   ├── middleware/       # JWT auth middleware
│   │   │   ├── token/            # JWT token implementation
│   │   │   └── docs/             # Auto-generated Swagger files
│   │   │
│   │   └── postgres/             # Temporary in-memory repositories
│   │       └── user_repo_memory.go
│   │
│   └── jobs/                     # Scheduled jobs (future use)
│
├── go.mod
└── README.md
```

---

## 🚀 Features Implemented

### ✔ User Registration

- **Endpoint**: `POST /users/register`
- Registers a new user and stores hashed password

### ✔ User Login with JWT

- **Endpoint**: `POST /users/login`
- Validates credentials and returns a signed JWT token

### ✔ Authentication Middleware

- Protects routes via `Bearer <token>` header

### ✔ Get Current User

- **Endpoint**: `GET /users/me`
- Returns data of the authenticated user

### ✔ Swagger UI & OpenAPI

- Auto-generated API docs via `/swagger/index.html`

### ✔ In-memory Repository

- Used for development
- Easy to switch to Postgres later without modifying domain/app layers

---

## 🔧 Running the Project

### 1. Install dependencies

```bash
go mod tidy
```

### 2. Install swag tool (Swagger generator)

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

### 3. Generate Swagger documentation

```bash
swag init -g cmd/powerbook/main.go -o internal/adapters/http/docs
```

### 4. Run server

```bash
go run ./cmd/powerbook
```

### 5. Test endpoints

#### Swagger UI
👉 http://localhost:8080/swagger/index.html

#### Health check
👉 http://localhost:8080/health

#### Registration
`POST /users/register`

#### Login
`POST /users/login`

#### Current user
`GET /users/me` with header:
```
Authorization: Bearer <token>
```

---

## 🧭 Next Steps (Roadmap)

The current project is the authentication foundation.

### Planned modules:

- [ ] Reading log system (LogReading use case)
- [ ] Competition aggregate (points, streaks, leaderboard)
- [ ] Badge/achievement system
- [ ] Telegram bot adapter
- [ ] Postgres repositories
- [ ] Redis leaderboard system
- [ ] Domain events and event bus

---

## 🏁 Conclusion

This project already includes:

- ✅ Full DDD + Hexagonal architecture
- ✅ JWT authentication
- ✅ Registration & login
- ✅ Secure endpoints
- ✅ Clean folder structure
- ✅ Swagger documentation
- ✅ Ready base for scaling into a real product

Everything is organized to support clean, scalable development as new features arrive.