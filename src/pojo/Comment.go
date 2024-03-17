package pojo

//评论
type Comment struct {
	ID            int    `gorm:"column:id" json:"ID"`
	UserId        int    `gorm:"column:userid" json:"UserId"`
	Content       string `gorm:"column:content" json:"Content"`
	LikeCount     int    `gorm:"column:likeCount" json:"LikeCount"`
	CreatTime     string `gorm:"column:creatTime" json:"CreatTime"`
	RootCommentId int    `gorm:"column:rootCommentId" json:"-"`
	ToCommentId   int    `gorm:"column:toCommentId" json:"ToCommentId"`
}

//评论前三条
type Comment_topthree struct {
	Comment
	SubComment []Comment `json:"SubComment"`
}
