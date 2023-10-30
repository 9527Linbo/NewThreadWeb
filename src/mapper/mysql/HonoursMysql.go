package mapper

import "NewThread/src/pojo"

type HonoursMysql struct{}

func NewHonoursMysql() *HonoursMysql {
	return &HonoursMysql{}
}

func (c *HonoursMysql) HonoursTypeMysql() ([]pojo.Honours, error) {
	var m []pojo.Honours
	err := Db.Raw("SELECT a.id, a.`name`, t.`name` teacher, g.`name` `group`, start_time FROM t_awardsName a LEFT JOIN t_teacher t ON t.user_id = a.teacher_id LEFT JOIN t_group g ON g.id = a.group_id ").Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return m, err
}

func (c *HonoursMysql) HonoursTeamMysql(id int) ([]pojo.HonoursTeam, error) {
	var m []pojo.HonoursTeam
	err := Db.Raw("SELECT rank,TeamName,number FROM t_awards WHERE type_id = ?", id).Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return m, err
}
