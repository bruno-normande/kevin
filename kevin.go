package main

import (
	"encoding/csv"
	//	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"kevin/internal/model"
	"kevin/internal/persistance"
)

func main() {
    // content, _ := loadCsv("testData/cibc_checking_2024_01.csv")
    dbAcessor := persistance.NewSquliteAccessor("testdb")
    dbAcessor.Insert(&model.Transaction{
        Date: time.Now(),
        Description: "teste1",
        Value: 2.99,
    })

    e := echo.New()
    e.GET("/", func(c echo.Context) error {
        return c.JSON(http.StatusOK, dbAcessor.SelectAllTransactions())
    })
    e.Logger.Fatal(e.Start(":1234"))
}

func loadCsv(fileName string) ([]model.Transaction, error) {
    file, err := os.Open(fileName)
    if err != nil {
        return nil, err 
    }
    defer file.Close()

    reader := csv.NewReader(file)

    records, err := reader.ReadAll()
    if err != nil {
        return nil, err 
    }
    var transactions []model.Transaction

    for _, record := range records {
        transaction := transactionFromRecord(record)
        // fmt.Println(transaction)
        transactions = append(transactions, transaction)
    }
    return transactions, nil
}

func transactionFromRecord(record []string) model.Transaction {
    dateString := record[0]
    date, _ := time.Parse("2006-01-02", dateString)
    description := record[1]
    value := record[2]
    var valueFloat float64
    if value != "" {
        valueFloat, _ = strconv.ParseFloat(value, 64)
        valueFloat = -valueFloat
    } else {
        valueFloat, _ = strconv.ParseFloat(record[3], 64)
    }
    //fmt.Printf("%v - %v - %v", date, description, value)

    return model.Transaction{
        Date: date, 
        Description: description,
        Value: valueFloat,
    }
}

