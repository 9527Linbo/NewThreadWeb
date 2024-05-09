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

// 获取用户头像
func (c *UserMysql) UserIcon(userids []int) ([]pojo.UserIcon, error) {
	var m []pojo.UserIcon
	err := Db.Raw("SELECT t_user.id, username,url FROM t_user where t_user.id in ?", userids).Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

// 获取用户密码
func (c *UserMysql) UserPwd(account string) (string, error) {
	var m string
	err := Db.Raw("SELECT password FROM t_user WHERE account = ?", account).Scan(&m).Error
	if err != nil {
		return "", err
	}
	return m, nil
}

// 插入用户数据
func (c *UserMysql) RegisterUser(usermsg pojo.T_user, db *gorm.DB) (int, error) {
	m := db.Create(&usermsg)
	rowsaffected := m.RowsAffected

	if rowsaffected == 0 {
		return 0, errors.New("Insert---User---Mesg---Error")
	}
	return usermsg.Id, nil
}

// 查询用户id 通过账号
func (c *UserMysql) SearhcUserId(account string) (int, error) {
	var m int
	err := Db.Raw("SELECT id FROM t_user WHERE account = ? ", account).Scan(&m).Error
	if err != nil {
		return 0, err
	}
	return m, nil
}

// 查询老师的用户id 通过名字
func (c *UserMysql) SearhcTeacherUserId(name string) (int, error) {
	var m int
	err := Db.Raw("SELECT user_id FROM t_teacher WHERE name = ? ", name).Scan(&m).Error
	if err != nil {
		return 0, err
	}
	return m, nil
}

// 插入老师表数据
func (c *UserMysql) InsertTeacher(teachermsg pojo.T_teacher, db *gorm.DB) error {
	m := db.Create(&teachermsg)
	rowsaffected := m.RowsAffected

	if rowsaffected == 0 {
		return errors.New("Insert---Teacher---Mesg---Error")
	}
	return nil
}

// 更新小组指导老师id
func (c *UserMysql) UpdateTGroupTeacherId(group string, teacherid int, db *gorm.DB) error {
	m := db.Exec("UPDATE t_group SET teacher_id = ? WHERE name = ?", teacherid, group)
	print(teacherid, group)
	rowsaffected := m.RowsAffected

	if rowsaffected == 0 {
		return errors.New("Update---Group-Teacher---Mesg---Error")
	}
	return nil
}

// 插入学生表数据
func (c *UserMysql) InsertStudent(studentmsg pojo.T_student, db *gorm.DB) error {
	m := db.Create(&studentmsg)
	rowsaffected := m.RowsAffected

	if rowsaffected == 0 {
		return errors.New("Insert---Student---Mesg---Error")
	}
	return nil
}

// 查询小组id通过小组名称
func (c *UserMysql) SearchGroupidBygroupname(groupname string) (int, error) {
	var m int
	err := Db.Raw("SELECT id FROM t_group WHERE name = ?", groupname).Scan(&m).Error
	if err != nil {
		return 0, err
	}
	return m, nil
}

// 插入毕业生表数据
func (c *UserMysql) InsertGraduate(graduatemsg pojo.T_graduate, db *gorm.DB) error {
	m := db.Create(&graduatemsg)
	rowsaffected := m.RowsAffected

	if rowsaffected == 0 {
		return errors.New("Insert---Graduate---Mesg---Error")
	}
	return nil
}

// 查询用户名称 通过id
func (c *UserMysql) SearhcUserName(userid int) (string, error) {
	var m string
	err := Db.Raw("SELECT username FROM t_user WHERE id = ? ", userid).Scan(&m).Error
	if err != nil {
		return "", err
	}
	return m, nil
}

func (c *UserMysql) UserList() ([]pojo.User, error) {
	var m []pojo.User
	err := Db.Raw("SELECT u.id,u.username,s.name,u.create_Time,u.url ,'学生' AS identity FROM t_user u INNER JOIN t_student s ON u.id = s.user_id UNION" +
		" SELECT u.id,u.username,t.name,u.create_Time,u.url ,'老师' AS identity FROM t_user u INNER JOIN t_teacher t ON u.id = t.user_id UNION" +
		" SELECT u.id,u.username,u.username,u.create_Time,u.url ,'普通用户' AS identity FROM t_user u WHERE u.id NOT IN (SELECT user_id FROM t_student UNION" +
		" SELECT user_id FROM t_teacher);").Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *UserMysql) DelUser(userid string, db *gorm.DB) (string, error) {
	var iconuuid string
	err := db.Raw("SELECT iconname FROM t_user where id = ?", userid).Scan(&iconuuid).Error
	if err != nil {
		return "", err
	}

	m := db.Exec("DELETE from t_user where id = ?", userid)
	if m.RowsAffected == 0 {
		return "", errors.New("Delete---User---Error")
	}
	return iconuuid, nil
}

func (c *UserMysql) DelTeacher(userid string, db *gorm.DB) error {
	m := db.Exec("DELETE from t_teacher where user_id = ?", userid)
	if m.RowsAffected == 0 {
		return errors.New("Delete---Teacher---Error")
	}
	return nil
}

func (c *UserMysql) DelStudent(userid string, db *gorm.DB) error {
	m := db.Exec("DELETE from t_student where user_id = ?", userid)
	if m.RowsAffected == 0 {
		return errors.New("Delete---student---Error")
	}
	return nil
}

func (c *UserMysql) DelGraduate(graduateId string, db *gorm.DB) (string, error) {
	var iconuuid string
	err := db.Raw("SELECT url FROM t_graduate where id = ?", graduateId).Scan(&iconuuid).Error
	if err != nil {
		return "", err
	}

	m := db.Exec("DELETE from t_graduate where id = ?", graduateId)
	if m.RowsAffected == 0 {
		return "", errors.New("Delete---Graduate---Error")
	}
	return iconuuid, nil
}
