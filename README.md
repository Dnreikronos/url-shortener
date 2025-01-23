# URL Shortener

A simple and efficient URL shortening service built using Golang. This application allows users to shorten long URLs and expand shortened URLs back to their original form. It uses an SQLite database for persistence and supports both unit and integration testing.

## Features
- Shorten long URLs into compact, easy-to-share links.
- Expand shortened URLs back to their original form.
- Persistent storage using SQLite.
- Unit and integration tests for core functionality.
- Designed for easy extensibility.

---

## Installation

### Prerequisites
- Go (1.20 or higher)
- SQLite (for local database storage)

### Clone the Repository
```bash
git clone https://github.com/Dnreikronos/url-shortener.git
cd url-shortener
```

### Install Dependencies
Ensure you have all the required Go modules:
```bash
go mod tidy
```

### Run the Application
```bash
go run ./main.go
```

---

## Usage

### Shorten a URL
Use the `--Shorten` flag to shorten a URL:
```bash
go run ./main.go --Shorten=https://example.com
```
Output:
```
Shortened URL: <short-key>
```

### Expand a URL
Use the `--Expand` flag to retrieve the original URL:
```bash
go run ./main.go --Expand=<short-key>
```
Output:
```
Original URL: https://example.com
```

### Error Handling
If the short key is not found, the application will display:
```
Error expanding URL or URL not found.
```

---

## Project Structure
```
url-shortener/
├── main.go             # Entry point for the application
├── shortener/          # Logic for URL shortening and expanding
├── storage/            # Database logic for storing and retrieving URLs
├── tests/              # Unit and integration tests
├── go.mod              # Go module file
└── README.md           # Project documentation
```

---

## Testing

### Run Unit Tests
```bash
cd tests
```
```bash
go test -v
```





