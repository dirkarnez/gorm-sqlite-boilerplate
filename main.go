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
		user.Saving = user.Saving.Add(decimal.NewFromFloat(0.01))
	}

	fmt.Println("Going to persist:", user.Saving)
	db.Save(&user)

	var userReQuery User
	db.First(&userReQuery, 1)
	log.Println("Saved userReQuery.Saving:", userReQuery.Saving)

	x := decimal.NewFromInt(10).Div(decimal.NewFromInt(3)) // 10/3
	y := decimal.NewFromInt(10).Div(decimal.NewFromInt(3)) // 10/3
	userReQuery.Saving = userReQuery.Saving.Sub(x.Round(2))
	userReQuery.Saving = userReQuery.Saving.Sub(y.Round(2))
	db.Save(&userReQuery)

	var userReQuery2 User
	db.First(&userReQuery2, 1)
	log.Println("Saved userReQuery2.Saving:", userReQuery2.Saving.Round(2)) // 10 - 3.33 - 3.33 = 3.34
	userReQuery2.Saving = userReQuery2.Saving.Round(2).Add(decimal.NewFromFloat32(0.005).Round(2))
	db.Save(&userReQuery2)

	var userReQuery3 User
	db.First(&userReQuery3, 1)
	log.Println("Saved userReQuery3.Saving:", userReQuery3.Saving.Round(2))

	log.Println("end")
}
