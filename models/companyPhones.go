package models

import (
  "github.com/jinzhu/gorm"
)

type CompanyPhone struct {
  gorm.Model
  CompanyID uint
  PhoneID uint
}
