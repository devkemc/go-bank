package repository

import (
	"github.com/devkemc/go-bank/internal/transaction/domain"
)

type TransactionRepository interface {
	Save(account *domain.Transaction) error
}
