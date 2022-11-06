package main

import (
	"fmt"
	"log"

	"github.com/shopspring/decimal"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string
	Saving decimal.Decimal `gorm:"type:decimal(13,6);"`
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
		log.Println("has `users table, dropping...`")
		db.Migrator().DropTable(&User{})
	}

	db.Migrator().CreateTable(&User{})

	var target = User{Name: "jinzhu1", Saving: decimal.NewFromInt(0)}
	target.ID = 1
	db.Create(&[]User{
		target,
		{Name: "jinzhu2"},
		{Name: "jinzhu3"},
	})

	log.Println("start")
	var user User
	db.First(&user, 1)
	log.Println("user.Saving:", user.Saving)

	for i := 0; i < 1000; i++ {
		user.Saving = user.Saving.Add(decimal.NewFromFloat(.01))
	}

	fmt.Println("Going to persist:", user.Saving)
	db.Save(&user)
	var userReQuery User
	db.First(&userReQuery, 1)
	log.Println("Saved userReQuery.Saving:", userReQuery.Saving)

	log.Println("end")
}
