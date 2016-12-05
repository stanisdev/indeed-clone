package models

import (
  "github.com/jinzhu/gorm"
)

type City struct {
  gorm.Model
  Name string `gorm:"type:varchar(50);unique;not null"`
  CountryId uint `gorm:"not null"`
}
