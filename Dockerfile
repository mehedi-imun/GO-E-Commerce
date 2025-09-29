# Simple Dockerfile for Go app
FROM golang:1.24-alpine

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./

# Install required packages and fetch dependencies
RUN apk update && apk upgrade && apk add --no-cache ca-certificates git build-base \
    && go mod tidy

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o main .

# Command to run
CMD ["./main"]
