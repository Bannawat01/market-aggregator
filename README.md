# Real-Time Market Data Aggregator

A high-performance backend service built with Go that demonstrates clean architecture, concurrent background processing, and real-time data aggregation. The system continuously fetches cryptocurrency prices from external APIs and serves historical data via a RESTful API.

## Key Features

- Concurrent background workers using goroutines and sync.WaitGroup to fetch multiple endpoints in parallel.
- Automated data pipeline with time.Ticker to aggregate market data every 5 seconds without blocking the API server.
- Layered architecture (Handler -> Service -> Repository) for maintainability and testability.
- RESTful API built with Gin to serve aggregated time-series data.
- Database integration with GORM and SQLite for persistence.
- Environment-based configuration using godotenv.

## Tech Stack

- Go (Golang)
- Gin
- GORM
- SQLite
- godotenv

## Architecture and Project Structure

```
market-aggregator/
├── cmd/
│   └── api/
│       └── main.go         # Application entry point and dependency injection
├── internal/
│   ├── handler/            # HTTP layer (Gin), handles requests and JSON responses
│   ├── service/            # Business logic layer
│   ├── repository/         # Database interaction layer (GORM)
│   ├── model/              # Data structures and database schemas
│   └── worker/             # Background daemon for concurrent data fetching
├── .env                    # Environment variables (target coins, ports)
└── go.mod                  # Go module dependencies
```

## Getting Started

### Prerequisites

- Go 1.20 or higher

### Installation

1. Clone the repository:

```bash
git clone https://github.com/Bannawat101/market-aggregator.git
cd market-aggregator
```

2. Install dependencies:

```bash
go mod tidy
```

3. Create a .env file in the root directory:

```bash
TARGET_COINS=BTCUSDT,ETHUSDT,DOGEUSDT,XRPUSDT
PORT=8080
```

4. Run the application:

```bash
go run cmd/api/main.go
```

## API Endpoints

### Get Latest Market Prices

Retrieves the most recently aggregated market prices.

- URL: `/api/v1/prices`
- Method: `GET`
- Query Parameters:
  - `limit` (optional): Number of records to return. Default is 10. Maximum is 100.

Example request:

```http
GET http://localhost:8080/api/v1/prices?limit=3
```

Example response:

```json
{
  "message": "Success",
  "data": [
    {
      "ID": 15,
      "CreatedAt": "2026-05-19T20:49:08.123Z",
      "UpdatedAt": "2026-05-19T20:49:08.123Z",
      "DeletedAt": null,
      "symbol": "BTCUSDT",
      "price": 77053.7
    },
    {
      "ID": 14,
      "CreatedAt": "2026-05-19T20:49:08.123Z",
      "UpdatedAt": "2026-05-19T20:49:08.123Z",
      "DeletedAt": null,
      "symbol": "ETHUSDT",
      "price": 2125.79
    }
  ]
}
```

## Concepts Demonstrated

- Concurrency with goroutines and sync.WaitGroup for non-blocking I/O.
- Graceful resource management using defer to close HTTP bodies and finalize waits.
- Separation of concerns across handlers, services, and repositories.
