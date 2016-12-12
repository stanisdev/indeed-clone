package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Emails model
 */
type Email struct {
  gorm.Model
  Value string `gorm:"type:varchar(100);unique;not null"`
  CompanyID uint `gorm:"not null"`
}
