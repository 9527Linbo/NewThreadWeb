package mapper

import (
	"NewThread/src/pojo"
	"errors"

	"gorm.io/gorm"
)

type UserMysql struct{}

func NewUserMysql() *UserMysql {
	return &UserMysql{}
}

func (c *UserMysql) UserIcon(userids []int) ([]pojo.UserIcon, error) {
	var m []pojo.UserIcon
	err := Db.Raw("SELECT t_user.id, username,url FROM t_user where t_user.id in ?", userids).Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *UserMysql) UserPwd(account string) (string, error) {
	var m string
	err := Db.Raw("SELECT password FROM t_user WHERE account = ?", account).Scan(&m).Error
	if err != nil {
		return "", err
	}
	return m, nil
}

func (c *UserMysql) RegisterUser(usermsg pojo.T_user) (int, error) {
	m := Db.Create(&usermsg)
	rowsaffected := m.RowsAffected

	if rowsaffected == 0 {
		return 0, errors.New("Insert---User---Mesg---Error")
	}
	return usermsg.Id, nil
}

func (c *UserMysql) RegisterTeacher(name string, userid int) error {
	m := Db.Exec("INSERT INTO t_teacher VALUES (NULL,?,?,NOW(),NOW());", name, userid)
	rowsaffected := m.RowsAffected

	if rowsaffected == 0 {
		return errors.New("Insert---Teacher---Mesg---Error")
	}
	return nil
}

func (c *UserMysql) SearhcUserId(account string) (int, error) {
	var m int
	err := Db.Raw("SELECT id FROM t_user WHERE account = ? ", account).Scan(&m).Error
	if err != nil {
		return 0, err
	}
	return m, nil
}

func (c *UserMysql) InsertTeacher(teachermsg pojo.T_teacher, db *gorm.DB) error {
	m := db.Create(&teachermsg)
	rowsaffected := m.RowsAffected

	if rowsaffected == 0 {
		return errors.New("Insert---Teacher---Mesg---Error")
	}
	return nil
}

func (c *UserMysql) UpdateTGroupTeacherId(group string, teacherid int, db *gorm.DB) error {
	m := db.Exec("UPDATE t_group SET teacher_id = ? WHERE name = ?", teacherid, group)
	rowsaffected := m.RowsAffected

	if rowsaffected == 0 {
		return errors.New("Update---Group-Teacher---Mesg---Error")
	}
	return nil
}

func (c *UserMysql) InsertStudent(studentmsg pojo.T_student, db *gorm.DB) error {
	m := db.Create(&studentmsg)
	rowsaffected := m.RowsAffected

	if rowsaffected == 0 {
		return errors.New("Insert---Student---Mesg---Error")
	}
	return nil
}

func (c *UserMysql) SearchGroupidBygroupname(groupname string) (int, error) {
	var m int
	err := Db.Raw("SELECT id FROM t_group WHERE name = ?", groupname).Scan(&m).Error
	if err != nil {
		return 0, err
	}
	return m, nil
}

func (c *UserMysql) InsertGraduate(graduatemsg pojo.T_graduate, db *gorm.DB) error {
	m := db.Create(&graduatemsg)
	rowsaffected := m.RowsAffected

	if rowsaffected == 0 {
		return errors.New("Insert---Graduate---Mesg---Error")
	}
	return nil
}
