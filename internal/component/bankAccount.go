package component

import (
	"kevin/internal/model"
	"kevin/internal/persistance"
)

type BankAccountComponent struct {
	DbAccessor persistance.DbAccessor
}

type NewBankAccountRequest struct {
	Name string `json:"name" form:"name" query:"name"`
}

type NewBankAccountResponse struct {
	Id      uint    `json:"id"`
	Name    string  `json:"name"`
	Balance float64 `json:"balance"`
}

func (componet *BankAccountComponent) NewBankAccount(request NewBankAccountRequest) NewBankAccountResponse {
	bankAccount := model.BankAccount{
		Name: request.Name,
	}

	componet.DbAccessor.SaveBankAccount(&bankAccount)
	return NewBankAccountResponse{
		Id:   bankAccount.ID,
		Name: bankAccount.Name,
	}
}
