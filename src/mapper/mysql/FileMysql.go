package mapper

import (
	"errors"
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
