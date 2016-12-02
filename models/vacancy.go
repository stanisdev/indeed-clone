package models

import (
  "github.com/jinzhu/gorm"
  "time"
)

type Vacancy struct {
  gorm.Model
  Title string
  ExpiredDate time.Time
}
