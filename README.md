# CEMA Backend

A  healthcare management system backend built with Go, providing APIs for managing doctors, patients, medical programs, and prescriptions.

## 🔧 Tech Stack

- **Language:** Go 1.23.1
- **Web Framework:** Gin
- **Database:** MySQL
- **Authentication:** JWT (JSON Web Tokens)
- **Dependencies Management:** Go Modules
- **Testing Framework:** Go Testing with Testify

## 🌟 Features

- **Doctor Management**
  - Registration and Authentication
  - JWT-based secure access
  - Department-based organization

- **Client/Patient Management**
  - Patient registration and profile management
  - Emergency contact information
  - Health metrics tracking (height, weight, age)
  - Program enrollment system

- **Medical Programs**
  - Program creation and management
  - Symptom tracking
  - Patient enrollment tracking

- **Prescription Management**
  - Create and update prescriptions
  - Track medicine history
  - Associate prescriptions with doctors and patients

## 📦 Project Structure

```
.
├── auth/           # Authentication and JWT handling
├── cmd/           # Application entry points
│   ├── app/      # Main application setup
│   └── main.go   # Main entry point
├── config/        # Configuration management
├── db/           # Database connection and migrations
│   └── migrations/ # SQL migration files
├── docs/         # Documentation (Postman collections)
├── logging/      # Logging utilities
├── service/      # Business logic and handlers
│   ├── clients/  # Client-related services
│   ├── doctors/  # Doctor-related services
│   └── programs/ # Program-related services
└── types/        # Shared types and interfaces
```

## 🚀 Getting Started

### Prerequisites

- Go 1.23.1 or higher
- MySQL
- Git

### Environment Setup

1. Clone the repository
2. Create a `.env` file in the root directory with the following variables:
```env
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_HOST=localhost
DB_PORT=3306
DB_NAME=your_db_name
JWT_SECRET=your_jwt_secret
PORT=8080
```

### Installation

1. Install dependencies:
```bash
go mod download
```

2. Run database migrations:
```bash
# Ensure your database is created first
mysql -u your_user -p your_database < db/migrations/000001_tables.up.sql
mysql -u your_user -p your_database < db/migrations/000002_enrollments.up.sql
mysql -u your_user -p your_database < db/migrations/000003_prescriptions.up.sql
```

3. Start the server:
```bash
go run cmd/main.go
```

## 📡 API Endpoints

### Doctors
- `POST /doctors/register` - Register a new doctor
- `POST /doctors/login` - Doctor login

### Clients
- `POST /clients/register` - Register a new client
- `POST /clients/search` - Search for a client
- `POST /clients/program-enroll` - Enroll client in a program
- `GET /clients/clients` - Get all clients
- `POST /clients/prescription` - Create prescription
- `PUT /clients/prescription` - Update prescription
- `DELETE /clients/delete` - Delete client

### Programs
- `POST /programs/register` - Create a new program
- `GET /programs/all` - Get all programs

## 🔒 Security

- Password hashing using bcrypt
- JWT-based authentication
- Protected routes with middleware
- Input validation and sanitization
- Environment variable management

## 🧪 Testing

Run the tests using:
```bash
go test ./...
```

The project includes unit tests for:
- Handler functions
- Authentication
- Database operations

## 📝 Documentation

Complete API documentation is available in the Postman collection at `docs/CEMA.postman_collection.json`

## 🏗️ Project Design

The project follows clean architecture principles:
- Service-based structure for separation of concerns
- Interface-driven design for flexibility and testing
- Middleware for cross-cutting concerns
- Centralized error handling and logging

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License.