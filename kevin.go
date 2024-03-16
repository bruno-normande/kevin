package main

import (
	"kevin/internal/component"
	"kevin/internal/persistance"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	dbAcessor := persistance.SqliteAccessorInstance("testdb")
	// Transaction
	// post transaction
	// get transactions - lista tranaçoes, podemos filtrar por conta ou datas
	// adicionar tag a uma transacao
	transactionComponent := component.TransactionComponent{
		DbAcessor: dbAcessor,
	}
	// Bank account
	// inserir
	// setar valor

	// Tag
	// criar tags
	// total gasto por tag - filtro por data
	e := echo.New()
	e.POST("/transaction", func(c echo.Context) error {
		var request component.PostTransactionRequest
		err := c.Bind(&request)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, transactionComponent.Post(request))
	})
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, dbAcessor.SelectAllTransactions())
	})
	e.Logger.Fatal(e.Start(":1234"))
}
