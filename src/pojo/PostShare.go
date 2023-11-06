package pojo

import "time"

type PostShare struct {
	Postid     string      `gorm:"column:article_id" json:"Article_id"`
	Title      string      `gorm:"column:title" json:"Title"`
	Content    string      `gorm:"column:content" json:"Content"`
	Userid     int         `gorm:"column:user_id" json:"User_id"`
	UpdateTime time.Time   `gorm:"column:update_time" json:"Update_time"`
	User       interface{} `gorm:"-" json:"User"`
	URL        string      `gorm:"column:url" json:"URL"`
}
