package pojo

type UserIcon struct {
	Userid int    `gorm:"column:id" json:"User_id"`
	Name   string `gorm:"column:username" json:"Name"`
	URL    string `gorm:"column:url" json:"URL"`
}

type RecvUserMsg struct {
	Account string `json:"account" form:"account"`
	Pwd     string `json:"-" form:"pwd"`
}
