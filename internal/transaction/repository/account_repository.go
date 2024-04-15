package repository

import (
	"go_bank/internal/transaction/domain"
)

type AccountRepository interface {
	Save(account *domain.Account) error
	FindById(id int64) (*domain.Account, error)
}
