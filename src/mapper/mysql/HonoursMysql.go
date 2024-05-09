package mapper

import (
	"NewThread/src/pojo"
	"errors"

	"gorm.io/gorm"
)

type HonoursMysql struct{}

func NewHonoursMysql() *HonoursMysql {
	return &HonoursMysql{}
}

func (c *HonoursMysql) HonoursTypeMysql() ([]pojo.Honours, error) {
	var m []pojo.Honours

	err := Db.Raw("SELECT a.id, a.`name`, t.`name` teacher, g.`name` `group`, start_time FROM t_race a LEFT JOIN t_teacher t ON t.user_id = a.teacher_id LEFT JOIN t_group g ON g.id = a.group_id ").Scan(&m).Error

	if err != nil {
		return nil, err
	}

	return m, err
}

func (c *HonoursMysql) HonoursTeamMysql(id int) ([]pojo.HonoursTeam, error) {
	var m []pojo.HonoursTeam

	err := Db.Raw("SELECT `rank`,teamname,number FROM t_raceawards WHERE race_id = ?", id).Scan(&m).Error

	if err != nil {
		return nil, err
	}

	return m, err
}

func (c *HonoursMysql) HonoursGraduateMysql() ([]pojo.StudentHonours, error) {

	var m []pojo.StudentHonours

	err := Db.Raw("SELECT s.id,name,description,img.url FROM t_graduate s LEFT JOIN t_user img ON s.user_id = img.id").Scan(&m).Error

	if err != nil {
		return nil, err
	}

	return m, err
}

func (c *HonoursMysql) HonoursProjectsMysql() ([]pojo.Project, error) {

	var m []pojo.Project

	err := Db.Raw("SELECT id,name,description,type FROM t_project WHERE Ismilestone = 0").Scan(&m).Error

	if err != nil {
		return nil, err
	}

	return m, err
}

func (c *HonoursMysql) HonoursProjectURLMysql(id int) ([]pojo.ImgURL, error) {

	var m []pojo.ImgURL

	err := Db.Raw("SELECT url FROM t_projectimage WHERE project_id = ?", id).Scan(&m).Error

	if err != nil {
		return nil, err
	}

	return m, err
}

func (c *HonoursMysql) HonoursProjectHonoursMysql(id int) ([]pojo.ProjectHonours, error) {

	var m []pojo.ProjectHonours

	err := Db.Raw("SELECT `rank`,time FROM t_projectawards p WHERE p.project_id = ?", id).Scan(&m).Error

	if err != nil {
		return nil, err
	}

	return m, err
}

func (c *HonoursMysql) HonoursProjectsMilestoneMysql() ([]pojo.Project, error) {

	var m []pojo.Project

	err := Db.Raw("SELECT id,name,description,type FROM t_project WHERE Ismilestone = 1 order by create_time desc limit 6").Scan(&m).Error

	if err != nil {
		return nil, err
	}

	return m, err
}

func (c *HonoursMysql) HonoursProjectsMilestonesMysql() ([]pojo.Project, error) {

	var m []pojo.Project

	err := Db.Raw("SELECT id,name,description,type FROM t_project WHERE Ismilestone = 1").Scan(&m).Error

	if err != nil {
		return nil, err
	}

	return m, err
}

func (c *HonoursMysql) InsertProjectMysql(project pojo.T_project, db *gorm.DB) (int, error) {

	m := db.Create(&project)
	rowsaffected := m.RowsAffected

	if rowsaffected == 0 {
		return 0, errors.New("Insert---Project---Mesg---Error")
	}
	return project.Id, nil
}

func (c *HonoursMysql) InsertProjectawardsMysql(projectawards []pojo.T_projectawards, db *gorm.DB) error {

	m := db.CreateInBatches(projectawards, len(projectawards))
	rowsaffected := m.RowsAffected

	if rowsaffected == 0 || rowsaffected < int64(len(projectawards)) {
		return errors.New("Insert---Project---Mesg---Error")
	}
	return nil
}

func (c *HonoursMysql) InsertRaceMysql(race pojo.T_race, db *gorm.DB) (int, error) {

	m := db.Create(&race)
	rowsaffected := m.RowsAffected

	if rowsaffected == 0 {
		return 0, errors.New("Insert---Race---Mesg---Error")
	}
	return race.Id, nil
}

func (c *HonoursMysql) InsertRaceawardsMysql(raceawards []pojo.T_raceawards, db *gorm.DB) error {

	m := db.CreateInBatches(raceawards, len(raceawards))
	rowsaffected := m.RowsAffected

	if rowsaffected == 0 || rowsaffected < int64(len(raceawards)) {
		return errors.New("Insert---Project---Mesg---Error")
	}
	return nil
}
