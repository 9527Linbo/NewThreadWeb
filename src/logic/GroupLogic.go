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
	return mapper.NewGroupMysql().GroupListMysql()
}

func (c *GroupLogic) GroupTeacherInfo() ([]pojo.Teacher, error) {

	//查询所有老师的姓名、负责小组、用户id
	mes, err := mapper.NewGroupMysql().GroupTeacherListAndGroupMysql()

	if err != nil {
		return nil, err
	}

	//根据老师的User_id循环查找position表.
	for i := range mes {

		position, err := mapper.NewGroupMysql().PositionByUserIdMysql(mes[i].Userid)

		if err != nil {
			return nil, err
		}

		mes[i].Position = position
	}

	return mes, err
}

func (c *GroupLogic) GroupStudentInfo(year int) ([]pojo.Student, error) {
	mes, err := mapper.NewGroupMysql().GroupStudentListAndWishesMysql(year)
	if err != nil {
		return nil, err
	}
	//根据学生的User_id循环查找position表.
	for i := range mes {
		position, err := mapper.NewGroupMysql().PositionByUserIdMysql(mes[i].Userid)
		if err != nil {
			return nil, err
		}
		mes[i].Position = position
	}
	return mes, err
}

func (c *GroupLogic) Yearlist() ([]int, error) {
	return mapper.NewGroupMysql().YearlistMysql()
}
