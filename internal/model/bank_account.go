package model

import (
    "gorm.io/gorm"
)

type BankAccount struct {
    gorm.Model
    Name string
    UserId uint
    Transactions []Transaction
}


