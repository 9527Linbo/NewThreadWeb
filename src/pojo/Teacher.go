package pojo

type Teacher struct {
	Id       int         `gorm:"column:id" json:"Teacher_id"`
	Name     string      `gorm:"column:name" json:"GroupName"`
	Userid   int         `gorm:"column:user_id" json:"User_id"`
	Group    string      `gorm:"column:group" json:"Group"`
	Position interface{} `gorm:"-" json:"Position"`
}
