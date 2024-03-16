package persistance

import (
	"kevin/internal/model"
)

type DbAccessor interface {
	Insert(transaction *model.Transaction)
	SelectAllTransactions() []model.Transaction
}
