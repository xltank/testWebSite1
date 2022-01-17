package main

import (
	"fmt"
	"website/db"
	"website/model"
)

func main() {
	db.InitMysql()

	var users []model.User
	r1 := db.Db.Find(&users)
	fmt.Println(r1)

	u := model.User{
		Name:       "Jason",
		Pass:       "123",
		Department: "test",
		Role:       "Admin",
		Email:      "jason@126.com",
	}
	r := db.Db.Create(&u)
	fmt.Println(r)
}
