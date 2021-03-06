package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Many to Many: Languages required for Job
 */
type VacancyLanguage struct {
  gorm.Model
  VacancyID uint `gorm:"not null"`
  LanguageID uint `gorm:"not null"`
}
