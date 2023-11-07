package pojo

type PostListShare struct {
	Postid     int    `gorm:"column:id" json:"Article_id"`
	Title      string `gorm:"column:title" json:"Title"`
	Content    string `gorm:"column:content" json:"Content"`
	URLArticle string `gorm:"column:url" json:"URLArticle"`
}

type PostListNews struct {
	Postid  int    `gorm:"column:id" json:"Article_id"`
	Title   string `gorm:"column:title" json:"Title"`
	Content string `gorm:"column:content" json:"Content"`
}

type PostListActivity struct {
	Postid     int    `gorm:"column:id" json:"Article_id"`
	Title      string `gorm:"column:title" json:"Title"`
	Content    string `gorm:"column:content" json:"Content"`
	URLArticle string `gorm:"column:url" json:"URLArticle"`
}
