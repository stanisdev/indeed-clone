package models

import (
  "github.com/jinzhu/gorm"
  "time"
)

type Vacancy struct {
  gorm.Model
  Title string
  ExpiredDate time.Time
  Description string `gorm:"type:text;not null"`
  SalarySize int
  IndustryID uint
  ExperienceLevelID uint
  PaymentRateID uint
  EducationLevelID uint
}
