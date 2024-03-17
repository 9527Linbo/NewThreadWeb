package pojo

type UserIcon struct {
	Userid int    `gorm:"column:id" json:"User_id"`
	Name   string `gorm:"column:username" json:"Name"`
	URL    string `gorm:"column:url" json:"URL"`
}
