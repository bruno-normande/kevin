package model

import (
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Email string
    BankAccounts []BankAccount
    Goals []Goal
    Tags []Tag
}
