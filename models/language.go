package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Additional languages required for perfomance of duties
 */
type Language struct {
  gorm.Model
  Name string `gorm:"type:varchar(50);unique;not null"`
}
