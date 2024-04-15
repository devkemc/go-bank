package domain_test

import (
	"github.com/devkemc/go-bank/internal/transaction/domain"
	"testing"
)

func TestAccountWithdraw(t *testing.T) {
	account := domain.NewAccount(2, 50.0, 30, true)

	// Test withdrawing an amount within balance and limit
	_, err := account.Withdraw(75.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test withdrawing an amount exceeding balance and limit
	_, err = account.Withdraw(200.0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestAccountDeposit(t *testing.T) {
	account := domain.NewAccount(2, 50.0, 30, true)

	// Test depositing a positive amount
	_, err := account.Deposit(50.0)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test depositing a non-positive amount
	_, err = account.Deposit(0.0)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
