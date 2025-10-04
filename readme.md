# Attendance System API

## Description

Attendance System API is a backend service developed using the **Go (Golang)** programming language with the **Go Fiber** framework.
This API is designed to simplify employee attendance management efficiently, featuring a lightweight architecture and high performance.

## Documentation

Full API documentation can be found here:
[Postman Docs](https://documenter.getpostman.com/view/37591253/2sB3QGuByq)

## Tech Stack

- **Programming Language:** Go (Golang)
- **Framework:** Go Fiber
- **Database:** MySQL

## Setup Tutorial

### Run With Docker

1. Clone the repository:

   ```bash
   git clone https://github.com/rudinengineer/attendance-api-gofiber.git
   cd attendance-api-gofiber
   ```

1. Copy the environment file:

   ```bash
   cp .env.example .env
   ```

1. Build and run using Docker:

   ```bash
   docker network create attendance-net
   docker compose up -d
   ```

1. The API will be accessible at:

   ```
   http://127.0.0.1:3000
   ```

---

🚀 You are ready to use Attendance System API!
