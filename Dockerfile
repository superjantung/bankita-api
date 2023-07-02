# Build stage
FROM golang:1.19-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files separately to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -o /app/main
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.16.2/migrate.linux-amd64.tar.gz | tar xvz

# Final stage
FROM alpine

# Set a non-root user
RUN adduser -D -g '' appuser
USER appuser

WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8080

CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]