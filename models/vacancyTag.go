package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Many to Many: Tags fot Job
 */
type VacancyTag struct {
  gorm.Model
  VacancyID uint `gorm:"not null"`
  TagID uint `gorm:"not null"`
}
