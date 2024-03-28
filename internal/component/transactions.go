package component

import (
	"kevin/internal/model"
	"kevin/internal/persistance"
	"time"
)

type TransactionComponent struct {
	DbAcessor persistance.DbAccessor
}

type NewTransactionRequest struct {
	Description    string    `json:"description" form:"description" query:"description"`
	Value          float64   `json:"value" form:"value" query:"value"`
	BanckAccountId uint      `json:"bank_account_id" form:"bank_account_id" query:"bank_account_id"`
	Date           time.Time `json:"date" form:"date" query:"date"`
}

type NewTransactionResponse struct {
	Id            uint
	Date          time.Time
	Description   string
	Value         float64
	BankAccountId uint
	Tag           Tag
}

type Tag struct {
	Id   uint
	Name string
	Img  string
}

func (component *TransactionComponent) NewTransaction(request NewTransactionRequest) NewTransactionResponse {
	transaction := model.Transaction{
		Date:          request.Date,
		Description:   request.Description,
		Value:         request.Value,
		BankAccountId: request.BanckAccountId,
	}

	component.DbAcessor.Insert(&transaction)
	return NewTransactionResponse{
		Id:            transaction.ID,
		Date:          transaction.Date,
		Description:   transaction.Description,
		Value:         transaction.Value,
		BankAccountId: transaction.BankAccountId,
	}
}
