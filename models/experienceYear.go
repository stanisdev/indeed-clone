package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * List of numbers that reletaed to years experience
 */
type ExperienceYear struct {
  gorm.Model
  Value int `gorm:"unique;not null"` // 1, 2, 3 ...
}
