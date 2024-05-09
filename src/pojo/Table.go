package pojo

import "time"

type T_user struct {
	Id         int       `gorm:"column:id" `
	Username   string    `gorm:"column:username" `
	Account    string    `gorm:"column:account"`
	Pwd        string    `gorm:"column:password"`
	Url        string    `gorm:"column:url"`
	Iconname   string    `gorm:"column:iconname"`
	CreatTime  time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

type T_teacher struct {
	Id           int       `gorm:"column:id" form:"id" `
	Name         string    `gorm:"column:name" form:"name" `
	Userid       int       `gorm:"column:user_id" form:"account"`
	Position     string    `gorm:"column:position" form:"position"`
	Pdescription string    `gorm:"column:pdescription" form:"pdescription"`
	CreatTime    time.Time `gorm:"column:create_time"`
	UpdateTime   time.Time `gorm:"column:update_time"`
}

type T_student struct {
	Id           int       `form:"id" `
	Name         string    `gorm:"column:name" form:"name" `
	Userid       int       `gorm:"column:user_id" form:"account"`
	GroupId      int       `gorm:"group_id" form:"group_id"`
	Class        int       `gorm:"class" form:"class"`
	Position     string    `gorm:"column:position" form:"position"`
	Pdescription string    `gorm:"column:pdescription" form:"pdescription"`
	CreatTime    time.Time `gorm:"column:create_time"`
	UpdateTime   time.Time `gorm:"column:update_time"`
}

type T_graduate struct {
	Id          int       `form:"id" `
	Name        string    `gorm:"column:name" form:"name" `
	Userid      int       `gorm:"column:user_id" form:"account"`
	Url         string    `gorm:"column:URL" form:"URL`
	Description string    `gorm:"column:description" form:"description"`
	CreatTime   time.Time `gorm:"column:create_time"`
	UpdateTime  time.Time `gorm:"column:update_time"`
}

type T_project struct {
	Id          int       `form:"id" `
	Name        string    `gorm:"column:name" form:"name" `
	Description string    `gorm:"column:description" form:"description"`
	Type        string    `gorm:"column:type" form:"type"`
	IsMilestone byte      `gorm:"column:Ismilestone" form:"Ismilestone"`
	CreatTime   time.Time `gorm:"column:create_time"`
	UpdateTime  time.Time `gorm:"column:update_time"`
}

type T_projectawards struct {
	Id         int       `form:"id" `
	ProjectId  int       `gorm:"column:project_id"`
	Rank       string    `gorm:"column:rank"`
	Time       string    `gorm:"column:time"`
	CreatTime  time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

type T_race struct {
	Id         int       `form:"id" `
	Name       string    `gorm:"column:name"`
	TeacherId  int       `gorm:"column:teacher_id"`
	GroupId    int       `gorm:"column:group_id"`
	Type       string    `gorm:"column:type"`
	StartTime  string    `gorm:"column:start_time"`
	EndTime    string    `gorm:"column:end_time"`
	CreatTime  time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

type T_raceawards struct {
	Id         int       `form:"id" `
	RaceId     int       `gorm:"column:race_id"`
	Rank       string    `gorm:"column:rank"`
	TeamName   string    `gorm:"column:TeamName"`
	Number     int       `gorm:"column:number"`
	CreatTime  time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

type T_article struct {
	Id         int       `form:"id" `
	Type       int       `gorm:"column:type" json:"type"`
	Title      string    `gorm:"column:title" json:"title"`
	Content    string    `gorm:"column:content" json:"content"`
	UserId     int       `gorm:"column:user_id" json:"userid"`
	CreatTime  time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

type T_articleimage struct {
	Id         int       `form:"id" `
	Url        string    `gorm:"column:url"`
	UserId     int       `gorm:"column:article_id" json:"id"`
	CreatTime  time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}
