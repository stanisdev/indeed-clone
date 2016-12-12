package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Many to Many (triple relationship): Vacancy + Profession + Year of Experience
 */
type VacancyExperience struct {
  gorm.Model
  VacancyID uint `gorm:"not null"`
  ExperienceProfessionID uint `gorm:"not null"`
  ExperienceYearID uint `gorm:"not null"`
}
