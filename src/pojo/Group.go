package pojo

type Group struct {
	Id          int    `gorm:"column:id" json:"Group_id"`
	GroupName   string `gorm:"column:name" json:"GroupName"`
	Description string `gorm:"column:description" json:"Description"`
}
