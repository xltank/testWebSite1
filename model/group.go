package model

type Group struct {
	Model
	Name  string `json:"name,omitempty"`
	Desc  string `json:"desc"`
	Users []User `json:"users" gorm:"many2many:user_group;"`
}

type GroupQueryParam struct {
	Keyword string `form:"keyword"`
	Offset  int    `form:"offset"`
	Limit   int    `form:"limit"`
}

type UserGroup struct {
	Model
	UserId  int    `json:"uid,omitempty" uri:"uid" `
	GroupId int    `json:"gid,omitempty" uri:"gid" `
	Role    string `json:"role,omitempty" uri:"role"`
}

/*func GroupCreate(group *Group) *Group {
	Db.Create(group)
	return group
}*/
