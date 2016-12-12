package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Many to Many: Vacancy and Level of Experience
 */
type VacancyExperienceLevel struct {
  gorm.Model
  VacancyID uint `gorm:"not null"`
  ExperienceLevelID uint `gorm:"not null"`
}
