package domain

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// Status representa os estados possíveis de uma fatura
type Status string

const (
	StatusPending  Status = "pending"
	StatusApproved Status = "approved"
	StatusRejected Status = "rejected"
)

// Invoice representa uma fatura no sistema de pagamento
type Invoice struct {
	ID             string
	AccountID      string
	Amount         float64
	Status         Status
	Description    string
	PaymentType    string
	CardLastDigits string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// CreditCard contém os dados sensíveis do cartão de crédito
type CreditCard struct {
	Number         string
	CVV            string
	ExpiryMonth    int
	ExpiryYear     int
	CardholderName string
}

// NewInvoice cria uma nova fatura e valida o valor
func NewInvoice(accountID string, amount float64, description string, paymentType string, card CreditCard) (*Invoice, error) {
	if amount <= 0 {
		return nil, ErrInvalidAmount
	}

	// Armazena apenas os últimos 4 dígitos do cartão por segurança
	lastDigits := card.Number[len(card.Number)-4:]

	return &Invoice{
		ID:             uuid.New().String(),
		AccountID:      accountID,
		Amount:         amount,
		Status:         StatusPending,
		Description:    description,
		PaymentType:    paymentType,
		CardLastDigits: lastDigits,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}, nil
}

// Process simula o processamento da fatura com regras de negócio
func (i *Invoice) Process() error {
	if i.Amount > 10000 {
		return nil
	}

	// Simula aprovação com 70% de chance
	randomSource := rand.New(rand.NewSource(time.Now().Unix()))
	var newStatus Status

	if randomSource.Float64() <= 0.7 {
		newStatus = StatusApproved
	} else {
		newStatus = StatusRejected
	}

	i.Status = newStatus
	i.UpdatedAt = time.Now()
	return nil
}

// UpdateStatus permite atualizar o status apenas de faturas pendentes
func (i *Invoice) UpdateStatus(newStatus Status) error {
	if i.Status != StatusPending {
		return ErrInvalidStatus
	}

	i.Status = newStatus
	i.UpdatedAt = time.Now()
	return nil
}
