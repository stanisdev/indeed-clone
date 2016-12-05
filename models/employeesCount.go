package models

import (
  "github.com/jinzhu/gorm"
)

type EmployeesCount struct {
  gorm.Model
  Value string `gorm:"type:varchar(50);unique;not null"`
}
