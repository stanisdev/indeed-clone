package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Kind of experience model
 */
type ExperienceLevel struct {
  gorm.Model
  Name string `gorm:"type:varchar(50);unique;not null"` // "Mid Level", "Senior Level", "Entry Level"
}
