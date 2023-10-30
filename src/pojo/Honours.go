package pojo

import "time"

type HonoursTeam struct {
	HonoursTeam string `gorm:"column:TeamName" json:"HonoursTeam"`
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
