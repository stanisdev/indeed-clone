package models

import (
  "github.com/jinzhu/gorm"
)

type ExperienceYear struct {
  gorm.Model
  VacancyID uint
  ExperienceTypeID uint
  ExperienceYearID uint
}
