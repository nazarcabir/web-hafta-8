# Aşama 1: Derleme
FROM golang:1.22-alpine AS builder

# build-base for CGO (SQLite), git for go mod
RUN apk add --no-cache build-base git

WORKDIR /app

# Copy all files
COPY . .

# Initialize module if it doesn't exist
RUN go mod init golearn || true

# Pre-create docs package to satisfy imports before swag init
RUN mkdir -p docs && echo "package docs" > docs/docs.go

# Install swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Tidy first to get dependencies (docs package now exists)
RUN go mod tidy

# Generate swagger docs
RUN /go/bin/swag init

# Tidy again just in case swag added new dependencies
RUN go mod tidy

# Build
RUN CGO_ENABLED=1 go build -o golearn .

# Aşama 2: Çalıştırma (minimal image)
FROM alpine:latest
RUN apk add --no-cache libc6-compat
WORKDIR /app
COPY --from=builder /app/golearn .
EXPOSE 8090
CMD ["./golearn"]
