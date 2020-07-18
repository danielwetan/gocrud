package controllers

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/danielwetan/gocrud/models"
)

// Get all books
func GetBooks(c *gin.Context) {
  var books []models.Book
  models.DB.Find(&books)

  c.JSON(http.StatusOK, gin.H{"data": books})
}

// Get single book
func GetBook(c *gin.Context) {
  var book models.Book

  if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found!"})
    return
  }

  c.JSON(http.StatusOK, gin.H{"data": book})
}

// Create new book
func CreateBook(c *gin.Context) {
  // Validate input
  var input models.CreateBookInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
  }

  book := models.Book{Title: input.Title, Author: input.Author}
  models.DB.Create(&book)

  c.JSON(http.StatusOK, gin.H{"data": book})
}

// Update a book
func UpdateBook(c *gin.Context) {
  var book models.Book
  if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
    return
  }

  // Validate input
  var input models.UpdateBookInput
  if err := c.ShouldBindJSON(&input); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
    return
  }

  models.DB.Model(&book).Updates(input)

  c.JSON(http.StatusOK, gin.H{"data": book})
}

// Delete a book
func DeleteBook(c *gin.Context) {
  var book models.Book
  if err := models.DB.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": "Data not found"})
    return
  }

  models.DB.Delete(&book)

  c.JSON(http.StatusOK, gin.H{"data": true})
}
