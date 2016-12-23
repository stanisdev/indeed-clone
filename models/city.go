package models

/**
 * City model
 */
type City struct {
  ID uint `gorm:"primary_key"`
  Name string `gorm:"type:varchar(50);unique;not null"`
  CountryID uint `gorm:"not null"`
}
