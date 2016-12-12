package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Duration of Job model
 */
type DurationJob struct {
  gorm.Model
  Name string `gorm:"type:varchar(50);unique;not null"` // Full-time, part-time...
}
