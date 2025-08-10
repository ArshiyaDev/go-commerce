# Go-Commerce 🛒
A simple e-commerce platform built with **Golang** featuring **JWT authentication** and a clean REST API.

## Features
- User registration & login with **bcrypt** password hashing
- **JWT-based authentication** (access tokens)
- Product CRUD (create, read, update, delete)
- Shopping cart & order creation
- PostgreSQL database with migrations
- Modular, maintainable Go project structure
- Secure coding practices (input validation, HTTPS-ready)

## Tech Stack
- **Backend:** Golang
- **Database:** PostgreSQL (SQLite for local/dev)
- **Auth:** JWT (`github.com/golang-jwt/jwt/v5`)
- **Password Hashing:** bcrypt (`golang.org/x/crypto/bcrypt`)
- **Router:** `net/http` (can switch to chi/gorilla/mux)


