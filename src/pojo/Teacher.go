package pojo

type Teacher struct {
	Name     string      `gorm:"column:name" json:"Name"`
	URL      string      `gorm:"column:url" json:"URL"`
	Userid   int         `gorm:"column:user_id" json:"User_id"`
	Group    string      `gorm:"column:group" json:"Group"`
	Position interface{} `gorm:"-" json:"Position"`
}
