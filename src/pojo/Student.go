package pojo

type Student struct {
	Name     string      `gorm:"column:name" json:"Name"`
	URL      string      `gorm:"column:url" json:"URL"`
	Userid   int         `gorm:"column:user_id" json:"User_id"`
	Group    string      `gorm:"column:group" json:"Group"`
	Year     int         `gorm:"column:year" json:"year"`
	Position interface{} `gorm:"-" json:"Position"`
}

type Year struct {
	Year int `form:"year"`
}
