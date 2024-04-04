package pojo

import "time"

type T_user struct {
	Id         int       `gorm:"column:id" `
	Username   string    `gorm:"column:username" `
	Account    string    `gorm:"column:account"`
	Pwd        string    `gorm:"column:password"`
	CreatTime  time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

type T_teacher struct {
	Id           int       `gorm:"column:id" form:"id" `
	Name         string    `gorm:"column:name" form:"name" `
	Userid       int       `gorm:"column:user_id" form:"account"`
	Position     string    `gorm:"column:position" form:"position"`
	Pdescription string    `gorm:"column:pdescription" form:"pdescription"`
	CreatTime    time.Time `gorm:"column:create_time"`
	UpdateTime   time.Time `gorm:"column:update_time"`
}

type T_student struct {
	Id           int       `form:"id" `
	Name         string    `gorm:"column:name" form:"name" `
	Userid       int       `gorm:"column:user_id" form:"account"`
	GroupId      int       `gorm:"group_id" form:"group_id"`
	Class        int       `gorm:"class" form:"class"`
	Position     string    `gorm:"column:position" form:"position"`
	Pdescription string    `gorm:"column:pdescription" form:"pdescription"`
	CreatTime    time.Time `gorm:"column:create_time"`
	UpdateTime   time.Time `gorm:"column:update_time"`
}

type T_graduate struct {
	Id          int       `form:"id" `
	Name        string    `gorm:"column:name" form:"name" `
	Userid      int       `gorm:"column:user_id" form:"account"`
	Url         string    `gorm:"column:URL" form:"URL`
	Description string    `gorm:"column:description" form:"description"`
	CreatTime   time.Time `gorm:"column:create_time"`
	UpdateTime  time.Time `gorm:"column:update_time"`
}
