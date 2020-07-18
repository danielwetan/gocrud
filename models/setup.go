package models

import (
  "github.com/jinzhu/gorm"
  _ "github.com/go-sql-driver/mysql"
)

var DB *gorm.DB

func Connect() {
  // create db connection
  database, err := gorm.Open("mysql", "root@tcp(127.0.0.1:3306)/gocrud")

  // check connection
  if err != nil {
    panic("Failed to connnect to database!")
  }

  // migrate the database schema
  database.AutoMigrate(&Book{})

  // we will use this variable in controller to access db
  DB = database
}
