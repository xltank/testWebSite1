package model

type Group struct {
	Model
	Name string `json:"name,omitempty"`
	Desc string `json:"desc"`
	//Users []User `json:"users" gorm:"many2many:user_group;"`
}

type GroupQueryParam struct {
	Keyword string `form:"keyword"`
	Offset  int    `form:"offset"`
	Limit   int    `form:"limit"`
}

/*
UserGroup
CREATE TABLE `user_group` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `group_id` int(11) NOT NULL,
  `createdAt` datetime NOT NULL,
  `updatedAt` datetime NOT NULL,
  `role` varchar(10) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_id_group_id` (`user_id`,`group_id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8mb4;
*/
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
