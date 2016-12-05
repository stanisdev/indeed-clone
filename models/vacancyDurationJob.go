package models

import (
  "github.com/jinzhu/gorm"
)

type VacancyDuration struct {
  gorm.Model
  VacancyID uint
  DurationJobID uint
}
