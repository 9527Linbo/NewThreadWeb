package pojo

import "time"

type HonoursTeam struct {
	HonoursTeam string `gorm:"column:teamname" json:"TeamName"`
	Number      int    `gorm:"column:number" json:"Number"`
	Rank        string `gorm:"column:rank" json:"Rank"`
}

type Honours struct {
	ID          int           `gorm:"column:id" json:"ID"`
	Name        string        `gorm:"column:name" json:"Name"`
	TeacherName string        `gorm:"column:teacher" json:"TeacherName"`
	Group       string        `gorm:"column:group" json:"Group"`
	StartTime   time.Time     `gorm:"column:start_time" json:"StartTime"`
	TeamMessage []HonoursTeam `gorm:"-" json:"HonoursTeams"`
}

type StudentHonours struct {
	Name        string `gorm:"column:name" json:"Name"`
	URL         string `gorm:"column:url" json:"URL"`
	Description string `gorm:"column:description" json:"Description"`
}

type Project struct {
	ID          int              `gorm:"column:id" json:"ID"`
	Name        string           `gorm:"column:name" json:"Name"`
	Description string           `gorm:"column:description" json:"Description"`
	Type        string           `gorm:"column:type" json:"type"`
	URL         []ImgURL         `gorm:"-" json:"URL"`
	Honours     []ProjectHonours `gorm:"-" json:"ProjectHonours"`
}

type ImgURL struct {
	URL string `gorm:"column:url" json:"URL"`
}

type ProjectHonours struct {
	Rank string `gorm:"column:rank" json:"Rank"`
	Time string `gorm:"column:time" json:"Time"`
}
