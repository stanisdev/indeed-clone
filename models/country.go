package models

/**
 * Country model
 */
type Country struct {
  ID uint `gorm:"primary_key"`
  Name string `gorm:"type:varchar(50);unique;not null"`
}
