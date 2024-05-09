package pojo

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
	StartTime   string        `gorm:"column:start_time" json:"StartTime"`
	TeamMessage []HonoursTeam `gorm:"-" json:"HonoursTeams"`
}

type StudentHonours struct {
	ID          string `gorm:"column:id" json:"ID"`
	Name        string `gorm:"column:name" json:"Name"`
	URL         string `gorm:"column:url" json:"URL"`
	Description string `gorm:"column:description" json:"Description"`
}

type Project struct {
	ID          int              `gorm:"column:id"`
	Name        string           `gorm:"column:name" json:"name"`
	Description string           `gorm:"column:description" json:"description"`
	Type        string           `gorm:"column:type" json:"type"`
	URL         []ImgURL         `gorm:"-" json:"url"`
	Honours     []ProjectHonours `gorm:"-" json:"projectHonours"`
}

type ImgURL struct {
	URL string `gorm:"column:url" json:"URL"`
}

type ProjectHonours struct {
	Rank string `gorm:"column:rank" json:"rank"`
	Time string `gorm:"column:time" json:"time"`
}
