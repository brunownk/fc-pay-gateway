# FC Pay Gateway

[![Go](https://img.shields.io/badge/Go-1.21-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](https://opensource.org/licenses/MIT)
[![Docker](https://img.shields.io/badge/Docker-24.0.5-blue.svg)](https://www.docker.com)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-15.4-blue.svg)](https://www.postgresql.org)
[![Kafka](https://img.shields.io/badge/Kafka-3.5.1-orange.svg)](https://kafka.apache.org)

A Go-based payment gateway service for FC Pay, handling account management, invoice processing, and payment transactions. Built with clean architecture principles and designed for scalability.

## Table of Contents

- [Getting Started](#getting-started)
- [Features](#features)
- [API Documentation](#api-documentation)
- [Architecture](#architecture)
- [Project Origin](#project-origin)

## 📝 Project Origin

This project is a fork and evolution of the original [Gateway de Pagamento](https://github.com/devfullcycle/imersao22/tree/main/go-gateway) developed during the Full Stack & Full Cycle Immersion course. 

The original project was created for educational purposes, and this version aims to:
- Deepen my understanding of Go and microservices
- Explore and implement best practices
- Experiment with different architectural patterns
- Add new features and improvements
- Create a more production-ready version

This is a personal learning journey to enhance my skills in backend development, distributed systems, and payment processing.

## 🌟 Features

<div align="center">

| Feature | Description |
|---------|-------------|
| 🔐 Account Management | API key authentication and account management |
| 💰 Payment Processing | Credit card payment processing |
| ✅ Automatic Approval | Transactions under $10,000 are automatically approved |
| 👀 Manual Review | High-value transactions (>$10,000) require manual review |
| 📊 Transaction History | Track transaction status and history |
| 📨 Kafka Integration | Asynchronous processing through Kafka |
| 🗄️ PostgreSQL | Data persistence with PostgreSQL |

</div>

## 🏗️ Architecture

### Tech Stack

<div align="center">

| Component | Technology |
|-----------|------------|
| Backend | Go 1.21+ |
| Database | PostgreSQL 16 |
| Message Queue | Apache Kafka |
| Containerization | Docker & Docker Compose |

</div>

### System Components

<div align="center">

```mermaid
graph LR
    A[Client] --> B[FC Pay Gateway]
    B --> C[PostgreSQL]
    B --> D[Kafka]
    D --> E[Antifraud Service]
    E --> D
    D --> B
```

</div>

## 🚀 Getting Started

### Prerequisites

<div align="center">

| Requirement | Version |
|-------------|---------|
| Go | 1.21+ |
| Docker | Latest |
| Docker Compose | Latest |
| migrate CLI | Latest |

</div>

### Installation Steps

1. **Clone the repository**
   ```bash
   git clone https://github.com/brunownk/fc-pay-gateway.git
   cd fc-pay-gateway
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Configure environment**
   ```bash
   cp .env.example .env
   # Edit .env with your configurations if needed
   ```

4. **Start services**
   ```bash
   docker-compose up -d
   ```

5. **Run migrations**
   ```bash
   migrate -path migrations \
           -database "postgresql://postgres:postgres@localhost:5432/fc_pay_gateway?sslmode=disable" \
           up
   ```

6. **Start the application**
   ```bash
   go run cmd/app/main.go
   ```

## 📚 API Documentation

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

## 📋 Transaction Rules

<div align="center">

| Rule | Description |
|------|-------------|
| 💸 Amount Threshold | Transactions under $10,000 are automatically approved |
| 🔍 Manual Review | Transactions of $10,000 or more require manual review |
| 🔄 Processing | All transactions are processed asynchronously through Kafka |
| ✅ Validation | Credit card information is validated before processing |

</div>

## 📖 Study Focus

<div align="center">

| Topic | Description |
|-------|-------------|
| 🔧 Go Fundamentals | Best practices and patterns |
| 🏗️ Microservices | Architecture patterns |
| 🌐 RESTful API | Design and implementation |
| 🗄️ Database | PostgreSQL operations |
| 📨 Message Queue | Kafka processing |
| 🐳 Docker | Containerization |
| 🔐 Authentication | Security and authorization |

</div>

## 🔗 Related Projects

<div align="center">

| Project | Description |
|---------|-------------|
| [Main Repository](https://github.com/brunownk/fc-pay) | Core project repository |
| [Web Interface](https://github.com/brunownk/fc-pay-web) | Web application interface |
| [Antifraud Service](https://github.com/brunownk/fc-pay-antifraud) | Fraud detection service |

</div>

## 📄 License

<div align="center">

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

</div> 