package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Phone model
 */
type Phone struct {
  gorm.Model
  Number string `gorm:"type:varchar(40);unique;not null"`
  CompanyID uint `gorm:"not null"`
}
