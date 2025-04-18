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

| Requirement | Version | Purpose |
|-------------|---------|---------|
| Go | 1.21+ | Gateway Service |
| Docker | Latest | Containerization |
| Docker Compose | Latest | Service Orchestration |
| migrate CLI | Latest | Database Migrations |

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

   > ⚠️ **Important**: Services must be started in order and each service must be healthy before starting the next.
   > The Gateway service creates the required Docker network and Kafka broker that other services will use.

   a. **Start Gateway services first**
   ```bash
   docker-compose up -d
   ```

   This will start:
   - PostgreSQL database (port 5432)
   - Kafka broker (port 9092)
   - Kafka initialization service
   - The gateway service (port 8080)

   b. **Wait for services to be healthy** (usually takes 1-2 minutes)
   ```bash
   # Check services status
   docker-compose ps
   
   # Expected output:
   # fc-pay-gateway-db-1        ... (healthy)
   # fc-pay-gateway-kafka-1     ... (healthy)
   # fc-pay-gateway-kafka-init-1... (exited)

   # Verify Kafka topics are created
   docker-compose exec kafka kafka-topics --bootstrap-server kafka:29092 --list
   
   # Expected output:
   # pending_transactions
   # transaction_results
   ```

   c. **Verify all components are working**
   ```bash
   # Check database
   docker-compose exec db pg_isready -U postgres
   # Expected: localhost:5432 - accepting connections

   # Check Kafka broker
   docker-compose exec kafka kafka-topics --bootstrap-server kafka:29092 --describe --topic pending_transactions
   # Should show topic details without errors

   # Check Gateway API
   curl http://localhost:8080/health
   # Expected: {"status":"ok"}
   ```

4. **Run database migrations**
   ```bash
   migrate -path db/migrations \
           -database "postgresql://postgres:postgres@localhost:5432/gateway?sslmode=disable" \
           up
   ```

5. **Next Steps**
   After the Gateway is running and healthy, proceed to set up:
   
   a. **Start Antifraud service**
   ```bash
   cd ../fc-pay-antifraud
   cp .env.example .env
   docker-compose up -d
   ```

   b. **Finally, start Web interface**
   ```bash
   cd ../fc-pay-web
   cp .env.example .env
   docker-compose up -d
   ```

### Troubleshooting Common Issues

1. **Kafka Connection Issues**
   ```bash
   # Restart Kafka if topics aren't visible
   docker-compose restart kafka
   
   # Wait for health check
   docker-compose ps
   
   # Verify topics again
   docker-compose exec kafka kafka-topics --bootstrap-server kafka:29092 --list
   ```

2. **Database Connection Issues**
   ```bash
   # Check database logs
   docker-compose logs db
   
   # Verify database is accepting connections
   docker-compose exec db pg_isready -U postgres
   ```

3. **Network Issues**
   ```bash
   # List networks
   docker network ls | grep fc-pay
   
   # Inspect network
   docker network inspect fc-pay-gateway_default
   ```

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