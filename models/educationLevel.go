package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * Duration of Job model
 */
type EducationLevel struct {
  gorm.Model
  Name string `gorm:"type:varchar(50);unique;not null"` // Bachelor's, Master's, Doctorate...
}
