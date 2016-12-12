package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Tags model
 */
type Tag struct {
  gorm.Model
  Name string `gorm:"type:varchar(30);unique;not null"`
}
