package pojo

type User struct {
	Id       int    `gorm:"column:id" json:"user_id"`
	Username string `gorm:"column:username" json:"username"`
}
