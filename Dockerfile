# Build stage
FROM golang:1.19-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files separately to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app/main

# Final stage
FROM alpine

# Set a non-root user
RUN adduser -D -g '' appuser
USER appuser

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080

CMD ["/app/main"]
