package models

import (
  "github.com/jinzhu/gorm"
  "time"
)

type ExperienceYear struct {
  gorm.Model
  Value int
}
