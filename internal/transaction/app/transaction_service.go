package app

import (
	"github.com/devkemc/go-bank/internal/transaction/usecase"
)

type TransactionService struct {
	transfer *usecase.Transfer
	withdraw *usecase.Withdraw
	deposit  *usecase.Deposit
}

func NewTransactionService(transfer *usecase.Transfer, withdraw *usecase.Withdraw, deposit *usecase.Deposit) *TransactionService {
	return &TransactionService{
		transfer: transfer,
		withdraw: withdraw,
		deposit:  deposit,
	}
}

func (s *TransactionService) Transfer(toId int64, fromId int64, amount float64) error {
	return s.transfer.Execute(toId, fromId, amount)
}

func (s *TransactionService) Withdraw(accountId int64, amount float64) error {
	return s.withdraw.Execute(accountId, amount)
}

func (s *TransactionService) Deposit(accountId int64, amount float64) error {
	return s.deposit.Execute(accountId, amount)
}
