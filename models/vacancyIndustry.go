package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Vacancy model
 */
type VacancyIndustry struct {
  gorm.Model
  VacancyID uint `gorm:"not null"`
  IndustryID uint `gorm:"not null"`
}
