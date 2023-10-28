package logic

import (
	mapper "NewThread/src/mapper/mysql"
	"NewThread/src/pojo"
)

type GroupLogic struct{} //所有size为0的变量都用的是同一块内存  zerobase

func NewGroupService() *GroupLogic {
	return &GroupLogic{}
}

func (c *GroupLogic) GroupInfo() ([]pojo.Group, error) {
	return mapper.NewGroupMysql().GroupList()
}

func (c *GroupLogic) GroupTeacherInfo() ([]pojo.Teacher, error) {
	//查询所有老师的姓名、负责小组、用户id
	mes, err := mapper.NewGroupMysql().GroupTeacherListAndGroup()
	if err != nil {
		return nil, err
	}

	//根据老师的User_id循环查找position表.
	for temp := range mes {
		position, err := mapper.NewGroupMysql().PositionByUserId(mes[temp].Userid)
		if err != nil {
			return nil, err
		}
		mes[temp].Position = position
	}
	return mes, err
}

func (c *GroupLogic) GroupStudentInfo() ([]pojo.Group, error) {
	return mapper.NewGroupMysql().GroupStudentList()
}
