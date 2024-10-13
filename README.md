# url-shortener
A simple URL shortener written in Go.

## Features
- Shortens long URLs into a hash-based short URL.
- Redirects users from the short URL to the original URL.

## How to Run

### Prerequisites
- Go installed on your machine

### Running Locally

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/url-shortener.git
   cd url-shortener
   ```

2. Run the application:
   ```bash
   go run main.go
   ```

3. Shorten a URL: Send a POST request to /shorten with the url parameter:
   ```bash
   curl -X POST -d "url=https://www.example.com" http://localhost:8080/shorten
   ```

4. Visit the shortened URL in your browser, or use curl:
   ```bash
   curl -v http://localhost:8080/s/xxxxxx
   ```

# Future Features

*Database support to persist shortened URLs.
*Error logging and better error handling.
