package models

import (
  "github.com/jinzhu/gorm"
)

type Language struct {
  gorm.Model
  Name string
}
