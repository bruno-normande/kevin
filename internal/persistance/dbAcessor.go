package persistance

import (
	"kevin/internal/model"
)

type DbAccessor interface {
	Insert(transaction *model.Transaction)
	SaveBankAccount(bankAccount *model.BankAccount)
	SelectAllTransactions() []model.Transaction
}
