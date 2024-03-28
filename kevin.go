package main

import (
	"kevin/internal/component"
	"kevin/internal/persistance"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	dbAcessor := persistance.SqliteAccessorInstance("testdb")
	e := echo.New()

	// Transaction
	transactionComponent := component.TransactionComponent{
		DbAcessor: dbAcessor,
	}
	// post transaction
	e.POST("/transaction", func(c echo.Context) error {
		var request component.NewTransactionRequest
		err := c.Bind(&request)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, transactionComponent.NewTransaction(request))
	})
	// get transactions - lista tranaçoes, podemos filtrar por conta ou datas
	// adicionar tag a uma transacao

	// Bank account
	bankAccountComponent := component.BankAccountComponent{
		DbAccessor: dbAcessor,
	}
	// inserir
	e.POST("/bank-account", func(c echo.Context) error {
		var request component.NewBankAccountRequest
		err := c.Bind(&request)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, bankAccountComponent.NewBankAccount(request))
	})
	// setar valor

	// Tag
	// criar tags
	// total gasto por tag - filtro por data
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, dbAcessor.SelectAllTransactions())
	})
	e.Logger.Fatal(e.Start(":1234"))
}
