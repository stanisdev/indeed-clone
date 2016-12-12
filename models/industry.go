package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Industry model
 */
type Industry struct {
  gorm.Model
  Name string `gorm:"type:varchar(50);unique;not null"` // Medicine, Hospitality
}
