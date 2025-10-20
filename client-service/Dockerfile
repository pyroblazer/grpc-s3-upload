# Build stage
FROM golang:1.22-alpine AS build

WORKDIR /app

# Copy source code
COPY . .

# Initialize module, tidy dependencies, and build binary
RUN go mod init client-service && \
    go mod tidy && \
    go build -o client

# Runtime stage
FROM alpine:3.19

WORKDIR /app

# Copy binary and config
COPY --from=build /app/client .
COPY config ./config

# Run the app
CMD ["./client"]
