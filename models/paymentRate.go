package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * How often payment will be paid
 */
type PaymentRate struct {
  gorm.Model
  Value string `gorm:"type:varchar(50);unique;not null"` // "per hour", "per day", "per month", "per year"
}
