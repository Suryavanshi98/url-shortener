# Use the official Golang image as the base image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o url-shortener .

# Expose the port on which the app will run
EXPOSE 8080

# Start the application
CMD ["./url-shortener"]
