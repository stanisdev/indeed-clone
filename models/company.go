package models

import (
  "github.com/jinzhu/gorm"
)

type Company struct {
  gorm.Model
  Name string
  CityID uint
  Url string
  EmployeesCountID uint
}
