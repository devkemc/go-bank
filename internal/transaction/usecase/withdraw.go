package usecase

import "go_bank/internal/transaction/repository"

type Withdraw struct {
	transactionRepository repository.TransactionRepository
	accountRepository     repository.AccountRepository
}

func NewWithdraw(transaction repository.TransactionRepository, account repository.AccountRepository) *Withdraw {
	return &Withdraw{
		transactionRepository: transaction,
		accountRepository:     account,
	}
}

func (w *Withdraw) Execute(accountId int64, amount float64) error {
	account, err := w.accountRepository.FindById(accountId)
	if err != nil {
		return err
	}
	transaction, err := account.Withdraw(amount)
	if err != nil {
		return err
	}
	err = w.accountRepository.Save(account)
	if err != nil {
		return err
	}
	err = w.transactionRepository.Save(transaction)
	if err != nil {
		return err
	}
	return nil
}
