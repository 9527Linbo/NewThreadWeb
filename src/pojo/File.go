package pojo

//文件
type FileList struct {
	Filename   string `gorm:"column:filename" json:"Name"`
	Fileuuid   string `gorm:"column:fileuuid" json:"FileUUID"`
	AtOSS      bool   `gorm:"column:atOSS" json:"AtOSS"`
	Username   string `gorm:"column:username" json:"UserName"`
	UpdateTime string `gorm:"-" json:"UpdateTime"`
	Size       string `gorm:"-" json:"size"`
}
