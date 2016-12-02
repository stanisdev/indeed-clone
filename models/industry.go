package models

import (
  "github.com/jinzhu/gorm"
)

type Industry struct {
  gorm.Model
  Name string
}
