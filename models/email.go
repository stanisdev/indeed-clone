package models

import (
  "github.com/jinzhu/gorm"
)

type Email struct {
  gorm.Model
  Value string `gorm:"type:varchar(100);unique;not null"`
  CompanyID uint `gorm:"not null"`
}
