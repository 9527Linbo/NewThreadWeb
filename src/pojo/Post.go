package pojo

import "time"

type Post struct {
	Postid     int64     `gorm:"column:id" json:"articleId"`
	Title      string    `gorm:"column:title" json:"title"`
	Content    string    `gorm:"column:content" json:"content"`
	UpdateTime time.Time `gorm:"column:update_time" json:"updateTime"`
	CreateTime time.Time `gorm:"column:create_time" json:"createTime"`
	UserName   string    `gorm:"column:username" json:"username"`
	UserImgUrl string    `gorm:"column:url" json:"userImgUrl"`
	PostImgUrl *[]Url    `gorm:"-" json:"articleImgUrl"`
}

type Url struct {
	Url string `gorm:"column:url" json:"Url"`
}
