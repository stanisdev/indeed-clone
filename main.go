package main

import (
  "github.com/gin-gonic/gin"
  "github.com/stanisdev/indeed-clone/services"
  "github.com/stanisdev/indeed-clone/handlers"
  "github.com/stanisdev/indeed-clone/middlewares"
  "os"
)

func main() {
  if len(os.Getenv("DB_MIGRATE")) > 0 {
    services.DbMigrate()
    return
  }
  router := gin.Default()
  htmlRender := services.NewRender()
  htmlRender.Debug = gin.IsDebugging()
  htmlRender.Layout = "layouts/default"
  router.HTMLRender = htmlRender.Create()

  router.Use(middlewares.DbConnect)

  router.GET("/ping", handlers.MainIndex)
  router.Run()
}
