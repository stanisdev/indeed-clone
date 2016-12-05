package models

import (
  "github.com/jinzhu/gorm"
)

type VacancyTag struct {
  gorm.Model
  VacancyID uint
  TagID uint
}
