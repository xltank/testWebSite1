package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type User struct {
	Name       string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Pass       string `json:"pass,omitempty"`
	Role       string `json:"role,omitempty"` // max role: sa > admin > editor > member
	Department string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func UserList(ctx *gin.Context) {
	time.Sleep(123 * time.Millisecond)
	ctx.JSON(200, gin.H{
		"rtn": 0,
		"data": gin.H{
			"list": []User{
				User{},
				User{},
			},
		},
	})
}

func UserLogin(c *gin.Context) {
	// fmt.Println(c.Request.Body)
	u := User{}
	if e := c.ShouldBind(&u); e != nil {
		fmt.Println(e)
		c.JSON(400, gin.H{
			"rtn":    400,
			"errMsg": e.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"rtn": 0,
		"data": gin.H{
			"name": u.Name,
			"pass": u.Pass,
		},
	})
}
