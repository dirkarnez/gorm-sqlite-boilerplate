package main

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
}
  
func main() {
	var err error
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("db open err")
	} else {
		log.Println("Connected to the database")
	}

	if db.Migrator().HasTable(&User{}) {
		log.Fatal("has `users table, dropping...`")
		db.Migrator().DropTable(&User{})
	}

	db.Migrator().CreateTable(&User{})

	db.Create(&[]User{
		{Name: "jinzhu1" }, 
		{Name: "jinzhu2" },
		{Name: "jinzhu3" },
	})
}
