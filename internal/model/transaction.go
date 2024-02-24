package model

import (
	"time"
    "gorm.io/gorm"
)

type Transaction struct {
    gorm.Model
	Date time.Time
	Description string
	Value float64
    BankAccountId uint
    TagId uint
}
