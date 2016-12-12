package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * User model
 */
type User struct {
  gorm.Model
  Name string `gorm:"type:varchar(40);not null"`
  Email string `gorm:"type:varchar(60);unique;not null"`
  Password string `gorm:"type:varchar(100);not null"`
  Salt string `gorm:"type:varchar(40);not null"`
}
