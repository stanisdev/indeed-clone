package models

import (
  "github.com/jinzhu/gorm"
)

type CompanyPhone struct {
  gorm.Model
  CompanyID uint `gorm:"not null"`
  PhoneID uint `gorm:"not null"`
}
