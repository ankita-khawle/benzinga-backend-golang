# Use official Golang image as a builder
FROM golang:1.19 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN go build -o server

# Use a minimal image to run the application
FROM alpine:latest

WORKDIR /main/
# Copy the built server from the builder
COPY --from=builder /main/server .

# Expose port 8080
EXPOSE 8080

# Run the server
CMD ["./server"]
