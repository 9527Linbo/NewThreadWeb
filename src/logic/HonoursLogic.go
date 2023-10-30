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
