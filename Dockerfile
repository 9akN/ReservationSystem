# ---- Build stage ----
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache git

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o bin/reservation-system ./cmd/...

# ---- Run stage ----
FROM alpine:3.19

WORKDIR /app

# Add CA certificates for HTTPS and timezone data
RUN apk add --no-cache ca-certificates tzdata

COPY --from=builder /app/bin/reservation-system .
COPY --from=builder /app/migrations ./migrations

EXPOSE 8080

ENTRYPOINT ["./reservation-system"]
