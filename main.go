package main

import (
  // "net/http"
  "github.com/gin-gonic/gin"

  "github.com/danielwetan/gocrud/models"
	"github.com/danielwetan/gocrud/controllers"
)

func main() {
  r := gin.Default()

  models.Connect()

	r.GET("/books", controllers.GetBooks)
  r.GET("/books/:id", controllers.GetBook)
	r.POST("/books", controllers.CreateBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books:id", controllers.DeleteBook)

  r.Run()
}
