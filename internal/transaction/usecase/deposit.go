package usecase

import "go_bank/internal/transaction/repository"

type Deposit struct {
	accountRepository     repository.AccountRepository
	transactionRepository repository.TransactionRepository
}

func NewDeposit(account repository.AccountRepository, transaction repository.TransactionRepository) *Deposit {
	return &Deposit{
		accountRepository:     account,
		transactionRepository: transaction,
	}
}

func (d *Deposit) Execute(accountId int64, amount float64) error {
	account, err := d.accountRepository.FindById(accountId)
	if err != nil {
		return err
	}
	transaction, err := account.Deposit(amount)
	if err != nil {
		return err
	}
	err = d.accountRepository.Save(account)
	if err != nil {
		return err
	}
	err = d.transactionRepository.Save(transaction)
	if err != nil {
		return err
	}
	return nil
}
