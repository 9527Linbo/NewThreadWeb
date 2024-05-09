package logic

import (
	mapper "NewThread/src/mapper/mysql"
	"NewThread/src/pojo"
	"time"
)

type HonoursLogic struct{} //所有size为0的变量都用的是同一块内存  zerobase

func NewHonoursService() *HonoursLogic {
	return &HonoursLogic{}
}

func (c *HonoursLogic) HonoursList() ([]pojo.Honours, error) {

	mes, err := mapper.NewHonoursMysql().HonoursTypeMysql()

	if err != nil {
		return nil, err
	}

	for i := range mes {

		honoursTeam, err := mapper.NewHonoursMysql().HonoursTeamMysql(mes[i].ID)

		if err != nil {
			return nil, err
		}

		mes[i].TeamMessage = honoursTeam
	}
	return mes, err
}

func (c *HonoursLogic) HonoursGraduate() ([]pojo.StudentHonours, error) {
	mes, err := mapper.NewHonoursMysql().HonoursGraduateMysql()

	if err != nil {
		return nil, err
	}

	return mes, err
}

func (c *HonoursLogic) HonoursProjects() (mes []pojo.Project, err error) {

	mes, err = mapper.NewHonoursMysql().HonoursProjectsMysql()

	if err != nil {
		return nil, err
	}

	//根据ID 查询URL
	for i := range mes {

		mes[i].URL, err = mapper.NewHonoursMysql().HonoursProjectURLMysql(mes[i].ID)

		if err != nil {
			return nil, err
		}
	}

	//根据ID 查询获得奖项
	for i := range mes {

		mes[i].Honours, err = mapper.NewHonoursMysql().HonoursProjectHonoursMysql(mes[i].ID)

		if err != nil {
			return nil, err
		}
	}
	//没出现错误则返回查询信息
	return mes, err
}

func (c *HonoursLogic) HonoursMilestone() (mes []pojo.Project, err error) {

	mes, err = mapper.NewHonoursMysql().HonoursProjectsMilestoneMysql()

	if err != nil {
		return nil, err
	}

	//根据ID 查询URL
	for i := range mes {

		mes[i].URL, err = mapper.NewHonoursMysql().HonoursProjectURLMysql(mes[i].ID)

		if err != nil {
			return nil, err
		}
	}

	//根据ID 查询获得奖项
	for i := range mes {

		mes[i].Honours, err = mapper.NewHonoursMysql().HonoursProjectHonoursMysql(mes[i].ID)

		if err != nil {
			return nil, err
		}
	}
	//没出现错误则返回查询信息
	return mes, err
}

func (c *HonoursLogic) HonoursMilestones() (mes []pojo.Project, err error) {

	mes, err = mapper.NewHonoursMysql().HonoursProjectsMilestonesMysql()

	if err != nil {
		return nil, err
	}
	//根据ID 查询URL
	for i := range mes {

		mes[i].URL, err = mapper.NewHonoursMysql().HonoursProjectURLMysql(mes[i].ID)

		if err != nil {
			return nil, err
		}
	}
	//根据ID 查询获得奖项
	for i := range mes {

		mes[i].Honours, err = mapper.NewHonoursMysql().HonoursProjectHonoursMysql(mes[i].ID)

		if err != nil {
			return nil, err
		}
	}
	//没出现错误则返回查询信息
	return mes, err
}

func (c *HonoursLogic) AddMilestone(projectmsg pojo.Project, vis byte) (pojo.Project, error) {
	project := pojo.T_project{
		Name:        projectmsg.Name,
		Type:        projectmsg.Type,
		Description: projectmsg.Description,
		IsMilestone: vis,
		CreatTime:   time.Now(),
		UpdateTime:  time.Now(),
	}

	tx := mapper.Db.Begin()

	//捕获异常
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := func() error {
		prjid, err := mapper.NewHonoursMysql().InsertProjectMysql(project, tx)
		if err != nil {
			return err
		}

		projectawards := make([]pojo.T_projectawards, len(projectmsg.Honours))
		for i, honour := range projectmsg.Honours {
			projectawards[i] = pojo.T_projectawards{
				ProjectId:  prjid,
				Rank:       honour.Rank,
				Time:       honour.Time,
				CreatTime:  time.Now(),
				UpdateTime: time.Now(),
			}
		}

		err = mapper.NewHonoursMysql().InsertProjectawardsMysql(projectawards, tx)
		if err != nil {
			return err
		}

		return tx.Commit().Error
	}(); err != nil {
		tx.Rollback()
		return pojo.Project{}, err
	}

	return projectmsg, nil
}

func (c *HonoursLogic) AddHonour(racemsg pojo.Honours, type_ string, endtime string) (pojo.Honours, error) {
	//查询 老师id和负责组id
	teacherid, err := mapper.NewUserMysql().SearhcTeacherUserId(racemsg.TeacherName)
	if err != nil {
		return pojo.Honours{}, err
	}

	groupid, err := mapper.NewUserMysql().SearchGroupidBygroupname(racemsg.Name)
	if err != nil {
		return pojo.Honours{}, err
	}

	//写入race表
	race := pojo.T_race{
		GroupId:    groupid,
		TeacherId:  teacherid,
		Name:       racemsg.Name,
		Type:       type_,
		StartTime:  racemsg.StartTime,
		EndTime:    endtime,
		CreatTime:  time.Now(),
		UpdateTime: time.Now(),
	}

	//将 获奖情况写入raceawards表

	tx := mapper.Db.Begin()

	//捕获异常
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := func() error {
		raceid, err := mapper.NewHonoursMysql().InsertRaceMysql(race, tx)
		if err != nil {
			return err
		}

		raceawards := make([]pojo.T_raceawards, len(racemsg.TeamMessage))
		for i, honour := range racemsg.TeamMessage {
			raceawards[i] = pojo.T_raceawards{
				RaceId:     raceid,
				Rank:       honour.Rank,
				Number:     honour.Number,
				TeamName:   honour.HonoursTeam,
				CreatTime:  time.Now(),
				UpdateTime: time.Now(),
			}
		}

		err = mapper.NewHonoursMysql().InsertRaceawardsMysql(raceawards, tx)
		if err != nil {
			return err
		}

		return tx.Commit().Error
	}(); err != nil {
		tx.Rollback()
		return pojo.Honours{}, err
	}

	return racemsg, nil
}
