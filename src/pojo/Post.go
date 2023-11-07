package pojo

import "time"

type PostShare struct {
	Postid     int            `gorm:"column:id" json:"Article_id"`
	Title      string         `gorm:"column:title" json:"Title"`
	Content    string         `gorm:"column:content" json:"Content"`
	UserName   string         `gorm:"column:username" json:"Username"`
	UpdateTime time.Time      `gorm:"column:update_time" json:"Update_time"`
	URLUser    string         `gorm:"column:url" json:"URL_user"`
	URLArticle []PostShareUrl `gorm:"-" json:"URL_article"`
}

type PostShareUrl struct {
	Url string `gorm:"column:url" json:"URL"`
}
