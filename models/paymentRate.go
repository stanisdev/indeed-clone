package models

import (
  "github.com/jinzhu/gorm"
)

type PaymentRate struct {
  gorm.Model
  Name string
}
