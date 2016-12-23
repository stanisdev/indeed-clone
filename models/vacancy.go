package models

import (
  "github.com/jinzhu/gorm"
  "time"
)

/**
 * Vacancy model
 */
type Vacancy struct {
  gorm.Model
  Title string `gorm:"type:varchar(255);not null"`
  ExpiredDate time.Time
  Description string `gorm:"type:text;not null"`
  SalarySize int `gorm:"not null"`
  PaymentRateID uint `gorm:"not null"`
  UserID uint `gorm:"not null"`
  CompanyID uint `gorm:"not null"`
}
