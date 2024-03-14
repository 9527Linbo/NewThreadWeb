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
