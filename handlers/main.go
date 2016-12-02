package handlers

import (
  "github.com/gin-gonic/gin"
  // "github.com/jinzhu/gorm"
  // "github.com/stanisdev/indeed-clone/models"
  "net/http"
)

/**
 * Index page
 */
func MainIndex(c *gin.Context) {
  // db := c.MustGet("db").(*gorm.DB)
  // var product models.Product
  // db.First(&product, 1)
  c.HTML(http.StatusOK, "main/index", gin.H{
    "title": "Hello there",
  })
}
