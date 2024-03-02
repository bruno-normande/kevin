package main

import (
	"github.com/labstack/echo/v4"
	"kevin/internal/persistance"
	"net/http"
)

func main() {
	dbAcessor := persistance.SqliteAccessorInstance("testdb")
	// dbAcessor.Insert(&model.Transaction{
	//     Date:        time.Now(),
	//     Description: "teste1",
	//     Value:       2.99,
	// })

	// Transaction
	// post transaction
	// get transactions - lista tranaçoes, podemos filtrar por conta ou datas
	// adicionar tag a uma transacao

	// Bank account
	// inserir
	// setar valor

	// Tag
	// criar tags
	// total gasto por tag - filtro por data
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, dbAcessor.SelectAllTransactions())
	})

	e.Logger.Fatal(e.Start(":1234"))
}
