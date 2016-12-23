package models

/**
 * How many employees? model
 */
type EmployeesCount struct {
  ID uint `gorm:"primary_key"`
  Value string `gorm:"type:varchar(50);unique;not null"` // "1 - 49", "50 - 149", "150 - 249"
}
