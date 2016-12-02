package models

import (
  "github.com/jinzhu/gorm"
)

type Country struct {
  gorm.Model
  Name string
}
