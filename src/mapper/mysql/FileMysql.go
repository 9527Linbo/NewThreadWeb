package mapper

import (
	"NewThread/src/pojo"
	"errors"
	"fmt"
)

type FileMysql struct{}

func NewFileMysql() *FileMysql {
	return &FileMysql{}
}

func (c *FileMysql) InsertFileMesg(filename string, fileuuid string, username string) error {
	m := Db.Exec("INSERT INTO t_files VALUES(null,?,?,?,FALSE,NOW(),NOW())", filename, username, fileuuid)
	rowsaffected := m.RowsAffected
	if rowsaffected == 0 {
		return errors.New("Insert---File---Mesg---Error")
	}
	return nil
}

func (c *FileMysql) File(fileuuid string) (pojo.FileList, error) {
	var m pojo.FileList
	err := Db.Raw("SELECT filename,username,fileuuid,atOSS from t_files where fileuuid = ?", fileuuid).Scan(&m).Error
	fmt.Print(m)
	if err != nil {
		return pojo.FileList{}, err
	}
	return m, nil
}

func (c *FileMysql) InsertIconMesg(url string, fileuuid string, userid int) error {
	m := Db.Exec("INSERT INTO t_imageuser VALUES(null,?,?,?,NOW(),NOW())", url, fileuuid, userid)
	rowsaffected := m.RowsAffected
	if rowsaffected == 0 {
		return errors.New("Insert---Icon---Mesg---Error")
	}
	return nil
}

func (c *FileMysql) SearchFileUUIDById(userid int) (string, error) {
	var fileuuid string
	err := Db.Raw("SELECT imagename FROM t_imageuser WHERE user_id =?", userid).Scan(&fileuuid).Error
	if err != nil {
		return "", err
	}
	return fileuuid, nil
}

func (c *FileMysql) DeleteIcon(fileuuid string) error {
	m := Db.Exec("DELETE FROM t_imageuser WHERE imagename = ?", fileuuid)
	rowsaffected := m.RowsAffected
	if rowsaffected == 0 {
		return errors.New("Delete---Icon---Mesg---Error")
	}
	return nil
}
