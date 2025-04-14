# FC Pay Gateway

[![Go](https://img.shields.io/badge/Go-1.21-blue.svg)](https://golang.org) [![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT) [![Docker](https://img.shields.io/badge/Docker-24.0.5-blue.svg)](https://www.docker.com) [![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15.4-blue.svg)](https://www.postgresql.org) [![Kafka](https://img.shields.io/badge/Kafka-3.5.1-orange.svg)](https://kafka.apache.org)

A Go-based payment gateway service for FC Pay, handling account management, invoice processing, and payment transactions. Built with clean architecture principles and designed for scalability.

## Table of Contents

- [Project Origin](#project-origin)
- [Features](#features)
- [Architecture](#architecture)
- [Getting Started](#getting-started)
- [API Documentation](#api-documentation)
- [Transaction Rules](#transaction-rules)
- [Study Focus](#study-focus)
- [Related Projects](#related-projects)
- [License](#license)

## Project Origin

This project is a fork and evolution of the original [Gateway de Pagamento](https://github.com/devfullcycle/imersao22/tree/main/go-gateway) developed during the Full Stack & Full Cycle Immersion course. 

The original project was created for educational purposes, and this version aims to:
- Deepen my understanding of Go and microservices
- Explore and implement best practices
- Experiment with different architectural patterns
- Add new features and improvements
- Create a more production-ready version

This is a personal learning journey to enhance my skills in backend development, distributed systems, and payment processing.

## Features

| Feature | Description |
|---------|-------------|
| 🔐 Account Management | API key authentication and account management |
| 💰 Payment Processing | Credit card payment processing |
| ✅ Automatic Approval | Transactions under $10,000 are automatically approved |
| 👀 Manual Review | High-value transactions (>$10,000) require manual review |
| 📊 Transaction History | Track transaction status and history |
| 📨 Kafka Integration | Asynchronous processing through Kafka |
| 🗄️ PostgreSQL | Data persistence with PostgreSQL |

## Architecture

### Tech Stack

| Component | Technology |
|-----------|------------|
| Backend | Go 1.21+ |
| Database | PostgreSQL 16 |
| Message Queue | Apache Kafka |
| Containerization | Docker & Docker Compose |

### System Components

```mermaid
graph LR
    A[Client] --> B[FC Pay Gateway]
    B --> C[PostgreSQL]
    B --> D[Kafka]
    D --> E[Antifraud Service]
    E --> D
    D --> B
```

## Getting Started

### Prerequisites

| Requirement | Version |
|-------------|---------|
| Go | 1.21+ |
| Docker | Latest |
| Docker Compose | Latest |
| migrate CLI | Latest |

### Installation Steps

1. **Clone the repository**
   ```bash
   git clone https://github.com/brunownk/fc-pay-gateway.git
   cd fc-pay-gateway
   ```

2. **Configure environment**
   ```bash
   cp .env.example .env
   # The default environment variables are already configured for Docker
   ```

3. **Start the services in order**

   > ⚠️ **Important**: The services must be started in the following order to ensure proper network and dependency initialization.

   a. **Start Gateway services first** (this will create the required network and Kafka broker)
   ```bash
   docker-compose up -d
   ```
   This will start:
   - PostgreSQL database (port 5432)
   - Kafka broker (port 9092)
   - Kafka initialization service
   - The gateway service (port 8080)

   b. **Start Antifraud service** (in fc-pay-antifraud directory)
   ```bash
   cd ../fc-pay-antifraud
   cp .env.example .env
   docker-compose up -d
   ```

   c. **Start Web interface** (in fc-pay-web directory)
   ```bash
   cd ../fc-pay-web
   cp .env.example .env
   docker-compose up -d
   ```

4. **Run migrations**
   ```bash
   migrate -path db/migrations \
           -database "postgresql://postgres:postgres@localhost:5432/gateway?sslmode=disable" \
           up
   ```

5. **Verify all services are running**
   ```bash
   docker-compose ps
   ```

6. **Access the application**
   - Gateway API: http://localhost:8080
   - Antifraud Service: http://localhost:3001
   - Web Interface: http://localhost:3000

7. **Run the application locally (optional)**
   ```bash
   go run cmd/app/main.go
   ```
   This will start the gateway service on port 8080. Make sure all required services (PostgreSQL and Kafka) are running before starting the application.

### Docker Network Configuration

The gateway service creates a Docker network named `fc-pay-network` that other services will use to communicate. The network configuration includes:

- Database: `db:5432`
- Kafka: `kafka:29092`
- Gateway API: `app:8080`

### Service Dependencies

The gateway service depends on:
- PostgreSQL 16 (for data persistence)
- Apache Kafka (for asynchronous processing)
- The antifraud service (for transaction validation)

### Health Checks

You can verify the services are healthy by:

1. **Database**
   ```bash
   docker-compose exec db pg_isready -U postgres
   ```

2. **Kafka**
   ```bash
   docker-compose exec kafka kafka-topics --bootstrap-server kafka:29092 --list
   ```

3. **Gateway API**
   ```bash
   curl http://localhost:8080/health
   ```

## API Documentation

### Create Account

```http
POST /accounts
Content-Type: application/json

{
    "name": "John Doe",
    "email": "john@example.com"
}
```

**Response:**
```json
{
    "id": "uuid",
    "name": "John Doe",
    "email": "john@example.com",
    "api_key": "generated-api-key",
    "balance": 0,
    "created_at": "timestamp",
    "updated_at": "timestamp"
}
```

### Create Invoice

```http
POST /invoice
Content-Type: application/json
X-API-Key: {your_api_key}

{
    "amount": 100.50,
    "description": "Product purchase",
    "payment_type": "credit_card",
    "card_number": "4111111111111111",
    "cvv": "123",
    "expiry_month": 12,
    "expiry_year": 2025,
    "cardholder_name": "John Doe"
}
```

**Response:**
```json
{
    "id": "uuid",
    "account_id": "uuid",
    "amount": 100.50,
    "status": "approved",
    "description": "Product purchase",
    "payment_type": "credit_card",
    "card_last_digits": "1111",
    "created_at": "timestamp",
    "updated_at": "timestamp"
}
```

### Get Invoice Details

```http
GET /invoice/{invoice_id}
X-API-Key: {your_api_key}
```

## Transaction Rules

| Rule | Description |
|------|-------------|
| 💸 Amount Threshold | Transactions under $10,000 are automatically approved |
| 🔍 Manual Review | Transactions of $10,000 or more require manual review |
| 🔄 Processing | All transactions are processed asynchronously through Kafka |
| ✅ Validation | Credit card information is validated before processing |

## Study Focus

| Topic | Description |
|-------|-------------|
| 🔧 Go Fundamentals | Best practices and patterns |
| 🏗️ Microservices | Architecture patterns |
| 🌐 RESTful API | Design and implementation |
| 🗄️ Database | PostgreSQL operations |
| 📨 Message Queue | Kafka processing |
| 🐳 Docker | Containerization |
| 🔐 Authentication | Security and authorization |

## Related Projects

| Project | Description |
|---------|-------------|
| [Main Repository](https://github.com/brunownk/fc-pay) | Core project repository |
| [Web Interface](https://github.com/brunownk/fc-pay-web) | Web application interface |
| [Antifraud Service](https://github.com/brunownk/fc-pay-antifraud) | Fraud detection service |

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details. 