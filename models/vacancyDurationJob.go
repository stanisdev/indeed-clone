package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Many to Many: Vacancy and Duration of Job
 */
type VacancyDurationJob struct {
  gorm.Model
  VacancyID uint `gorm:"not null"`
  DurationJobID uint `gorm:"not null"`
}
