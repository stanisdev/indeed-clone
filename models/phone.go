package models

import (
  "github.com/jinzhu/gorm"
)

type Phone struct {
  gorm.Model
  Number string
}
