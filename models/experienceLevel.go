package models

import (
  "github.com/jinzhu/gorm"
)

type ExperienceLevel struct {
  gorm.Model
  Name string `gorm:"type:varchar(50);unique;not null"`
}
