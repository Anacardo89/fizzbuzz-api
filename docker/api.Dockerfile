FROM golang:1.25.0-trixie AS builder
ARG APP_PATH=/fizzbuzz-api
WORKDIR $APP_PATH
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o fizzbuzz-api ./cmd


FROM debian:trixie-slim
ARG APP_PATH=/fizzbuzz-api
WORKDIR $APP_PATH
RUN apt-get update && apt-get install -y ca-certificates netcat-openbsd
RUN rm -rf /var/lib/apt/lists/*
COPY --from=builder $APP_PATH .
COPY scripts/wait-for-db.sh ./scripts/

ENV PORT=8080

EXPOSE ${PORT}

ENTRYPOINT ["/fizzbuzz-api/scripts/wait-for-db.sh"]
