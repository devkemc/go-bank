package domain

type Transaction struct {
	amount          float64
	transactionType TransactionType
	accountId       int64
}

func NewTransaction(amount float64, transactionType TransactionType, accountId int64) *Transaction {
	return &Transaction{
		amount:          amount,
		transactionType: transactionType,
		accountId:       accountId,
	}
}
