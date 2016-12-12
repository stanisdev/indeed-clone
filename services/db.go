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
    &models.DurationJob{},
    &models.EducationLevel{},
    &models.Email{},
    &models.EmployeesCount{},
    &models.ExperienceLevel{},
    &models.ExperienceProfession{},
    &models.ExperienceYear{},
    &models.Industry{},
    &models.Language{},
    &models.PaymentRate{},
    &models.Phone{},
    &models.Tag{},
    &models.User{},
    &models.Vacancy{},
    &models.VacancyDuration{},
    &models.VacancyEducationLevel{},)

  con.Model(&models.City{}).AddForeignKey("country_id", "countries(id)", "CASCADE", "CASCADE")
  con.Model(&models.Company{}).AddForeignKey("city_id", "cities(id)", "CASCADE", "CASCADE")
  con.Model(&models.Company{}).AddForeignKey("employees_count_id", "employees_counts(id)", "CASCADE", "CASCADE")
  con.Model(&models.Email{}).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE")
  con.Model(&models.VacancyDuration{}).AddForeignKey("vacancy_id", "vacancies(id)", "CASCADE", "CASCADE")
  con.Model(&models.VacancyDuration{}).AddForeignKey("duration_job_id", "duration_jobs(id)", "CASCADE", "CASCADE")

  con.Model(&models.VacancyEducationLevel{}).AddForeignKey("vacancy_id", "vacancies(id)", "CASCADE", "CASCADE")
  con.Model(&models.VacancyEducationLevel{}).AddForeignKey("education_level_id", "education_levels(id)", "CASCADE", "CASCADE")
}
