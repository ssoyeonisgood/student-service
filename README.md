# Student Service

A **Go + gRPC + PostgreSQL** microservice for managing student records.  
This project allows creating, reading, updating, and deleting student data via gRPC APIs and provides a simple CLI client to interact with the service.

---

## Features

- gRPC-based API server
- PostgreSQL database for persistent storage
- Student CRUD operations:
  - Create Student
  - Get Student
  - Update Student
  - Delete Student
  - List Students
- CLI client to interact with the gRPC server
- Environment variables support for sensitive data (DB credentials)

---

## Tech Stack

- **Language:** Go
- **API:** gRPC
- **Database:** PostgreSQL
- **ORM/DB Driver:** `database/sql` with `lib/pq`
- **Environment Variables:** `.env` + `github.com/joho/godotenv`
- **Protobuf:** `protoc` + Go plugin

---

## Prerequisites

- Go 1.23+ installed
- PostgreSQL installed and running
- `protoc` (Protocol Buffers) installed
- `git` installed

---

## Setup

1. **Clone the repository**

```bash
git clone https://github.com/yourusername/student-service.git
cd student-service

2. **Set up PostgreSQL database**

CREATE DATABASE studentdb;
CREATE TABLE IF NOT EXISTS students (
  id UUID PRIMARY KEY,
  name TEXT NOT NULL,
  age INT NOT NULL,
  class_id TEXT,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);


3. ****Install dependencies**
**
go mod tidy


4. **Generate gRPC code from .proto files**

protoc --go_out=pb --go-grpc_out=pb proto/student.proto


5. **Running the Project**

- Start gRPC server

go run cmd/server/main.go


- Run CLI client

go run client.go


6. **Use the menu to create, update, delete, or list students.**

