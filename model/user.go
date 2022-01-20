package model

import (
	. "website/db"
)

type User struct {
	Model
	Name       string  `json:"name,omitempty" binding:"required"`
	Email      string  `json:"email,omitempty" binding:"required"`
	Pass       string  `json:"pass,omitempty"`
	Role       string  `json:"role,omitempty"` // max role: sa > admin > editor > member
	Department string  `json:"department,omitempty"`
	Groups     []Group `json:"groups" binding:"required" gorm:"many2many:user_group;"`
}

type UserQueryParam struct {
	Keyword string `form:"keyword"`
	Offset  int    `form:"offset"`
	Limit   int    `form:"limit"`
}

type UserLoginParam struct {
	Email string `binding:"required"`
	Pass  string `binding:"required"`
}

/*func UserGetOneByEmail(email string) *User {
	var user User
	db.Db.Where(map[string]string{email: email}).Find(&user)

	return &user
}*/

func UserCreate(user *User) *User {
	Db.Create(&user)
	return user
}
