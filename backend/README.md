# Go Backend Server

This project is a basic backend server built with Go that connects to a MongoDB database. It provides a simple API for managing items.

## Project Structure

```
go-backend-server
├── cmd
│   └── main.go          # Entry point of the application
├── internal
│   ├── config
│   │   └── config.go    # Configuration management
│   ├── db
│   │   └── mongo.go      # MongoDB connection management
│   ├── handlers
│   │   └── handlers.go    # HTTP handler functions
│   ├── models
│   │   └── models.go      # Data models
│   └── routes
│       └── routes.go      # Application routes
├── go.mod                # Module dependencies
├── go.sum                # Module checksums
└── README.md             # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd go-backend-server
   ```

2. **Install dependencies:**
   ```
   go mod tidy
   ```

3. **Set up environment variables:**
   Create a `.env` file in the root directory with the following content:
   ```
   PORT=8080
   DB_HOST=localhost
   DB_USER=user
   DB_PASSWORD=password
   DB_NAME=dbname
   ```

4. **Run the server:**
   ```
   go run cmd/main.go
   ```

## Usage

- The server listens on the specified port (default: 8080).
- You can access the API endpoints to manage items.

## License

This project is licensed under the MIT License.