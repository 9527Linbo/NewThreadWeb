package mapper

import (
	"NewThread/src/pojo"
)

type GroupMysql struct{}

func NewGroupMysql() *GroupMysql {
	return &GroupMysql{}
}

func (c *GroupMysql) GroupList() ([]pojo.Group, error) {
	var m []pojo.Group
	if err := Db.Raw("select id,name,description from t_group").Scan(&m).Error; err != nil {
		return []pojo.Group{}, err
	}
	return m, nil
}

/*
查询 所有老师 和 老师负责的小组
返回 老师的用户ID、Name、Group
*/
func (c *GroupMysql) GroupTeacherListAndGroup() ([]pojo.Teacher, error) {
	var m []pojo.Teacher
	err := Db.Raw("SELECT t.id,t.user_id,t.`name`,g.name `group` FROM t_teacher t LEFT JOIN t_group g ON t.id = g.teacher_id").Scan(&m).Error
	if err != nil {
		return []pojo.Teacher{}, err
	}
	return m, nil
}

func (c *GroupMysql) PositionByUserId(userid int) ([]string, error) {
	var m []string
	if err := Db.Raw("SELECT name FROM t_position where user_id = ?", &userid).Scan(&m).Error; err != nil {
		return nil, err
	}
	return m, nil
}

func (c *GroupMysql) GroupStudentList() ([]pojo.Group, error) {
	var m []pojo.Group
	if err := Db.Raw("select id,name,description from t_group").Scan(&m).Error; err != nil {
		return []pojo.Group{}, err
	}
	return m, nil
}
