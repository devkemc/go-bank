package domain

import (
	"errors"
)

type Account struct {
	id       int64
	balance  float64
	limit    float64
	isActive bool
}

func NewAccount(id int64, balance float64, limit float64, isActive bool) *Account {
	return &Account{
		id:       id,
		balance:  balance,
		limit:    limit,
		isActive: isActive,
	}
}
func (a *Account) canWithdraw(amount float64) error {
	if amount > (a.balance + a.limit) {
		return errors.New("Insufficient balance")
	}
	if !a.isActive {
		return errors.New("Account must be active")
	}
	return nil
}

func (a *Account) canDeposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Deposit amount must be greater than zero")
	}
	if !a.isActive {
		return errors.New("Account must be active")
	}
	return nil
}

func (a *Account) Withdraw(amount float64) (*Transaction, error) {
	if err := a.canWithdraw(amount); err != nil {
		return nil, err
	}
	a.balance -= amount
	transaction := NewTransaction(amount, DEBIT, a.id)
	return transaction, nil
}

func (a *Account) Deposit(amount float64) (*Transaction, error) {
	if err := a.canDeposit(amount); err != nil {
		return nil, err
	}
	a.balance += amount
	transaction := NewTransaction(amount, CREDIT, a.id)
	return transaction, nil
}
