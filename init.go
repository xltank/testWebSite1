package main

import (
	"fmt"
	"website/db"
	"website/router"
)

func main() {
	db.InitMysql()

	var users []router.User
	r1 := db.Db.Find(&users)
	fmt.Println(r1)

	u := router.User{
		Name:       "Jason",
		Pass:       "123",
		Department: "test",
		Role:       "Admin",
		Email:      "jason@126.com",
	}
	r := db.Db.Create(&u)
	fmt.Println(r)
}
