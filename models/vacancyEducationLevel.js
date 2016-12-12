package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Many to Many: Vacancy and Level of Education
 */
type VacancyEducationLevel struct {
  gorm.Model
  VacancyID uint `gorm:"not null"`
  EducationLevelID uint `gorm:"not null"`
}
