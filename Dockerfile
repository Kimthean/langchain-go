# Use the official Golang image as the base
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go application
RUN go build -o main ./cmd/main.go

# Expose the port your application will run on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]