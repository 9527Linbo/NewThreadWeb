package pojo

type PageShare struct {
	CurrentPage  int              `gorm:"-" json:"currentPage"`
	TotalPageNum int              `gorm:"column:sum" json:"totalPage"`
	PostList     *[]PostListShare `json:"postList"`
}

type PageNews struct {
	CurrentPage  int             `gorm:"-" json:"currentPage"`
	TotalPageNum int             `gorm:"column:sum" json:"totalPage"`
	PostList     *[]PostListNews `json:"postList"`
}

type PageActivity struct {
	CurrentPage  int                 `gorm:"-" json:"currentPage"`
	TotalPageNum int                 `gorm:"column:sum" json:"totalPage"`
	PostList     *[]PostListActivity `json:"postList"`
}
