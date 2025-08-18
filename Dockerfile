FROM golang:1.25.0-trixie AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o api ./cmd


FROM debian:trixie-slim
WORKDIR /app
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
COPY --from=builder /app/api .
COPY migrations ./migrations

ENV PORT=8080
ENV DATABASE_URL=postgres://user:pass@host:5432/dbname?sslmode=disable

EXPOSE ${PORT}

CMD ["./api"]
