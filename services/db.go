package services

import (
  "github.com/jinzhu/gorm"
  "github.com/stanisdev/indeed-clone/models"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "fmt"
)

const importingError = "Error while importing fixtures"

/**
 * Connecting to DB
 */
func DbConnect() *gorm.DB {
  db, err := gorm.Open("mysql", "root:root@/indeed_clone?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("Failed to connect database")
  }
  db.LogMode(true)
  return db
}

/**
 * Import fixtures
 */
func ImportFixtures() {
  con := DbConnect()
  tx := con.Begin()

  fixtures := getFixtures()
  for id, name := range fixtures["Countries"].(map[uint]string) {
    country := models.Country{
      ID: id,
      Name: name,
    }
    if err := tx.Create(&country).Error; err != nil {
      tx.Rollback()
      fmt.Println(importingError)
      return
    }
  }
  for id, city := range fixtures["Cities"].(map[uint][]interface{}) {
    cityInstance := models.City{
      ID: id,
      Name: city[1].(string),
      CountryID: uint(city[0].(int)),
    }
    if err := tx.Create(&cityInstance).Error; err != nil {
      tx.Rollback()
      fmt.Println(importingError)
      return
    }
  }
  for id, employeesCount := range fixtures["EmployeesCount"].(map[uint]string) {
    employeesCountInstance := models.EmployeesCount{
      ID: id,
      Value: employeesCount,
    }
    if err := tx.Create(&employeesCountInstance).Error; err != nil {
      tx.Rollback()
      fmt.Println(importingError)
      return
    }
  }
  tx.Commit()
  fmt.Println("All fixtures have been imported successfully")
}

/**
 * Return fixtures
 */
func getFixtures() map[string]interface{} {
  return map[string]interface{} {
    "Countries": map[uint]string{1: "USA", 2: "Canada"},
    "Cities": map[uint][]interface{}{
      1: []interface{}{1, "Chicago"},
      2: []interface{}{1, "New York"},
      3: []interface{}{2, "Toronto"},
    },
    "EmployeesCount": map[uint]string{
      1: "1 - 49",
      2: "50 - 149",
      3: "150 - 249",
      4: "250 - 499",
      5: "500 - 749",
      6: "750 - 999",
      7: "1000 +",
    },
  }
}

/**
 * Migration
 */
func DbMigrate() {
  con := DbConnect()
  tx := con.Begin()
  err := tx.AutoMigrate(
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
    &models.VacancyDurationJob{},
    &models.VacancyEducationLevel{},
    &models.VacancyExperience{},
    &models.VacancyExperienceLevel{},
    &models.VacancyIndustry{},
    &models.VacancyLanguage{},
    &models.VacancyTag{},).Error

  if err != nil {
    tx.Rollback()
    fmt.Println("Error while migration")
    return
  }

  foreignKeys := make([]func() error, 0, 22)
  foreignKeys = append(foreignKeys, func() error {
    return tx.Model(&models.City{}).AddForeignKey("country_id", "countries(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.Company{}).AddForeignKey("city_id", "cities(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.Company{}).AddForeignKey("employees_count_id", "employees_counts(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.Email{}).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.VacancyDurationJob{}).AddForeignKey("vacancy_id", "vacancies(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.VacancyDurationJob{}).AddForeignKey("duration_job_id", "duration_jobs(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.VacancyEducationLevel{}).AddForeignKey("vacancy_id", "vacancies(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.VacancyEducationLevel{}).AddForeignKey("education_level_id", "education_levels(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.VacancyExperience{}).AddForeignKey("vacancy_id", "vacancies(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.VacancyExperience{}).AddForeignKey("experience_profession_id", "experience_professions(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.VacancyExperience{}).AddForeignKey("experience_year_id", "experience_years(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.VacancyExperienceLevel{}).AddForeignKey("vacancy_id", "vacancies(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.VacancyExperienceLevel{}).AddForeignKey("experience_level_id", "experience_levels(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.VacancyIndustry{}).AddForeignKey("vacancy_id", "vacancies(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.VacancyIndustry{}).AddForeignKey("industry_id", "industries(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.VacancyLanguage{}).AddForeignKey("vacancy_id", "vacancies(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.VacancyLanguage{}).AddForeignKey("language_id", "languages(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.VacancyTag{}).AddForeignKey("vacancy_id", "vacancies(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.VacancyTag{}).AddForeignKey("tag_id", "tags(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.Vacancy{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.Vacancy{}).AddForeignKey("payment_rate_id", "payment_rates(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.Phone{}).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE").Error
  })
  foreignKeys = append(foreignKeys, func() error {
    return con.Model(&models.Vacancy{}).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE").Error
  })
  for _, function := range foreignKeys {
    if err := function(); err != nil {
      tx.Rollback()
      fmt.Println("Error while migration")
      return
    }
  }
  tx.Commit()
  fmt.Println("Done")
}
