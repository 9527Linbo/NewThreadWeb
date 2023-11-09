package pojo

type PostListShare struct {
	Postid     int64  `gorm:"column:id" json:"articleId"`
	Title      string `gorm:"column:title" json:"title"`
	Content    string `gorm:"column:content" json:"content"`
	PostImgUrl string `gorm:"column:url" json:"articleImgUrl"`
}

type PostListNews struct {
	Postid  int64  `gorm:"column:id" json:"articleId"`
	Title   string `gorm:"column:title" json:"title"`
	Content string `gorm:"column:content" json:"content"`
}

type PostListActivity struct {
	Postid     int64  `gorm:"column:id" json:"articleId"`
	Title      string `gorm:"column:title" json:"title"`
	Content    string `gorm:"column:content" json:"content"`
	PostImgUrl string `gorm:"column:url" json:"articleImgUrl"`
}
