package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	CurrentSlash string
}
  
func main() {
	log.Println("starting")
	var err error
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("db open err")
	} else {
		log.Println("Connected to the database")
	}

	db.AutoMigrate(&User{})
}
