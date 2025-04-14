# FC Pay Gateway

Core payment processing service of FC Pay system, implemented in Go. Handles payment transactions, integrates with payment providers, and manages transaction states.

## Navigation

- [Main Repository](https://github.com/brunownk/fc-pay)
- [Web Interface](https://github.com/brunownk/fc-pay-web)
- [Antifraud Service](https://github.com/brunownk/fc-pay-antifraud)

## Topics

- go
- golang
- payment-gateway
- microservices
- postgresql
- backend
- study-project
- api

## Features

- Basic payment processing
- Simple account management
- Transaction history
- Basic Kafka integration
- PostgreSQL database

## Tech Stack

- Go 1.21+
- PostgreSQL
- Apache Kafka
- Docker

## Getting Started

1. Install dependencies:
```bash
go mod download
```

2. Set up environment variables:
```env
HTTP_PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=fc_pay_gateway
DB_SSL_MODE=disable
```

3. Run with Docker:
```bash
docker-compose up -d
```

## Study Focus

This project focuses on:
- Go fundamentals
- Basic microservices patterns
- Simple API design
- Basic database operations
- Introduction to Kafka

## License

MIT 