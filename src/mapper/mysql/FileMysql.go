package mapper

import (
	"NewThread/src/pojo"
	"errors"

	"gorm.io/gorm"
)

type FileMysql struct{}

func NewFileMysql() *FileMysql {
	return &FileMysql{}
}

func (c *FileMysql) InsertFileMesg(filename string, fileuuid string, username string, url string) error {
	m := Db.Exec("INSERT INTO t_files VALUES(null,?,?,?,?,NOW(),NOW())", filename, username, fileuuid, url)
	rowsaffected := m.RowsAffected
	if rowsaffected == 0 {
		return errors.New("Insert---File---Mesg---Error")
	}
	return nil
}

func (c *FileMysql) File(fileuuid string) (pojo.FileList, error) {
	var m pojo.FileList
	err := Db.Raw("SELECT filename,username,fileuuid,url from t_files where fileuuid = ?", fileuuid).Scan(&m).Error
	if err != nil {
		return pojo.FileList{}, err
	}
	return m, nil
}

func (c *FileMysql) InsertIconMesg(url string, fileuuid string, userid int, db *gorm.DB) error {
	m := db.Exec("UPDATE t_user SET url = ?, iconname = ?, update_time = NOW() WHERE id = ?", url, fileuuid, userid)
	rowsaffected := m.RowsAffected
	if rowsaffected == 0 {
		return errors.New("UPDATE---Icon---Mesg---Error")
	}
	return nil
}

func (c *FileMysql) SearchFileUUIDById(userid int) (string, error) {
	var fileuuid string
	err := Db.Raw("SELECT iconname FROM t_user WHERE id =?", userid).Scan(&fileuuid).Error
	if err != nil {
		return "", err
	}
	return fileuuid, nil
}

func (c *FileMysql) InsertPostImgMesg(url string, postid int, db *gorm.DB) error {
	m := db.Exec("INSERT INTO t_articleimage VALUES(null,?,?,NOW(),NOW())", url, postid)
	rowsaffected := m.RowsAffected
	if rowsaffected == 0 {
		return errors.New("Insert---postimg---Mesg---Error")
	}
	return nil
}

func (c *FileMysql) InsertHonourImgMesg(url string, honourid int, db *gorm.DB) error {
	m := db.Exec("INSERT INTO t_projectimage VALUES(NULL,?,?,NOW(),NOW())", url, honourid)
	rowsaffected := m.RowsAffected
	if rowsaffected == 0 {
		return errors.New("Insert---postimg---Mesg---Error")
	}
	return nil
}
