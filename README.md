# FizzBuzz API

## Project Overview
This is a production-ready FizzBuzz REST API with statistics tracking and authentication. It is implemented in Go, uses PostgreSQL for persistence, and is fully tested with unit and integration tests.

## Endpoints
- **/fizzbuzz**: Generate a FizzBuzz sequence based on input parameters.
- **/stats**: Retrieve the most frequently requested FizzBuzz sequence.
- **/stats/all**: Retrieve all FizzBuzz queries with pagination. Requires authentication.
- **/auth/register**: Register a new user.
- **/auth/login**: Log in and receive a token.
- **Health check**: `/` endpoint for liveness.

## Project Structure
```
cmd/            # Application entry point
config/         # Configuration and models
docker/         # Dockerfiles
docs/           # Swagger API documentation
internal/       # Core logic, API, middleware, server and repositories
migrations/     # SQL migrations for PostgreSQL
pkg/            # Utility packages: crypto, db and logger
scripts/        # Utility scripts for migration and DB readiness
sample.env      # Sample environment variables
```

## Requirements
- Go 1.25
- Docker & Docker Compose
- PostgreSQL 15+

## Getting Started

1. **Clone the repository**
```bash
git clone git@github.com:Anacardo89/fizzbuzz-api.git
cd fizzbuzz-api
```

2. **Set up environment variables**
```bash
cp sample.env .env
# edit .env as needed
```

3. **Run with Docker Compose**
```bash
docker-compose up -d
```

## API Documentation
Full OpenAPI specification is available in `docs/swagger.yaml`.
You can open it with [Swagger Editor](https://editor.swagger.io/) or any compatible tool.

## Testing
Run all tests (unit + integration):
```bash
go test -v ./...
```

### Notes
- Integration tests use in-memory mock databases.
- Token authentication is required for `/stats/all`.

## Endpoints Summary
| Endpoint | Method | Description | Auth Required |
|----------|--------|-------------|---------------|
| `/` | GET | Health check | No |
| `/fizzbuzz` | GET | Generate FizzBuzz sequence | No |
| `/stats` | GET | Top FizzBuzz query | No |
| `/stats/all` | GET | All FizzBuzz queries with pagination | Yes (Bearer Token) |
| `/auth/register` | POST | Register a user | No |
| `/auth/login` | POST | Login and receive JWT | No |


## Example Requests
All examples assume the API is running locally at `http://localhost:8080`.

### Health Check
#### Postman
**Method:** `GET`
**URL:** `http://localhost:8080/`
#### Curl
```bash
curl -X GET http://localhost:8080/
```
**Response:**
```json
{
  "status": "OK"
}
```

### FizzBuzz
#### Postman
**Method:** `GET`
**URL:** `http://localhost:8080/fizzbuzz?int1=3&int2=5&str1=Fizz&str2=Buzz&limit=16`

#### Curl
```bash
curl -X GET "http://localhost:8080/fizzbuzz?int1=3&int2=5&str1=Fizz&str2=Buzz&limit=16"
```

**Response:**
```json
{
  "payload": ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz","16"]
}
```

### Stats
#### Postman
**Method:** `GET`
**URL:** `http://localhost:8080/stats`

#### Curl
```bash
curl -X GET http://localhost:8080/stats
```

**Response:**
```json
{
  "int1": 3,
  "int2": 5,
  "str1": "Fizz",
  "str2": "Buzz",
  "hits": 10
}
```

### Stats All
#### Postman
**Method:** `GET`
**URL:** `http://localhost:8080/stats/all?offset=0&limit=10`
**Requires Token**

#### Curl
```bash
curl -X GET "http://localhost:8080/stats/all?offset=0&limit=10" -H "Authorization: Bearer <your_jwt_token>"
```

**Response:**
```json
{
  "stats": [
    {"int1":3,"int2":5,"str1":"Fizz","str2":"Buzz","hits":10}
  ],
  "stats_len": 1,
  "offset": 0,
  "limit": 10
}
```

### Register
#### Postman
**Method:** `POST`
**URL:** `http://localhost:8080/auth/register`
**Body**
```json
{
  "username": "user",
  "password": "test"
}
```

#### Curl
```bash
curl -X POST http://localhost:8080/auth/register \
-H "Content-Type: application/json" \
-d '{"username":"user1","password":"pass123"}'
```

**Response:**
```json
{
  "user_id": "<uuid>"
}
```

### Login
#### Postman
**Method:** `POST`
**URL:** `http://localhost:8080/auth/login`
**Body**
```json
{
  "username": "user",
  "password": "test"
}
```

#### Curl
```bash
curl -X POST http://localhost:8080/auth/login \
-H "Content-Type: application/json" \
-d '{"username":"user1","password":"pass123"}'
```

**Response:**
```json
{
  "token": "<jwt_token>"
}
```