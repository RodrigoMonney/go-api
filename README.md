# Go Clean Architecture API

This is a simple REST API built with Go using the [Gin](https://github.com/gin-gonic/gin) framework and following the principles of Clean Architecture.

The API allows you to create and retrieve users stored in memory.

## 🔧 Technologies

- [Go](https://golang.org/) 1.21+
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- Clean Architecture (Domain-Driven, Separation of Concerns)

## 📁 Folder Structure

```bash
nome-api/
├── cmd/                # Application entry point
│   └── server/main.go
├── internal/
│   ├── domain/         # Entities and interfaces (core)
│   ├── usecase/        # Business use cases
│   ├── interface/
│   │   ├── handler/    # HTTP handlers (controllers)
│   │   └── repository/ # Infrastructure (data access)
```

## 🚀 How to Run

1. Clone the repository:

```bash
git clone https://github.com/yourusername/nome-api.git
cd nome-api
```

2. Install dependencies:

```bash
go mod tidy
```

3. Run the server:

```bash
go run cmd/server/main.go
```

The API will be available at `http://localhost:8080`

## 📌 Available Routes

| Method | Route        | Description       |
| ------ | ------------ | ----------------- |
| GET    | `/users`     | List all users    |
| GET    | `/users/:id` | Get a user by ID  |
| POST   | `/users`     | Create a new user |

### Sample payload for `POST /users`:

```json
{
  "name": "John Doe",
  "email": "john@example.com"
}
```

## 🧪 Testing (coming soon)

You can easily add unit tests for use cases and handlers.

## 📄 License

This project is licensed under the MIT License.
