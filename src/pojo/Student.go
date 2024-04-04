package pojo

type Student struct {
	Name     string      `gorm:"column:name" json:"Name"`
	URL      string      `gorm:"column:url" json:"URL"`
	Userid   int         `gorm:"column:user_id" json:"User_id"`
	Group    string      `gorm:"column:group" json:"Group"`
	Class    int         `gorm:"column:Class" json:"Class"`
	Position interface{} `gorm:"-" json:"Position"`
}

type Class struct {
	Class int `form:"class"`
}
