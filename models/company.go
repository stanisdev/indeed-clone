package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Company model
 */
type Company struct {
  gorm.Model
  Name string `gorm:"not null;unique"`
  CityID uint `gorm:"not null"`
  Url string `gorm:"type:varchar(100)"`
  EmployeesCountID uint
}
