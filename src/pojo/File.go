package pojo

//文件
type FileList struct {
	Filename   string `gorm:"column:filename" json:"FileName"`
	Fileuuid   string `gorm:"column:fileuuid" json:"FileUUID"`
	URL        string `gorm:"column:url" json:"URL"`
	Username   string `gorm:"column:username" json:"UserName"`
	UpdateTime string `gorm:"-" json:"UpdateTime"`
	Size       string `gorm:"-" json:"size"`
}
