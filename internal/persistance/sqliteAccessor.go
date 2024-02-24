package persistance

import (
	"fmt"
	"kevin/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteAccessor struct {
    db *gorm.DB
}

func NewSquliteAccessor(dbName string) *sqliteAccessor {
    db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
    if err != nil {
        panic("couldn't initializate the DB")
    }

    db.AutoMigrate(&model.User{})
    db.AutoMigrate(&model.BankAccount{})
    db.AutoMigrate(&model.Goal{})
    db.AutoMigrate(&model.Tag{})
    db.AutoMigrate(&model.Transaction{})

    return &sqliteAccessor{
        db: db,
    }
}

func (accessor sqliteAccessor) Insert(transaction *model.Transaction) {
    accessor.db.Create(transaction)
}

func (accessor sqliteAccessor) SelectAllTransactions() []model.Transaction {
    var transactions []model.Transaction
    result := accessor.db.Find(&transactions)
    fmt.Println("found: ", result.RowsAffected)
    return transactions
 }
