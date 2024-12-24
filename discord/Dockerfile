# Stage 1: Build the application
FROM golang:1.23-alpine AS builder

# Install necessary tools
RUN apk add --no-cache git

# Set the working directory
WORKDIR /app

# Copy backend and common directories to the build context
COPY discord/ ./discord
COPY common/ ./common

# Change to the backend directory
WORKDIR /app/discord

# Cache dependencies
RUN go mod download

# Build the backend binary
RUN CGO_ENABLED=0 GOOS=linux go build -o discord .

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/discord/discord .

# Expose the application port
EXPOSE 8555

# Command to run the application
CMD ["./discord"]
