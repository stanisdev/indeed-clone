package services

import (
  "github.com/jinzhu/gorm"
  "github.com/stanisdev/indeed-clone/models"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

type Product struct {
  gorm.Model
  Code string
  Price uint
}

func DbConnect() *gorm.DB {
  db, err := gorm.Open("mysql", "root:root@/indeed_clone?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("Failed to connect database")
  }
  db.LogMode(true)
  return db
}

func DbMigrate()  {
  con := DbConnect()
  con.AutoMigrate(
    &models.Country{},
    &models.City{},
    &models.Company{},
    &models.EmployeesCount{},
    &models.Phone{},
    &models.CompanyPhone{},
    &models.User{})

  con.Model(&models.City{}).AddForeignKey("country_id", "countries(id)", "CASCADE", "CASCADE")
  con.Model(&models.Company{}).AddForeignKey("city_id", "cities(id)", "CASCADE", "CASCADE")
  con.Model(&models.Company{}).AddForeignKey("employees_count_id", "employees_counts(id)", "CASCADE", "CASCADE")
  con.Model(&models.CompanyPhone{}).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE")
  con.Model(&models.CompanyPhone{}).AddForeignKey("phone_id", "phones(id)", "CASCADE", "CASCADE")
}
