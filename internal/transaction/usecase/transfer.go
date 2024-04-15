package usecase

import "go_bank/internal/transaction/repository"

type Transfer struct {
	accountRepository     repository.AccountRepository
	transactionRepository repository.TransactionRepository
}

func NewTransfer(accountRepository repository.AccountRepository, transaction repository.TransactionRepository) *Transfer {
	return &Transfer{
		accountRepository:     accountRepository,
		transactionRepository: transaction,
	}
}
func (t *Transfer) Execute(toId int64, fromId int64, amount float64) error {
	toAccount, err := t.accountRepository.FindById(toId)
	if err != nil {
		return err
	}
	fromAccount, err := t.accountRepository.FindById(fromId)
	if err != nil {
		return err
	}
	transactionFrom, errFrom := fromAccount.Withdraw(amount)
	if errFrom != nil {
		return errFrom
	}
	transactionTo, errTo := toAccount.Deposit(amount)
	if errTo != nil {
		return errTo
	}
	err = t.accountRepository.Save(fromAccount)
	if err != nil {
		return err
	}
	err = t.accountRepository.Save(toAccount)
	if err != nil {
		return err
	}
	err = t.transactionRepository.Save(transactionFrom)
	if err != nil {
		return err
	}
	err = t.transactionRepository.Save(transactionTo)
	if err != nil {
		return err
	}
	return nil
}
