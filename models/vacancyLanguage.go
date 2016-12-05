package models

import (
  "github.com/jinzhu/gorm"
)

type VacancyLanguage struct {
  gorm.Model
  VacancyID uint
  LanguageID uint
}
