package persistance

import (
	"fmt"
	"kevin/internal/model"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type sqliteAccessor struct {
	db *gorm.DB
}

// instancia singleton
var instance DbAccessor

// é usado para garantir que so vai ser executado uma vez
var once sync.Once

func SqliteAccessorInstance(dbName string) DbAccessor {
	once.Do(func() {
		db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			panic("couldn't initializate the DB")
		}

		db.AutoMigrate(&model.User{})
		db.AutoMigrate(&model.BankAccount{})
		db.AutoMigrate(&model.Goal{})
		db.AutoMigrate(&model.Tag{})
		db.AutoMigrate(&model.Transaction{})

		instance = &sqliteAccessor{
			db: db,
		}
	})

	return instance
}

func (accessor sqliteAccessor) Insert(transaction *model.Transaction) {
	accessor.db.Save(transaction) // TODO: how to deal with errors using gorm?
}

func (accessor sqliteAccessor) SaveBankAccount(bankAccount *model.BankAccount) {
	accessor.db.Save(bankAccount)
}

func (accessor sqliteAccessor) SelectAllTransactions() []model.Transaction {
	var transactions []model.Transaction
	result := accessor.db.Find(&transactions)
	fmt.Println("found: ", result.RowsAffected)
	return transactions
}
