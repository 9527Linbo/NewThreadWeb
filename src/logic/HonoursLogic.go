package logic

import (
	mapper "NewThread/src/mapper/mysql"
	"NewThread/src/pojo"
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

func (c *HonoursLogic) HonoursStudents() ([]pojo.StudentHonours, error) {
	mes, err := mapper.NewHonoursMysql().HonoursStudentsMysql()

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
