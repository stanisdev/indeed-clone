package models

import (
  "github.com/jinzhu/gorm"
)

type EmployeesCount struct {
  gorm.Model
  Value string
}
