package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Country model
 */
type Country struct {
  gorm.Model
  Name string `gorm:"type:varchar(50);unique;not null"`
}
