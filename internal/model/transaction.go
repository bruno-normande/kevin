package model

import (
	"gorm.io/gorm"
	"time"
)

type Transaction struct {
	gorm.Model
	Date          time.Time
	Description   string
	Value         float64
	BankAccountId uint
	TagId         uint
}
