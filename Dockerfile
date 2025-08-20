# Stage 1: Build Go binary
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum, lalu download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy semua source code
COPY . .

# Build binary (ubah `./cmd` sesuai dengan folder main.go kamu)
RUN go build -o inventory-system ./cmd

# Stage 2: Create runtime image
FROM alpine:3.18

# Set working directory
WORKDIR /app

# Copy binary hasil build dari stage 1
COPY --from=builder /app/inventory-system .

# Jalankan aplikasi
CMD ["./inventory-system"]
