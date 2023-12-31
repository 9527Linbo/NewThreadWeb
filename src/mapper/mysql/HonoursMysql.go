package mapper

import "NewThread/src/pojo"

type HonoursMysql struct{}

func NewHonoursMysql() *HonoursMysql {
	return &HonoursMysql{}
}

func (c *HonoursMysql) HonoursTypeMysql() ([]pojo.Honours, error) {
	var m []pojo.Honours
	err := Db.Raw("SELECT a.id, a.`name`, t.`name` teacher, g.`name` `group`, start_time FROM t_awardsname a LEFT JOIN t_teacher t ON t.user_id = a.teacher_id LEFT JOIN t_group g ON g.id = a.group_id ").Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return m, err
}

func (c *HonoursMysql) HonoursTeamMysql(id int) ([]pojo.HonoursTeam, error) {
	var m []pojo.HonoursTeam
	err := Db.Raw("SELECT `rank`,teamname,number FROM t_awards WHERE type_id = ?", id).Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return m, err
}

func (c *HonoursMysql) HonoursStudentsMysql() ([]pojo.StudentHonours, error) {
	var m []pojo.StudentHonours
	err := Db.Raw("SELECT name,description,img.url FROM t_studentgraduate s LEFT JOIN t_imageuser img ON s.user_id = img.user_id").Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return m, err
}

func (c *HonoursMysql) HonoursProjectsMysql() ([]pojo.Project, error) {
	var m []pojo.Project
	err := Db.Raw("SELECT id,name,description,type FROM t_project").Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return m, err
}

func (c *HonoursMysql) HonoursProjectURLMysql(id int) ([]pojo.ImgURL, error) {
	var m []pojo.ImgURL
	err := Db.Raw("SELECT url FROM t_imageproject WHERE project_id = ?", id).Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return m, err
}

func (c *HonoursMysql) HonoursProjectHonoursMysql(id int) ([]pojo.ProjectHonours, error) {
	var m []pojo.ProjectHonours
	err := Db.Raw("SELECT `name`,`rank`,time FROM t_awardsproject p LEFT JOIN t_awardsname n ON p.type_id = n.id WHERE p.project_id = ?", id).Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return m, err
}

func (c *HonoursMysql) HonoursProjectsMilestoneMysql() ([]pojo.Project, error) {
	var m []pojo.Project
	err := Db.Raw("SELECT id,name,description,type FROM t_project order by create_time desc limit 6").Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return m, err
}

func (c *HonoursMysql) HonoursProjectsMilestonesMysql() ([]pojo.Project, error) {
	var m []pojo.Project
	err := Db.Raw("SELECT id,name,description,type FROM t_project").Scan(&m).Error
	if err != nil {
		return nil, err
	}
	return m, err
}
