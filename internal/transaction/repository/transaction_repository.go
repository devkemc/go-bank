package repository

import (
	"go_bank/internal/transaction/domain"
)

type TransactionRepository interface {
	Save(account *domain.Transaction) error
}
