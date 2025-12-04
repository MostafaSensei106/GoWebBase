# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application as a static binary
# Using CGO_ENABLED=0 to create a statically linked executable
# Using -ldflags="-w -s" to strip debug information and reduce binary size
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-w -s" -o /gowebbase .

# Stage 2: Create the final, minimal image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the static binary from the builder stage
COPY --from=builder /gowebbase .

# Copy the static assets required by the web application
COPY --from=builder /app/static ./static

# Expose the port the application runs on
EXPOSE 8080

# The command to run the application
CMD ["./gowebbase"]
