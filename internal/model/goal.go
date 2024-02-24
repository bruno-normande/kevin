package model

import (
    "gorm.io/gorm"
)

type Goal struct {
    gorm.Model
    Name string
    Value float64
    UserId uint
}
