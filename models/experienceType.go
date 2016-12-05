package models

import (
  "github.com/jinzhu/gorm"
)

type ExperienceType struct {
  gorm.Model
  Title string
}
