package models

import (
  "github.com/jinzhu/gorm"
)

/**
 * List of pofessions that required to Job
 */
type ExperienceProfession struct {
  gorm.Model
  Title string `gorm:"type:varchar(100);unique;not null"`
}
