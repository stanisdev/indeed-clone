package middlewares

import (
  "github.com/gin-gonic/gin"
  "github.com/stanisdev/indeed-clone/services"
)

func DbConnect(c *gin.Context) {
  c.Set("db", services.DbConnect())
  c.Next()
}
