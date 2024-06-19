# Use the official Golang image as the base image
FROM golang:1.22.4-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the go application
RUN go build .

# Using a lightweight alpine image
FROM alpine:latest

#Install CA certificates
RUN apk --no-cache add ca-certificates

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the builder stage to the final stage
COPY --from=builder /app/authentication .
COPY --from=builder /app/.env .

# Expose port 8000 to the outside world
EXPOSE 7000

# Command to run the executable
CMD ["./authentication"]