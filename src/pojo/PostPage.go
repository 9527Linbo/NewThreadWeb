package pojo

type PageShare struct {
	CurrentPage  int `gorm:"-" json:"Current_page"`
	TotalPageNum int `gorm:"column:sum" json:"Total_page"`
	PostList     *[]PostListShare
}

type PageNews struct {
	CurrentPage  int `gorm:"-" json:"Current_page"`
	TotalPageNum int `gorm:"column:sum" json:"Total_page"`
	PostList     *[]PostListNews
}

type PageActivity struct {
	CurrentPage  int `gorm:"-" json:"Current_page"`
	TotalPageNum int `gorm:"column:sum" json:"Total_page"`
	PostList     *[]PostListActivity
}
