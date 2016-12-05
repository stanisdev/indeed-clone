package models

import (
  "github.com/jinzhu/gorm"
)

type EducationLevel struct {
  gorm.Model
  Name string `gorm:"type:varchar(50);unique;not null"`
}
