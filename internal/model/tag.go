package model

import (
    "gorm.io/gorm"
)

type Tag struct {
    gorm.Model
    Name string
    Img string
    Transactions []Transaction
    UserId uint
}
